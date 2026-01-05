package opbool

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/public/runtimeapi"
)

func Install_bool_logical_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_LNOT] = _instruction_lnot
	opcodeMap[opcodes.OP_LOR] = _instruction_lor
	opcodeMap[opcodes.OP_LAND] = _instruction_land
	opcodeMap[opcodes.OP_LXOR] = _instruction_lxor
}

func _instruction_lnot(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	aValue := stack.Pop()
	CheckBool(vm, aValue)

	a := aValue.GetRawValue().(bool)
	stack.Push(ToBoolValue(!a))
}

func _instruction_lor(vm runtimeapi.VM, frame runtimeapi.Frame) {
	_bool_instruction(vm, func(a, b bool) bool {
		return a || b
	})
}

func _instruction_land(vm runtimeapi.VM, frame runtimeapi.Frame) {
	_bool_instruction(vm, func(a, b bool) bool {
		return a && b
	})
}

func _instruction_lxor(vm runtimeapi.VM, frame runtimeapi.Frame) {
	_bool_instruction(vm, func(a, b bool) bool {
		return (a && !b) || (!a && b)
	})
}
