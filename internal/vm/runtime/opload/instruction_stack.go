package opload

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/public/runtimeapi"
)

func install_stack_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_DUP] = execDUP
	opcodeMap[opcodes.OP_SWP] = execSWP
	opcodeMap[opcodes.OP_POP] = execPOP
}

func execDUP(vm runtimeapi.VM, frame runtimeapi.Frame) {
	top := vm.GetValueStack().GetTop()
	vm.GetValueStack().Push(top)
}

func execSWP(vm runtimeapi.VM, frame runtimeapi.Frame) {
	first := vm.GetValueStack().Pop()
	last := vm.GetValueStack().Pop()
	vm.GetValueStack().Push(first)
	vm.GetValueStack().Push(last)
}

func execPOP(vm runtimeapi.VM, frame runtimeapi.Frame) {
	vm.GetValueStack().Pop()
}
