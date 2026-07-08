package opfloat

import (
	"focal-vm/internal/bytecode/opcodes"
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
	rtpool := vm.GetRTValuePool()

	_float_instruction(vm, func(a, b float32) {
		stack.Push(rtpool.GetOrMakeRTValueF32(a + b))
	}, func(a, b float64) {
		stack.Push(rtpool.GetOrMakeRTValueF64(a + b))
	})
}

func _instruction_fsub(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_float_instruction(vm, func(a, b float32) {
		stack.Push(rtpool.GetOrMakeRTValueF32(a - b))
	}, func(a, b float64) {
		stack.Push(rtpool.GetOrMakeRTValueF64(a - b))
	})
}

func _instruction_fmul(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_float_instruction(vm, func(a, b float32) {
		stack.Push(rtpool.GetOrMakeRTValueF32(a * b))
	}, func(a, b float64) {
		stack.Push(rtpool.GetOrMakeRTValueF64(a * b))
	})
}

func _instruction_fdiv(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_float_instruction(vm, func(a, b float32) {
		stack.Push(rtpool.GetOrMakeRTValueF32(a / b))
	}, func(a, b float64) {
		stack.Push(rtpool.GetOrMakeRTValueF64(a / b))
	})
}
