package opmath

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/public/runtimeapi"
)

func Install_relational_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_LT] = _instruction_lt
	opcodeMap[opcodes.OP_GT] = _instruction_gt
	opcodeMap[opcodes.OP_LE] = _instruction_le
	opcodeMap[opcodes.OP_GE] = _instruction_ge
}

func _instruction_lt(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_number_instruction(vm, true, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a < b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a < b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a < b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a < b))
	}, func(a, b float32) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a < b))
	}, func(a, b float64) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a < b))
	})
}

func _instruction_gt(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_number_instruction(vm, true, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a > b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a > b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a > b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a > b))
	}, func(a, b float32) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a > b))
	}, func(a, b float64) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a > b))
	})
}

func _instruction_le(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_number_instruction(vm, true, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a <= b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a <= b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a <= b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a <= b))
	}, func(a, b float32) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a <= b))
	}, func(a, b float64) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a <= b))
	})
}

func _instruction_ge(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_number_instruction(vm, true, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a >= b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a >= b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a >= b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a >= b))
	}, func(a, b float32) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a >= b))
	}, func(a, b float64) {
		stack.Push(rtpool.GetOrMakeRTValueBool(a >= b))
	})
}
