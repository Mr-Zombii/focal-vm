package opbool

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/public/runtimeapi"
)

func Install_bool_logical_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_LNOT] = _instruction_lnot
	opcodeMap[opcodes.OP_LOR] = _instruction_lor
	opcodeMap[opcodes.OP_LAND] = _instruction_land
	opcodeMap[opcodes.OP_LXOR] = _instruction_lxor
}

func _instruction_lnot(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	aValue := stack.Pop()
	CheckBool(vm, aValue)

	a := aValue.(*rtvalue.RTValueBool).GetValue()
	stack.Push(rtpool.GetOrMakeRTValueBool(!a))
	aValue.DecRefCount()
}

func _instruction_lor(vm runtimeapi.VM, _ runtimeapi.Frame) {
	_bool_instruction(vm, func(a, b bool) bool {
		return a || b
	})
}

func _instruction_land(vm runtimeapi.VM, _ runtimeapi.Frame) {
	_bool_instruction(vm, func(a, b bool) bool {
		return a && b
	})
}

func _instruction_lxor(vm runtimeapi.VM, _ runtimeapi.Frame) {
	_bool_instruction(vm, func(a, b bool) bool {
		return (a && !b) || (!a && b)
	})
}
