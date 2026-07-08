package bcio

import (
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/bytecode/constants"
	"focal-vm/internal/bytecode/spec"
	"math"
	"strconv"
)

type ModuleWriter struct {
	arr    []byte
	module *spec.BCModule
}

func NewWriter(module *spec.BCModule) *ModuleWriter {
	return &ModuleWriter{arr: []byte{}, module: module}
}

func (mw *ModuleWriter) GetBytes() []byte {
	return mw.arr
}

func (mw *ModuleWriter) WriteModule() {
	mw.writeU8ArrLE(spec.BC_MAGIC)
	mw.writeU8LE(mw.module.GetBytecodeMajorVersion())
	mw.writeU8LE(mw.module.GetBytecodeMinorVersion())

	mw.writeU32LE(uint32(len(mw.module.GetName())))
	mw.writeU8ArrLE([]byte(mw.module.GetName()))
	mw.writeTypePool(mw.module.GetTypePool())
	mw.writeConstantPool(mw.module.GetConstantPool())
	mw.writeU32LE(uint32(len(mw.module.GetFunctions())))
	for _, v := range mw.module.GetFunctions() {
		mw.writeFunction(v)
	}
}

func (mw *ModuleWriter) writeTypePool(pool *bctypes.TypePool) {
	typeCount := pool.Size()
	mw.writeU32LE(uint32(typeCount))
	for i := range typeCount {
		typeValue := pool.GetType(i)
		tag := typeValue.GetTag()
		mw.writeU8LE(uint8(tag))
		switch tag {
		case bctypes.BCTYPE_I8:
			continue
		case bctypes.BCTYPE_I16:
			continue
		case bctypes.BCTYPE_I32:
			continue
		case bctypes.BCTYPE_I64:
			continue
		case bctypes.BCTYPE_U8:
			continue
		case bctypes.BCTYPE_U16:
			continue
		case bctypes.BCTYPE_U32:
			continue
		case bctypes.BCTYPE_U64:
			continue
		case bctypes.BCTYPE_F32:
			continue
		case bctypes.BCTYPE_F64:
			continue
		case bctypes.BCTYPE_UTFSTRING:
			continue
		case bctypes.BCTYPE_BOOLEAN:
			continue
		case bctypes.BCTYPE_MODULE:
			name := typeValue.(*bctypes.ModuleType).GetName()
			mw.writeU32LE(uint32(len(name)))
			mw.writeU8ArrLE([]byte(name))
		case bctypes.BCTYPE_ARRAY:
			mw.writeI32LE(typeValue.(*bctypes.ArrayType).GetElementTypeIndex())
		case bctypes.BCTYPE_STRUCT:
			struc := typeValue.(*bctypes.StructType)
			name := struc.GetName()
			mw.writeU32LE(uint32(len(name)))
			mw.writeU8ArrLE([]byte(name))

			types := struc.GetFieldTypeIndexes()
			mw.writeU32LE(uint32(len(types)))
			mw.writeI32ArrLE(types)
		case bctypes.BCTYPE_FUNCTION:
			fn := typeValue.(*bctypes.FunctionType)
			paramTypes := fn.GetParamTypeIndexes()
			mw.writeU32LE(uint32(len(paramTypes)))
			mw.writeI32ArrLE(paramTypes)
			returnTypes := fn.GetReturnTypeIndexes()
			mw.writeU32LE(uint32(len(returnTypes)))
			mw.writeI32ArrLE(returnTypes)

		}
	}
}

func (mw *ModuleWriter) writeConstantPool(pool *constants.ConstantPool) {
	constantCount := pool.Size()
	mw.writeU32LE(uint32(constantCount))
	for i := range constantCount {
		constant := pool.GetConstant(i)
		mw.writeU8LE(uint8(constant.GetTag()))
		switch con := constant.(type) {
		case *constants.ConstantBoolean:
			if con.GetValue() {
				mw.writeU8LE(1)
				continue
			}
			mw.writeU8LE(0)
		case *constants.ConstantInt8:
			mw.writeU8LE(uint8(con.GetValue()))
		case *constants.ConstantInt16:
			mw.writeU16LE(uint16(con.GetValue()))
		case *constants.ConstantInt32:
			mw.writeU32LE(uint32(con.GetValue()))
		case *constants.ConstantInt64:
			mw.writeU64LE(uint64(con.GetValue()))
		case *constants.ConstantFloat32:
			mw.writeU32LE(math.Float32bits(con.GetValue()))
		case *constants.ConstantFloat64:
			mw.writeU64LE(math.Float64bits(con.GetValue()))
		case *constants.ConstantUTF8String:
			mw.writeU32LE(uint32(con.GetLength()))
			mw.writeU8ArrLE([]uint8(con.GetValue()))
		default:
			panic("Unknown Constant Tag #" + strconv.Itoa(int(con.GetTag())))
		}
	}
}

