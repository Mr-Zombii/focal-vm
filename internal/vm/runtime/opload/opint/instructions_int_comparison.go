package opint

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/vm/runtime"
	"focal-vm/public/runtimeapi"
)

func Install_int_comparison_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_ILT] = _instruction_ilt
	opcodeMap[opcodes.OP_IGT] = _instruction_igt
	opcodeMap[opcodes.OP_ILE] = _instruction_ile
	opcodeMap[opcodes.OP_IGE] = _instruction_ige
}

func _instruction_ilt(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_int_instruction(vm, true, func(a, b int8) {
		stack.Push(runtime.NewBooleanValue(a < b))
	}, func(a, b int16) {
		stack.Push(runtime.NewBooleanValue(a < b))
	}, func(a, b int32) {
		stack.Push(runtime.NewBooleanValue(a < b))
	}, func(a, b int64) {
		stack.Push(runtime.NewBooleanValue(a < b))
	})
}

func _instruction_igt(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_int_instruction(vm, true, func(a, b int8) {
		stack.Push(runtime.NewBooleanValue(a > b))
	}, func(a, b int16) {
		stack.Push(runtime.NewBooleanValue(a > b))
	}, func(a, b int32) {
		stack.Push(runtime.NewBooleanValue(a > b))
	}, func(a, b int64) {
		stack.Push(runtime.NewBooleanValue(a > b))
	})
}

func _instruction_ile(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_int_instruction(vm, true, func(a, b int8) {
		stack.Push(runtime.NewBooleanValue(a <= b))
	}, func(a, b int16) {
		stack.Push(runtime.NewBooleanValue(a <= b))
	}, func(a, b int32) {
		stack.Push(runtime.NewBooleanValue(a <= b))
	}, func(a, b int64) {
		stack.Push(runtime.NewBooleanValue(a <= b))
	})
}

func _instruction_ige(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_int_instruction(vm, true, func(a, b int8) {
		stack.Push(runtime.NewBooleanValue(a >= b))
	}, func(a, b int16) {
		stack.Push(runtime.NewBooleanValue(a >= b))
	}, func(a, b int32) {
		stack.Push(runtime.NewBooleanValue(a >= b))
	}, func(a, b int64) {
		stack.Push(runtime.NewBooleanValue(a >= b))
	})
}
