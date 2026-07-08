package bcio

import (
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/bytecode/constants"
	"focal-vm/internal/bytecode/spec"
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

	tpool := br.readTypePool()
	cpool := br.readConstantPool()

	module := spec.NewBCModule(
		versionMajor, versionMinor,
		string(name), tpool, cpool,
	)

	funcCount := br.readU32LE()
	funArray := make([]*spec.BCFunction, funcCount)
	for i := range funcCount {
		funArray[i] = br.readFunction(module)
	}
	module.SetFunctions(funArray)

	return module
}

func (br *ModuleReader) readTypePool() *bctypes.TypePool {
	pool := bctypes.NewTypePool()
	count := br.readU32LE()
	for range count {
		pool.CopyToPool(br.readType(pool))
	}
	return pool
}

func (br *ModuleReader) readType(tpool *bctypes.TypePool) bctypes.BCType {
	tag := bctypes.BCTypeTag(br.readU8LE())
	switch tag {
	case bctypes.BCTYPE_I8:
		return bctypes.I8
	case bctypes.BCTYPE_I16:
		return bctypes.I16
	case bctypes.BCTYPE_I32:
		return bctypes.I32
	case bctypes.BCTYPE_I64:
		return bctypes.I64
	case bctypes.BCTYPE_F32:
		return bctypes.F32
	case bctypes.BCTYPE_F64:
		return bctypes.F64
	case bctypes.BCTYPE_U8:
		return bctypes.U8
	case bctypes.BCTYPE_U16:
		return bctypes.U16
	case bctypes.BCTYPE_U32:
		return bctypes.U32
	case bctypes.BCTYPE_U64:
		return bctypes.U64
	case bctypes.BCTYPE_UTFSTRING:
		return bctypes.UTF_STRING
	case bctypes.BCTYPE_BOOLEAN:
		return bctypes.BOOL

	case bctypes.BCTYPE_ARRAY:
		elemType := br.readI32LE()
		return bctypes.NewArrayType(tpool, elemType)

	case bctypes.BCTYPE_STRUCT:
		nameLen := br.readI32LE()
		nameBytes := br.readU8ArrLE(nameLen)
		name := string(nameBytes)
		typeCount := br.readI32LE()
		types := br.readI32ArrLE(typeCount)
		return bctypes.NewStructType(tpool, name, types...)
	case bctypes.BCTYPE_MODULE:
		nameLen := br.readI32LE()
		nameBytes := br.readU8ArrLE(nameLen)
		name := string(nameBytes)
		return bctypes.NewModuleType(tpool, name)
	case bctypes.BCTYPE_FUNCTION:
		paramTypeCount := br.readI32LE()
		paramTypes := br.readI32ArrLE(paramTypeCount)

		returnTypeCount := br.readI32LE()
		returnTypes := br.readI32ArrLE(returnTypeCount)
		return bctypes.NewFunctionType(tpool, paramTypes, returnTypes...)
	default:
		panic("Unknown Type Tag #" + strconv.Itoa(int(tag)))
	}
	return nil
}

func (br *ModuleReader) readConstantPool() *constants.ConstantPool {
	pool := constants.NewConstantPool()
	constantCount := br.readU32LE()
	for range constantCount {
		tag := constants.ConstantTag(br.readU8LE())
		switch tag {
		case constants.ConstantTagBoolean:
			pool.GetOrCreateI8(int8(br.readU8LE()))
		case constants.ConstantTagInt8:
			pool.GetOrCreateI8(int8(br.readU8LE()))
		case constants.ConstantTagInt16:
			pool.GetOrCreateI16(int16(br.readU16LE()))
		case constants.ConstantTagInt32:
			pool.GetOrCreateI32(br.readI32LE())
		case constants.ConstantTagInt64:
			pool.GetOrCreateI64(int64(br.readU64LE()))
		case constants.ConstantTagFloat32:
			pool.GetOrCreateF32(math.Float32frombits(br.readU32LE()))
		case constants.ConstantTagFloat64:
			pool.GetOrCreateF64(math.Float64frombits(br.readU64LE()))
		case constants.ConstantTagUTF8String:
			pool.GetOrCreateUTF8(string(br.readU8ArrLE(br.readI32LE())))

		default:
			panic("Unknown Constant Tag #" + strconv.Itoa(int(tag)))
		}
	}
	return pool
}

func (br *ModuleReader) readFunction(module *spec.BCModule) *spec.BCFunction {
	modifier := br.readU8LE()
	nameIdx := br.readI32LE()
	typeIdx := br.readI32LE()
	paramNameCount := br.readI32LE()
	paramNameIndexes := br.readI32ArrLE(paramNameCount)
	codeLen := br.readI32LE()
	code := br.readU8ArrLE(codeLen)

	return spec.NewBCFunction(module, modifier, nameIdx, typeIdx, paramNameIndexes, code)
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

func (br *ModuleReader) ReadVariableLEU32(size int32) uint32 {
	v := uint32(0)
	for i := range size {
		v <<= 8
		v |= uint32(br.arr[br.idx+i])
	}
	br.idx += size
	return v
}

func (br *ModuleReader) ReadVariableLEU64(size int32) uint64 {
	v := uint64(0)
	for i := range size {
		v <<= 8
		v |= uint64(br.arr[br.idx+i])
	}
	br.idx += size
	return v
}

func (br *ModuleReader) readI8LE() int8 {
	v := br.arr[br.idx]
	br.idx++
	return int8(v)
}

func (br *ModuleReader) readI8ArrLE(size int32) []int8 {
	out := make([]int8, size)
	for i := range size {
		out[i] = int8(br.arr[br.idx+i])
	}
	br.idx += size
	return out
}

func (br *ModuleReader) readI16LE() int16 {
	low := int16(br.readI8LE())
	high := int16(br.readI8LE())
	return low | high<<8
}

func (br *ModuleReader) readI16ArrLE(size int32) []int16 {
	out := make([]int16, size)
	for i := range size {
		out[i] = br.readI16LE()
	}
	return out
}

func (br *ModuleReader) readI32LE() int32 {
	low := int32(br.readI16LE())
	high := int32(br.readI16LE())
	return low | high<<16
}

func (br *ModuleReader) readI32ArrLE(size int32) []int32 {
	out := make([]int32, size)
	for i := range size {
		out[i] = br.readI32LE()
	}
	return out
}

func (br *ModuleReader) readI64LE() int64 {
	return int64(br.readI32LE()) | int64(br.readI32LE())<<32
}

func (br *ModuleReader) readI64ArrLE(size int32) []int64 {
	out := make([]int64, size)
	for i := range size {
		out[i] = br.readI64LE()
	}
	return out
}

func (br *ModuleReader) ReadVariableLEI32(size int32) int32 {
	v := int32(0)
	for i := range size {
		v <<= 8
		v |= int32(br.arr[br.idx+i])
	}
	br.idx += size
	return v
}

func (br *ModuleReader) ReadVariableLEI64(size int32) int64 {
	v := int64(0)
	for i := range size {
		v <<= 8
		v |= int64(br.arr[br.idx+i])
	}
	br.idx += size
	return v
}
