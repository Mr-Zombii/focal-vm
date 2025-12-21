package opload

import (
	"focal-lang/internal/bytecode/opcodes"
	"focal-lang/internal/vm/api"
)

func installStack(opcodeMap []api.OpcodeImpl) {
	opcodeMap[opcodes.OP_DUP] = execDUP
	opcodeMap[opcodes.OP_SWP] = execSWP
	opcodeMap[opcodes.OP_POP] = execPOP
}

func execDUP(vm api.VM, frame api.Frame) {
	top := vm.GetStack().GetTopValue()
	vm.GetStack().PushValue(top)
}

func execSWP(vm api.VM, frame api.Frame) {
	first := vm.GetStack().PopValue()
	last := vm.GetStack().PopValue()
	vm.GetStack().PushValue(first)
	vm.GetStack().PushValue(last)
}

func execPOP(vm api.VM, frame api.Frame) {
	vm.GetStack().PopValue()
}
