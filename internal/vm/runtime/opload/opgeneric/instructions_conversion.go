package opgeneric

import (
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/public/runtimeapi"
)

func Install_conversion_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_CONV_TO_I8] = _conv_instruction[int8](bctypes.I8, func(pool *rtvalue.RTValuePool, i int8) rtvalue.RTValue {
		return pool.GetOrMakeRTValueI8(i)
	})
	opcodeMap[opcodes.OP_CONV_TO_I16] = _conv_instruction[int16](bctypes.I16, func(pool *rtvalue.RTValuePool, i int16) rtvalue.RTValue {
		return pool.GetOrMakeRTValueI16(i)
	})
	opcodeMap[opcodes.OP_CONV_TO_I32] = _conv_instruction[int32](bctypes.I32, func(pool *rtvalue.RTValuePool, i int32) rtvalue.RTValue {
		return pool.GetOrMakeRTValueI32(i)
	})
	opcodeMap[opcodes.OP_CONV_TO_I64] = _conv_instruction[int64](bctypes.I64, func(pool *rtvalue.RTValuePool, i int64) rtvalue.RTValue {
		return pool.GetOrMakeRTValueI64(i)
	})
	opcodeMap[opcodes.OP_CONV_TO_F32] = _conv_instruction[float32](bctypes.F32, func(pool *rtvalue.RTValuePool, f float32) rtvalue.RTValue {
		return pool.GetOrMakeRTValueF32(f)
	})
	opcodeMap[opcodes.OP_CONV_TO_F64] = _conv_instruction[float64](bctypes.F64, func(pool *rtvalue.RTValuePool, f float64) rtvalue.RTValue {
		return pool.GetOrMakeRTValueF64(f)
	})
}
