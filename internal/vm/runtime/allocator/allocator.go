package allocator

import (
	"errors"
	"fmt"
	"focal-vm/internal/util"
	"strconv"
	"unsafe"
)

type AnyPointer interface{}

type Allocator interface {
	Alloc(size int32) unsafe.Pointer
	Realloc(ptr AnyPointer, size int32) unsafe.Pointer
	Free(ptr AnyPointer)
	SizeOf(ptr AnyPointer) int32
	Zero(ptr AnyPointer)
	IsInvalidOrFree(ptr AnyPointer) bool
	GetBlock(ptr AnyPointer) *MemBlock
	CopyString(ptr AnyPointer, v string) (*string, *byte)
	String() string
	FreeAllBlocks()
	GetTotalBackingSize() int
	CopyStringAsCStr(ptr AnyPointer, v string) *byte
	GetAllocationCount() int32
	GetTimesAllocCalled() int32
	GetTimesFreeCalled() int32
}

type _Allocator struct {
	bytesPreAllocated int32
	memory            []byte
	blocks            []*MemBlock
	blockMap          map[uintptr]*MemBlock
	furthestBlock     *MemBlock
	blockCount        int
	usedBlockCount    int
	freedBlockCount   int
	freedSpace        int32
	usedSpace         int32
	lastAvalibleIdx   int32
	allocationCount   int32
	timesAllocCalled  int32
	timesFreeCalled   int32
}

func (a *_Allocator) FreeAllBlocks() {
	for _, v := range a.blocks {
		if !v.free {
			a.freeBlock(v)
		}
	}
}

func (a *_Allocator) String() string {
	//return fmt.Sprint(a.blocks)
	return fmt.Sprint(a.memory)
}

type MemBlock struct {
	Ptr   unsafe.Pointer
	start int32
	Size  int32
	idx   int32
	free  bool
}

func newMemBlock(a *_Allocator, ptr unsafe.Pointer, start int32, size int32) *MemBlock {
	block := &MemBlock{
		Ptr:   ptr,
		start: start,
		Size:  size,
		idx:   int32(a.blockCount),
	}
	a.blockCount++
	a.usedBlockCount++
	a.freedSpace -= size
	a.usedSpace += size

	a.blockMap[uintptr(block.Ptr)] = block
	a.blocks = append(a.blocks, block)

	return block
}

func (mv *MemBlock) setFree(a *_Allocator, doFree bool) {
	if doFree {
		a.usedBlockCount--
		a.freedBlockCount++
		a.freedSpace += mv.Size
		a.usedSpace -= mv.Size
	} else {
		a.usedBlockCount++
		a.freedBlockCount--
		a.freedSpace -= mv.Size
		a.usedSpace += mv.Size
	}
	mv.free = doFree
}

func (mv *MemBlock) deleteBlock(a *_Allocator) {
	a.blocks = append(a.blocks[:mv.idx], a.blocks[mv.idx+1:]...)
	a.blockCount--
	delete(a.blockMap, uintptr(mv.Ptr))
	mv.idx = -1
	for i := range a.blocks {
		a.blocks[i].idx = int32(i)
	}
}

func (mb *MemBlock) String() string {
	return fmt.Sprintf("MemBlock: {%d, %d, %d, %v} ", mb.Ptr, mb.start, mb.Size, mb.free)
}

func NewAllocator(baseSize int32) Allocator {
	return &_Allocator{bytesPreAllocated: baseSize, freedSpace: baseSize, memory: make([]byte, baseSize), blocks: []*MemBlock{}, blockMap: make(map[uintptr]*MemBlock)}
}

func (a *_Allocator) GetBlock(ptr AnyPointer) *MemBlock {
	return a.blockMap[uintptr(GetPtr(ptr))]
}

var extraToExpandBy int32 = 0

func (a *_Allocator) expand_backing(sizeToAlloc int32) {
	sizeToAlloc += extraToExpandBy

	//fmt.Println("Expanding Memory size of", a.bytesPreAllocated, "byte(s) by", sizeToAlloc, "byte(s) to", a.bytesPreAllocated+sizeToAlloc, "byte(s)")
	a.memory = append(a.memory, make([]byte, sizeToAlloc)...)
	a.bytesPreAllocated += sizeToAlloc
}

