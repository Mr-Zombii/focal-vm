package opfloat

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/vm/runtime/opload/opbool"
	"focal-vm/public/runtimeapi"
)

func Install_float_comparison_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_FLT] = _instruction_flt
	opcodeMap[opcodes.OP_FGT] = _instruction_fgt
	opcodeMap[opcodes.OP_FLE] = _instruction_fle
	opcodeMap[opcodes.OP_FGE] = _instruction_fge
}

/*
[stack-in]:
├─> floatValue A
└─> floatValue B

[stack-out]:
└─> floatValue (takes largest bit-count)
*/
func _instruction_flt(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_float_instruction(vm, func(a, b float32) {
		stack.Push(opbool.ToBoolValue(a < b))
	}, func(a, b float64) {
		stack.Push(opbool.ToBoolValue(a < b))
	})
}

func _instruction_fgt(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_float_instruction(vm, func(a, b float32) {
		stack.Push(opbool.ToBoolValue(a > b))
	}, func(a, b float64) {
		stack.Push(opbool.ToBoolValue(a > b))
	})
}

func _instruction_fle(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_float_instruction(vm, func(a, b float32) {
		stack.Push(opbool.ToBoolValue(a <= b))
	}, func(a, b float64) {
		stack.Push(opbool.ToBoolValue(a <= b))
	})
}

func _instruction_fge(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_float_instruction(vm, func(a, b float32) {
		stack.Push(opbool.ToBoolValue(a >= b))
	}, func(a, b float64) {
		stack.Push(opbool.ToBoolValue(a >= b))
	})
}
