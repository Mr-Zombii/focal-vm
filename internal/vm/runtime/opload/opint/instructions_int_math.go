package opint

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/public/runtimeapi"
)

func Install_int_arithmatic_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_IADD] = _instruction_iadd
	opcodeMap[opcodes.OP_ISUB] = _instruction_isub
	opcodeMap[opcodes.OP_IDIV] = _instruction_idiv
	opcodeMap[opcodes.OP_IMUL] = _instruction_imul
	opcodeMap[opcodes.OP_IMOD] = _instruction_imod
}

func _instruction_iadd(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_int_instruction(vm, true, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueI8(a + b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueI16(a + b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueI32(a + b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueI64(a + b))
	})
}

func _instruction_isub(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_int_instruction(vm, true, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueI8(a - b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueI16(a - b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueI32(a - b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueI64(a - b))
	})
}

func _instruction_imul(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_int_instruction(vm, true, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueI8(a * b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueI16(a * b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueI32(a * b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueI64(a * b))
	})
}

func _instruction_idiv(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_int_instruction(vm, true, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueI8(a / b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueI16(a / b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueI32(a / b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueI64(a / b))
	})
}

func _instruction_imod(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_int_instruction(vm, true, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueI8(a % b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueI16(a % b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueI32(a % b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueI64(a % b))
	})
}
