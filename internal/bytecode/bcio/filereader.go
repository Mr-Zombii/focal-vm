package bcio

import (
	"focal-lang/internal/bytecode/constants"
	"focal-lang/internal/bytecode/spec"
	"math"
	"strconv"
)

type ModuleReader struct {
	idx int32
	arr []byte
}

func NewReader(contents []byte) *ModuleReader {
	return &ModuleReader{idx: 0, arr: contents}
}

func (br *ModuleReader) ReadModule() *spec.BCModule {
	magic := br.readU8ArrLE(6)
	if !(magic[0] == spec.BC_MAGIC[0] && magic[1] == spec.BC_MAGIC[1] &&
		magic[2] == spec.BC_MAGIC[2] && magic[3] == spec.BC_MAGIC[3] &&
		magic[4] == spec.BC_MAGIC[4] && magic[5] == spec.BC_MAGIC[5]) {
		panic("Invalid FOCAL file type or corrupt magic number!")
	}

	versionMajor := br.readU8LE()
	versionMinor := br.readU8LE()

	nameLength := br.readU32LE()
	name := br.readU8ArrLE(int32(nameLength))

	pool := br.readConstantPool()
	mainFunctionIndex := br.readU32LE()

	module := spec.NewBCModule(
		versionMajor, versionMinor,
		string(name), int32(mainFunctionIndex),
		pool,
	)

	funcCount := br.readU32LE()
	funArray := make([]*spec.BCFunction, funcCount)
	for i := range funcCount {
		funArray[i] = br.readFunction()
	}
	module.SetFunctions(funArray)

	return module
}

func (br *ModuleReader) readConstantPool() *constants.ConstantPool {
	pool := constants.NewConstantPool()
	constantCount := br.readU32LE()
	for range constantCount {
		tag := constants.ConstantTag(br.readU8LE())
		switch tag {
		case constants.ConstantTagInt8:
			pool.GetOrCreateI8(int8(br.readU8LE()))
		case constants.ConstantTagInt16:
			pool.GetOrCreateI16(int16(br.readU16LE()))
		case constants.ConstantTagInt32:
			pool.GetOrCreateI32(int32(br.readU32LE()))
		case constants.ConstantTagInt64:
			pool.GetOrCreateI64(int64(br.readU64LE()))
		case constants.ConstantTagFloat32:
			pool.GetOrCreateF32(math.Float32frombits(br.readU32LE()))
		case constants.ConstantTagFloat64:
			pool.GetOrCreateF64(math.Float64frombits(br.readU64LE()))
		case constants.ConstantTagUTF8String:
			pool.GetOrCreateUTF8(string(br.readU8ArrLE(int32(br.readU32LE()))))
		default:
			panic("Unknown Constant Tag #" + strconv.Itoa(int(tag)))
		}
	}
	return pool
}

func (br *ModuleReader) readFunction() *spec.BCFunction {
	nameIdx := br.readU32LE()
	modifier := br.readU8LE()
	codeLen := br.readU32LE()
	code := br.readU8ArrLE(int32(codeLen))

	return spec.NewBCFunction(int32(nameIdx), modifier, code)
}

func (br *ModuleReader) skip(n int32) {
	br.idx += n
}

func (br *ModuleReader) readU8LE() uint8 {
	v := br.arr[br.idx]
	br.idx++
	return v
}

func (br *ModuleReader) readU8ArrLE(size int32) []uint8 {
	out := make([]uint8, size)
	for i := range size {
		out[i] = br.arr[br.idx+i]
	}
	br.idx += size
	return out
}

func (br *ModuleReader) readU16LE() uint16 {
	low := uint16(br.readU8LE())
	high := uint16(br.readU8LE())
	return low | high<<8
}

func (br *ModuleReader) readU16ArrLE(size int32) []uint16 {
	out := make([]uint16, size)
	for i := range size {
		out[i] = br.readU16LE()
	}
	return out
}

func (br *ModuleReader) readU32LE() uint32 {
	low := uint32(br.readU16LE())
	high := uint32(br.readU16LE())
	return low | high<<16
}

func (br *ModuleReader) readU32ArrLE(size int32) []uint32 {
	out := make([]uint32, size)
	for i := range size {
		out[i] = br.readU32LE()
	}
	return out
}

func (br *ModuleReader) readU64LE() uint64 {
	return uint64(br.readU32LE()) | uint64(br.readU32LE())<<32
}

func (br *ModuleReader) readU64ArrLE(size int32) []uint64 {
	out := make([]uint64, size)
	for i := range size {
		out[i] = br.readU64LE()
	}
	return out
}

func (br *ModuleReader) ReadVariableLE32(size int32) uint32 {
	v := uint32(0)
	for i := range size {
		v <<= 8
		v |= uint32(br.arr[br.idx+i])
	}
	br.idx += size
	return v
}

func (br *ModuleReader) ReadVariableLE64(size int32) uint64 {
	v := uint64(0)
	for i := range size {
		v <<= 8
		v |= uint64(br.arr[br.idx+i])
	}
	br.idx += size
	return v
}
