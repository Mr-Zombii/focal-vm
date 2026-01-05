package opgeneric

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/vm/runtime/opload/opbool"
	"focal-vm/public/runtimeapi"
	"reflect"
)

func Install_equality_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_EQ] = _instruction_eq
	opcodeMap[opcodes.OP_NEQ] = _instruction_neq
}

func _instruction_eq(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	aValue := stack.Pop()
	bValue := stack.Pop()

	if aValue.GetTag() != bValue.GetTag() {
		stack.Push(opbool.ToBoolValue(false))
		return
	}

	stack.Push(opbool.ToBoolValue(reflect.DeepEqual(aValue.GetRawValue(), bValue.GetRawValue())))
}

func _instruction_neq(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	aValue := stack.Pop()
	bValue := stack.Pop()

	if aValue.GetTag() == bValue.GetTag() {
		stack.Push(opbool.ToBoolValue(true))
		return
	}

	stack.Push(opbool.ToBoolValue(!reflect.DeepEqual(aValue.GetRawValue(), bValue.GetRawValue())))
}