func (a *_Allocator) find_consecutive_freed_blocks(size int32) ([]*MemBlock, int32, error) {
	if a.blockCount == 0 || a.freedBlockCount == 0 {
		return nil, -1, errors.New("no freed blocks to check")
	}

	var consecutiveBlocks []*MemBlock
	var totalSize int32

	for _, mainBlock := range a.blocks {
		if !mainBlock.free {
			continue
		}

		if mainBlock.Size >= size {
			return []*MemBlock{mainBlock}, totalSize, nil
		}

		consecutiveBlocks, totalSize = a.findConsecutiveFreeBlocksTo(mainBlock, size)
		fmt.Println(consecutiveBlocks)
		if totalSize >= size {
			break
		}
	}

	if totalSize < size {
		return nil, -1, errors.New("could not find enough consecutive free blocks to meet target total size of " + strconv.Itoa(int(totalSize)) + " bytes")
	}

	return consecutiveBlocks, totalSize, nil
}

func (a *_Allocator) realloc_freed_blocks(size int32) (*MemBlock, error) {
	if a.freedBlockCount == 0 || a.blockCount == 0 {
		return nil, errors.New("no freed blocks were found to operate on")
	}

	mainBlock := a.findFreeBlock(size)
	var currentBlockSize int32
	if mainBlock == nil {
		var consecutiveBlocks []*MemBlock
		var err error

		consecutiveBlocks, currentBlockSize, err = a.find_consecutive_freed_blocks(size)
		if err != nil {
			return nil, err
		}
		mainBlock = a.merge(consecutiveBlocks...)
	} else {
		currentBlockSize = mainBlock.Size
	}

	if currentBlockSize > size {
		mainBlock, _ = a.split(mainBlock, size)
	}

	mainBlock.setFree(a, false)
	a.blockMap[uintptr(mainBlock.Ptr)] = mainBlock

	return mainBlock, nil
}

func (a *_Allocator) has_avaliable_space() bool {
	return a.freedSpace != 0
}

func (a *_Allocator) has_avalible_space_for_size(size int32) bool {
	return a.freedSpace >= size
}

//func (a *_Allocator) alloc_no_reused(size int32) (*MemBlock, error) {
//	var blockStart int32
//	var blockSize int32
//	var block *MemBlock
//	if a.blockCount != 0 {
//		blockStart = 0
//
//		if blockSize > size {
//			mainBlock, _ :=
//		}
//	}
//}

// func (a *_Allocator) realloc(block *MemBlock, size int32) (*MemBlock, error) {
// 	var blockOut *MemBlock
// 	var err error
// 	if block.Size > size {
// 		blockOut, _ = a.split(block, size)
// 		return blockOut, nil
// 	}
// 	if block.Size < size {
// 		blockOut, err = a.realloc_freed_blocks(size)
// 		if err == nil {
// 			a.copyRaw(block.start, block.Size, blockOut.start)
// 			a.freeBlock(block)
// 			return blockOut, nil
// 		}
// 		return nil, err
// 	}
// 	return block, nil
// }

// func (a *_Allocator) alloc(size int32) (*MemBlock, error) {
// 	if !a.has_avaliable_space() && !a.has_avalible_space_for_size(size) {
// 		return a.realloc_freed_blocks(size)
// 	}
// 	if !a.has_avalible_space_for_size(size) {
// 		a.expand_backing(size - a.freedSpace)
// 	}
// 	var memBlock *MemBlock
// 	if a.furthestBlock != nil && a.furthestBlock.free {
// 		memBlock = a.furthestBlock
// 		if memBlock.Size > size {
// 			memBlock, _ = a.split(memBlock, size)
// 		}
// 		if memBlock.Size < size {
// 			lastAvalibleIdx
// 		}
// 		return memBlock, nil
// 	}
// }

func (a *_Allocator) GetAllocationCount() int32 {
	return a.allocationCount
}

func (a *_Allocator) GetTimesAllocCalled() int32 {
	return a.timesAllocCalled
}

func (a *_Allocator) GetTimesFreeCalled() int32 {
	return a.timesFreeCalled
}