func (mw *ModuleWriter) writeFunction(fun *spec.BCFunction) {
	mw.writeU8LE(fun.GetModifier())
	mw.writeI32LE(fun.GetNameIndex())
	mw.writeI32LE(fun.GetTypeIdx())
	mw.writeI32LE(int32(len(fun.GetParamNameIndexes())))
	mw.writeI32ArrLE(fun.GetParamNameIndexes())
	mw.writeI32LE(int32(len(fun.GetCode())))
	mw.writeU8ArrLE(fun.GetCode())
}

func (mw *ModuleWriter) writeU8LE(v uint8) {
	mw.arr = append(mw.arr, v)
}

func (mw *ModuleWriter) writeU8ArrLE(v []uint8) {
	mw.arr = append(mw.arr, v...)
}

func (mw *ModuleWriter) writeU16LE(v uint16) {
	mw.arr = append(append(mw.arr, uint8(v&0xFF)), uint8((v&0xFF00)>>8))
}

func (mw *ModuleWriter) writeU16ArrLE(v []uint16) {
	for i := range len(v) {
		mw.writeU16LE(v[i])
	}
}

func (mw *ModuleWriter) writeU32LE(v uint32) {
	mw.writeU16LE(uint16(v & 0xFFFF))
	mw.writeU16LE(uint16((v & 0xFFFF0000) >> 16))
}

func (mw *ModuleWriter) writeU32ArrLE(v []uint32) {
	for i := range len(v) {
		mw.writeU32LE(v[i])
	}
}

func (mw *ModuleWriter) writeU64LE(v uint64) {
	mw.writeU32LE(uint32(v & 0xFFFFFFFF))
	mw.writeU32LE(uint32((v & 0xFFFFFFFF00000000) >> 32))
}

func (mw *ModuleWriter) writeU64ArrLE(v []uint64) {
	for i := range len(v) {
		mw.writeU64LE(v[i])
	}
}

func (mw *ModuleWriter) writeVariableLEU32(v uint32, size int32) {
	o := v
	for range size {
		mw.arr = append(mw.arr, uint8(o&0xFF))
		o >>= 8
	}
}

func (mw *ModuleWriter) writeVariableLEU64(v uint64, size int32) {
	o := v
	for range size {
		mw.arr = append(mw.arr, uint8(o&0xFF))
		o >>= 8
	}
}

func (mw *ModuleWriter) writeI8LE(v int8) {
	mw.writeU8LE(uint8(v))
}

func (mw *ModuleWriter) writeI8ArrLE(v []int8) {
	for i := range v {
		mw.writeU8LE(uint8(v[i]))
	}
}

func (mw *ModuleWriter) writeI16LE(v int16) {
	mw.writeU16LE(uint16(v))
}

func (mw *ModuleWriter) writeI16ArrLE(v []int16) {
	for i := range len(v) {
		mw.writeI16LE(v[i])
	}
}

func (mw *ModuleWriter) writeI32LE(v int32) {
	mw.writeU32LE(uint32(v))
}

func (mw *ModuleWriter) writeI32ArrLE(v []int32) {
	for i := range len(v) {
		mw.writeI32LE(v[i])
	}
}

func (mw *ModuleWriter) writeI64LE(v int64) {
	mw.writeU64LE(uint64(v))
}

func (mw *ModuleWriter) writeI64ArrLE(v []int64) {
	for i := range len(v) {
		mw.writeI64LE(v[i])
	}
}

func (mw *ModuleWriter) writeVariableLEI32(v int32, size int32) {
	o := v
	for range size {
		mw.arr = append(mw.arr, uint8(o&0xFF))
		o >>= 8
	}
}

func (mw *ModuleWriter) writeVariableLEI64(v int64, size int32) {
	o := v
	for range size {
		mw.arr = append(mw.arr, uint8(o&0xFF))
		o >>= 8
	}
}
