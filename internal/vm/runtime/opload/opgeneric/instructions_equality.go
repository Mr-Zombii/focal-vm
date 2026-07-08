package opgeneric

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/public/runtimeapi"
	"reflect"
)

func Install_equality_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_EQ] = _instruction_eq
	opcodeMap[opcodes.OP_NEQ] = _instruction_neq
}

func _instruction_eq(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	aValue := stack.Pop()
	bValue := stack.Pop()

	if aValue.GetType() != bValue.GetType() || aValue.GetTag() != bValue.GetTag() {
		stack.Push(rtpool.GetOrMakeRTValueBool(false))
		aValue.DecRefCount()
		bValue.DecRefCount()
		return
	}

	stack.Push(rtpool.GetOrMakeRTValueBool(reflect.DeepEqual(aValue, bValue)))
	aValue.DecRefCount()
	bValue.DecRefCount()
}

func _instruction_neq(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	aValue := stack.Pop()
	bValue := stack.Pop()

	if aValue.GetType() != bValue.GetType() || aValue.GetTag() != bValue.GetTag() {
		stack.Push(rtpool.GetOrMakeRTValueBool(true))
		aValue.DecRefCount()
		bValue.DecRefCount()
		return
	}

	stack.Push(rtpool.GetOrMakeRTValueBool(!reflect.DeepEqual(aValue, bValue)))
	aValue.DecRefCount()
	bValue.DecRefCount()
}