func (a *_Allocator) Alloc(size int32) unsafe.Pointer {
	a.timesAllocCalled++
	a.allocationCount++
	if size == 8 {
		block, _ := a.realloc_freed_blocks(size)
		return block.Ptr
	}
	var memIdx int32 = 0
	var block *MemBlock

	if a.blockCount != 0 {
		freeBlock := a.findFreeBlock(size)
		if freeBlock != nil && freeBlock.Size != size {
			freeBlock.setFree(a, false)
			freeBlock = a.blockMap[uintptr(a.Realloc(freeBlock.Ptr, size))]
		}
		if freeBlock != nil {
			freeBlock.setFree(a, false)
			block = freeBlock
			memIdx = freeBlock.start
			a.zero(freeBlock)
		} else {
			endBlock := a.findEndBlock()

			if endBlock.free {
				endBlock.Size = size
				endBlock.setFree(a, false)
				block = endBlock
				memIdx = endBlock.start
				a.zero(endBlock)
			} else {
				memIdx = endBlock.start + endBlock.Size
			}
		}
	}
	if memIdx+size > a.bytesPreAllocated {
		a.expand_backing((memIdx + size) - a.bytesPreAllocated)
	}

	ptr := &a.memory[memIdx]
	unsafePtr := unsafe.Pointer(ptr)
	//fmt.Println("Allocated:", unsafePtr, "Size:", size)

	if block == nil {
		block = newMemBlock(
			a,
			unsafePtr,
			memIdx,
			size,
		)
	}

	if a.blockCount != 1 {
		if memIdx > a.furthestBlock.start {
			a.furthestBlock = block
		}
	} else {
		a.furthestBlock = block
	}

	a.blockMap[uintptr(unsafePtr)] = block

	return unsafePtr
}

func (a *_Allocator) Realloc(ptr AnyPointer, size int32) unsafe.Pointer {
	unsafePtr := GetPtr(ptr)
	block := a.blockMap[uintptr(unsafePtr)]
	if block == nil {
		panic("Tried to realloc unallocated blocks or freed blocks!")
	}

	oldSize := block.Size
	sizeDiff := size - oldSize
	if sizeDiff < 1 {
		a.zeroRaw(block.start+size, -sizeDiff)
		block.Size = size
	}
	if sizeDiff > 1 {
		var moveBlock bool
		for _, otherBlock := range a.blocks {
			if otherBlock.Ptr == unsafePtr {
				continue
			}
			otherStart := otherBlock.start
			otherEnd := otherBlock.start + otherBlock.Size - 1

			for i := range size {
				blockAddr := block.start + i
				if blockAddr >= otherStart && blockAddr <= otherEnd {
					moveBlock = true
				}
				break
			}
			if moveBlock {
				break
			}
		}

		if !moveBlock {
			block.Size = size
			return unsafePtr
		}
		delete(a.blockMap, uintptr(unsafePtr))

		freeBlock := a.findFreeBlock(size)
		if freeBlock == nil {
			endBlock := a.findEndBlock()
			newIdx := endBlock.start + endBlock.Size
			a.zeroRaw(newIdx, size)
			blockPtr := unsafe.Pointer(&a.memory[newIdx])
			a.blockMap[uintptr(blockPtr)] = block
			block.Size = size
			block.Ptr = blockPtr
			block.start = newIdx
			return blockPtr
		}
		var blockToUse = freeBlock
		if freeBlock.Size != size {
			blockToUse, _ = a.split(freeBlock, size)
		}
		a.blockMap[uintptr(blockToUse.Ptr)] = blockToUse
		a.freeBlock(block)
		return blockToUse.Ptr
	}
	return unsafePtr
}

func (a *_Allocator) findEndBlock() *MemBlock {
	if a.blockCount == 0 {
		return nil
	}

	//var largestStart int32
	//var furthestBlock *MemBlock
	//for _, v := range a.blocks {
	//	if v.start >= largestStart {
	//		furthestBlock = v
	//	}
	//}

	return a.furthestBlock
}

func GetPtr(ptr AnyPointer) unsafe.Pointer {
	unsafePtrPtr := unsafe.Add(unsafe.Pointer(&ptr), util.PointerSize)
	unsafePtr := *(*unsafe.Pointer)(unsafePtrPtr)
	return unsafePtr
}

func (a *_Allocator) Zero(ptr AnyPointer) {
	unsafePtr := GetPtr(ptr)
	block := a.blockMap[uintptr(unsafePtr)]
	if block == nil {
		panic("Tried to zero unallocated or free blocks!")
	}

	a.zero(block)
}

func (a *_Allocator) zeroRaw(start int32, size int32) {
	for i := range size {
		a.memory[start+i] = 0
	}
}

func (a *_Allocator) copyRaw(aStart, aSize int32, bStart int32) {
	for i := range aSize {
		a.memory[bStart+i] = a.memory[aStart+i]
	}
}

func (a *_Allocator) findFreeBlock(size int32) *MemBlock {
	for _, block := range a.blocks {
		if block.Size >= size && block.free {
			return block
		}
	}
	return nil
}

func (a *_Allocator) split(block *MemBlock, size int32) (*MemBlock, *MemBlock) {
	if !block.free {
		panic("Tried to split non-free blocks!")
	}

	sizeA := size
	sizeB := block.Size - sizeA

	blockBStart := block.start + sizeA
	blockBPtr := &a.memory[blockBStart]
	blockBUnsafePtr := unsafe.Pointer(blockBPtr)

	blockB := newMemBlock(
		a,
		blockBUnsafePtr,
		blockBStart,
		sizeB,
	)
	a.freedSpace += sizeB
	a.usedSpace -= sizeB
	a.freeBlock(blockB)

	return block, blockB
}

