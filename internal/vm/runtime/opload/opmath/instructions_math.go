package opmath

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/public/runtimeapi"
	"math"
)

func Install_arithmatic_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_ADD] = _instruction_add
	opcodeMap[opcodes.OP_SUB] = _instruction_sub
	opcodeMap[opcodes.OP_DIV] = _instruction_div
	opcodeMap[opcodes.OP_MUL] = _instruction_mul
	opcodeMap[opcodes.OP_MOD] = _instruction_mod
}

func _instruction_add(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_number_instruction(vm, true, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueI8(a + b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueI16(a + b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueI32(a + b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueI64(a + b))
	}, func(a, b float32) {
		stack.Push(rtpool.GetOrMakeRTValueF32(a + b))
	}, func(a, b float64) {
		stack.Push(rtpool.GetOrMakeRTValueF64(a + b))
	})
}

func _instruction_sub(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_number_instruction(vm, true, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueI8(a - b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueI16(a - b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueI32(a - b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueI64(a - b))
	}, func(a, b float32) {
		stack.Push(rtpool.GetOrMakeRTValueF32(a - b))
	}, func(a, b float64) {
		stack.Push(rtpool.GetOrMakeRTValueF64(a - b))
	})
}

func _instruction_mul(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_number_instruction(vm, true, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueI8(a * b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueI16(a * b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueI32(a * b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueI64(a * b))
	}, func(a, b float32) {
		stack.Push(rtpool.GetOrMakeRTValueF32(a * b))
	}, func(a, b float64) {
		stack.Push(rtpool.GetOrMakeRTValueF64(a * b))
	})
}

func _instruction_div(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_number_instruction(vm, true, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueI8(a / b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueI16(a / b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueI32(a / b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueI64(a / b))
	}, func(a, b float32) {
		stack.Push(rtpool.GetOrMakeRTValueF32(a / b))
	}, func(a, b float64) {
		stack.Push(rtpool.GetOrMakeRTValueF64(a / b))
	})
}

func _instruction_mod(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_number_instruction(vm, true, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueI8(a % b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueI16(a % b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueI32(a % b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueI64(a % b))
	}, func(a, b float32) {
		stack.Push(rtpool.GetOrMakeRTValueF32(float32(math.Mod(float64(a), float64(b)))))
	}, func(a, b float64) {
		stack.Push(rtpool.GetOrMakeRTValueF64(math.Mod(a, b)))
	})
}
