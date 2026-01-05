package opgeneric

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/vm/runtime"
	"focal-vm/public/runtimeapi"
)

func Install_conversion_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_CONV_TO_I8] = _conv_instruction[int8](runtimeapi.ValueTagInt8, runtime.NewInt8Value)
	opcodeMap[opcodes.OP_CONV_TO_I16] = _conv_instruction[int16](runtimeapi.ValueTagInt8, runtime.NewInt16Value)
	opcodeMap[opcodes.OP_CONV_TO_I32] = _conv_instruction[int32](runtimeapi.ValueTagInt8, runtime.NewInt32Value)
	opcodeMap[opcodes.OP_CONV_TO_I64] = _conv_instruction[int64](runtimeapi.ValueTagInt8, runtime.NewInt64Value)
	opcodeMap[opcodes.OP_CONV_TO_F32] = _conv_instruction[float32](runtimeapi.ValueTagInt8, runtime.NewFloat32Value)
	opcodeMap[opcodes.OP_CONV_TO_F64] = _conv_instruction[float64](runtimeapi.ValueTagInt8, runtime.NewFloat64Value)
}
