package opload

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/public/runtimeapi"
)

func installStack(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_DUP] = execDUP
	opcodeMap[opcodes.OP_SWP] = execSWP
	opcodeMap[opcodes.OP_POP] = execPOP
}

func execDUP(vm runtimeapi.VM, frame runtimeapi.Frame) {
	top := vm.GetStack().GetTopValue()
	vm.GetStack().PushValue(top)
}

func execSWP(vm runtimeapi.VM, frame runtimeapi.Frame) {
	first := vm.GetStack().PopValue()
	last := vm.GetStack().PopValue()
	vm.GetStack().PushValue(first)
	vm.GetStack().PushValue(last)
}

func execPOP(vm runtimeapi.VM, frame runtimeapi.Frame) {
	vm.GetStack().PopValue()
}
