package bcio

import (
	"focal-lang/internal/bytecode/constants"
	"focal-lang/internal/bytecode/spec"
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
	mw.writeConstantPool(mw.module.GetConstantPool())
	mw.writeU32LE(uint32(mw.module.GetMainFunctionIndex()))
	mw.writeU32LE(uint32(len(mw.module.GetFunctions())))
	for _, v := range mw.module.GetFunctions() {
		mw.writeFunction(v)
	}
}

func (mw *ModuleWriter) writeConstantPool(pool *constants.ConstantPool) {
	constantCount := pool.Size()
	mw.writeU32LE(uint32(constantCount))
	for i := range constantCount {
		constant := pool.GetConstant(i)
		mw.writeU8LE(uint8(constant.GetTag()))
		switch con := constant.(type) {
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
	mw.writeU32LE(uint32(fun.GetNameIndex()))
	mw.writeU8LE(fun.GetModifier())
	mw.writeU32LE(uint32(len(fun.GetCode())))
	mw.writeU8ArrLE(fun.GetCode())
}

func (bw *ModuleWriter) writeU8LE(v uint8) {
	bw.arr = append(bw.arr, v)
}

func (bw *ModuleWriter) writeU8ArrLE(v []uint8) {
	bw.arr = append(bw.arr, v...)
}

func (bw *ModuleWriter) writeU16LE(v uint16) {
	bw.arr = append(append(bw.arr, uint8(v&0xFF)), uint8((v&0xFF00)>>8))
}

func (bw *ModuleWriter) writeU16ArrLE(v []uint16) {
	for i := range len(v) {
		bw.writeU16LE(v[i])
	}
}

func (bw *ModuleWriter) writeU32LE(v uint32) {
	bw.writeU16LE(uint16(v & 0xFFFF))
	bw.writeU16LE(uint16((v & 0xFFFF0000) >> 16))
}

func (bw *ModuleWriter) writeU32ArrLE(v []uint32) {
	for i := range len(v) {
		bw.writeU32LE(v[i])
	}
}

func (bw *ModuleWriter) writeU64LE(v uint64) {
	bw.writeU32LE(uint32(v & 0xFFFFFFFF))
	bw.writeU32LE(uint32((v & 0xFFFFFFFF00000000) >> 32))
}

func (bw *ModuleWriter) writeU64ArrLE(v []uint64) {
	for i := range len(v) {
		bw.writeU64LE(v[i])
	}
}

func (bw *ModuleWriter) writeVaribleLE32(v uint32, size int32) {
	o := v
	for range size {
		bw.arr = append(bw.arr, uint8(o&0xFF))
		o >>= 8
	}
}

func (bw *ModuleWriter) writeVaribleLE64(v uint64, size int32) {
	o := v
	for range size {
		bw.arr = append(bw.arr, uint8(o&0xFF))
		o >>= 8
	}
}
