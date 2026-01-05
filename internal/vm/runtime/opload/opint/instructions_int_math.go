package opint

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/vm/runtime"
	"focal-vm/public/runtimeapi"
)

func Install_int_arithmatic_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_IADD] = _instruction_iadd
	opcodeMap[opcodes.OP_ISUB] = _instruction_isub
	opcodeMap[opcodes.OP_IDIV] = _instruction_idiv
	opcodeMap[opcodes.OP_IMUL] = _instruction_imul
	opcodeMap[opcodes.OP_IMOD] = _instruction_imod
}

func _instruction_iadd(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_int_instruction(vm, true, func(a, b int8) {
		stack.Push(runtime.NewInt8Value(a + b))
	}, func(a, b int16) {
		stack.Push(runtime.NewInt16Value(a + b))
	}, func(a, b int32) {
		stack.Push(runtime.NewInt32Value(a + b))
	}, func(a, b int64) {
		stack.Push(runtime.NewInt64Value(a + b))
	})
}

func _instruction_isub(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_int_instruction(vm, true, func(a, b int8) {
		stack.Push(runtime.NewInt8Value(a - b))
	}, func(a, b int16) {
		stack.Push(runtime.NewInt16Value(a - b))
	}, func(a, b int32) {
		stack.Push(runtime.NewInt32Value(a - b))
	}, func(a, b int64) {
		stack.Push(runtime.NewInt64Value(a - b))
	})
}

func _instruction_imul(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_int_instruction(vm, true, func(a, b int8) {
		stack.Push(runtime.NewInt8Value(a * b))
	}, func(a, b int16) {
		stack.Push(runtime.NewInt16Value(a * b))
	}, func(a, b int32) {
		stack.Push(runtime.NewInt32Value(a * b))
	}, func(a, b int64) {
		stack.Push(runtime.NewInt64Value(a * b))
	})
}

func _instruction_idiv(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_int_instruction(vm, true, func(a, b int8) {
		stack.Push(runtime.NewInt8Value(a / b))
	}, func(a, b int16) {
		stack.Push(runtime.NewInt16Value(a / b))
	}, func(a, b int32) {
		stack.Push(runtime.NewInt32Value(a / b))
	}, func(a, b int64) {
		stack.Push(runtime.NewInt64Value(a / b))
	})
}

func _instruction_imod(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_int_instruction(vm, true, func(a, b int8) {
		stack.Push(runtime.NewInt8Value(a % b))
	}, func(a, b int16) {
		stack.Push(runtime.NewInt16Value(a % b))
	}, func(a, b int32) {
		stack.Push(runtime.NewInt32Value(a % b))
	}, func(a, b int64) {
		stack.Push(runtime.NewInt64Value(a % b))
	})
}
