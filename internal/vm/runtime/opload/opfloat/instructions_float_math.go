package opfloat

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/vm/runtime"
	"focal-vm/public/runtimeapi"
)

func Install_float_arithmatic_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_FADD] = _instruction_fadd
	opcodeMap[opcodes.OP_FSUB] = _instruction_fsub
	opcodeMap[opcodes.OP_FDIV] = _instruction_fdiv
	opcodeMap[opcodes.OP_FMUL] = _instruction_fmul
}

func _instruction_fadd(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_float_instruction(vm, func(a, b float32) {
		stack.Push(runtime.NewFloat32Value(a + b))
	}, func(a, b float64) {
		stack.Push(runtime.NewFloat64Value(a + b))
	})
}

func _instruction_fsub(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_float_instruction(vm, func(a, b float32) {
		stack.Push(runtime.NewFloat32Value(a - b))
	}, func(a, b float64) {
		stack.Push(runtime.NewFloat64Value(a - b))
	})
}

func _instruction_fmul(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_float_instruction(vm, func(a, b float32) {
		stack.Push(runtime.NewFloat32Value(a * b))
	}, func(a, b float64) {
		stack.Push(runtime.NewFloat64Value(a * b))
	})
}

func _instruction_fdiv(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_float_instruction(vm, func(a, b float32) {
		stack.Push(runtime.NewFloat32Value(a / b))
	}, func(a, b float64) {
		stack.Push(runtime.NewFloat64Value(a / b))
	})
}
