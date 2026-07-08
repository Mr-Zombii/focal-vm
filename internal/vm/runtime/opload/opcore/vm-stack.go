package opcore

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/erroring"
	"focal-vm/public/runtimeapi"
)

func Install_stack_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_DUP] = _instruction_dup
	opcodeMap[opcodes.OP_SWP] = _instruction_swp
	opcodeMap[opcodes.OP_POP] = _instruction_pop
}

func _instruction_dup(vm runtimeapi.VM, _ runtimeapi.Frame) {
	top := vm.GetValueStack().GetTop()
	if top == nil {
		erroring.GlobalErrorHandler.Panic("Instruction Dup", "Cannot duplication nil value or empty stack")
		return
	}
	top.IncRefCount()
	vm.GetValueStack().Push(top)
}

func _instruction_swp(vm runtimeapi.VM, _ runtimeapi.Frame) {
	first := vm.GetValueStack().Pop()
	last := vm.GetValueStack().Pop()
	vm.GetValueStack().Push(first)
	vm.GetValueStack().Push(last)
}

func _instruction_pop(vm runtimeapi.VM, _ runtimeapi.Frame) {
	vm.GetValueStack().Pop().DecRefCount()
}