func (a *_Allocator) findConsecutiveFreeBlocksTo(mainBlock *MemBlock, targetSize int32) ([]*MemBlock, int32) {
	totalSize := mainBlock.Size
	out := []*MemBlock{mainBlock}

	nextBlockStart := mainBlock.start + mainBlock.Size
	for totalSize < targetSize {
		consecutive := false
		for _, v := range a.blocks {
			if v.start == mainBlock.start {
				continue
			}

			if v.start == nextBlockStart {
				totalSize += v.Size
				nextBlockStart = v.start + v.Size
				out = append(out, v)
				consecutive = true
			}
		}
		if !consecutive {
			break
		}
	}

	return out, totalSize
}

func (a *_Allocator) merge(otherBlocks ...*MemBlock) *MemBlock {
	mainBlock := otherBlocks[0]
	var totalSize int32
	var nextBlockStart int32

	for _, otherBlock := range otherBlocks {
		if nextBlockStart != otherBlock.start {
			panic("Tried to merge non-consecutive blocks!")
		}
		nextBlockStart = otherBlock.start + otherBlock.Size
		totalSize += otherBlock.Size
		otherBlock.deleteBlock(a)
	}

	mainBlock.Size = totalSize

	return mainBlock
}

func (a *_Allocator) zero(block *MemBlock) {
	blockStart := block.start
	blockSize := block.Size

	for i := range blockSize {
		a.memory[blockStart+i] = 0
	}
}

/*
	string {
		dataPtr *byte 4/8 bytes
		len int 4/8 bytes
	}
*/

func GetTotalStringSize(v string) int32 {
	return 2*(util.PointerSize) + int32(len(v))
}

func StrLen(ptr AnyPointer) int32 {
	basePtr := GetPtr(ptr)
	var offs int64 = 0
	for {
		value := *(*byte)(unsafe.Add(basePtr, offs))
		if value == 0 {
			break
		}
		offs++
	}
	return int32(offs)
}

func (a *_Allocator) CopyStringAsCStr(ptr AnyPointer, v string) *byte {
	basePtr := GetPtr(ptr)
	bytes := []byte(v)
	for i, b := range bytes {
		currentPtr := (*byte)(unsafe.Add(basePtr, i))
		*currentPtr = b
	}
	endPtr := (*byte)(unsafe.Add(basePtr, len(bytes)))
	*endPtr = 0

	return (*byte)(basePtr)
}

func (a *_Allocator) CopyString(ptr AnyPointer, v string) (*string, *byte) {
	strLen := len(v)
	strBytes := []byte(v)
	dataOffs := util.PointerSize * 2

	stringPointer := GetPtr(ptr)
	dataPointer := unsafe.Add(stringPointer, dataOffs)

	bytePtr := dataPointer
	for _, b := range strBytes {
		*(*byte)(bytePtr) = b
		bytePtr = unsafe.Add(bytePtr, 1)
	}

	byteArrayPtrPtr := (**byte)(stringPointer)
	*byteArrayPtrPtr = (*byte)(dataPointer)
	lenPtr := (*int)(unsafe.Add(stringPointer, util.PointerSize))
	*lenPtr = strLen

	return (*string)(stringPointer), (*byte)(dataPointer)
}

func (a *_Allocator) SizeOf(ptr AnyPointer) int32 {
	block := a.blockMap[uintptr(GetPtr(ptr))]
	if block == nil {
		panic("Tried to get size of unallocated blocks or freed blocks!")
	}
	return block.Size
}

func (a *_Allocator) IsInvalidOrFree(ptr AnyPointer) bool {
	block := a.blockMap[uintptr(GetPtr(ptr))]
	return block == nil || block.free
}

func (a *_Allocator) Free(ptr AnyPointer) {
	a.timesFreeCalled++
	a.allocationCount--
	block := a.blockMap[uintptr(GetPtr(ptr))]
	if block == nil {
		panic("Tried to free pointer that was not allocated or was already freed")
	}

	//fmt.Println("Freed", block.Ptr, "Size", block.Size)
	a.freeBlock(block)
}

func (a *_Allocator) freeBlock(block *MemBlock) {
	block.setFree(a, true)
	a.zero(block)
}

func (a *_Allocator) GetTotalBackingSize() int {
	return len(a.memory)
}
