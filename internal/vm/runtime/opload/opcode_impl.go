package opload

import (
	"fmt"
	"focal-lang/internal/bytecode/opcodes"
	"focal-lang/internal/vm/api"
)

func InstallOpcodes(vm api.VM) {
	opcodeMap := make([]api.OpcodeImpl, opcodes.OPCODE_COUNT)
	installScope(opcodeMap)
	installCall(opcodeMap)
	installStack(opcodeMap)

	opcodeMap[opcodes.OP_PRINT] = execPrint
	opcodeMap[opcodes.OP_RET] = execRet
	vm.InstallOpcodeMap(opcodeMap)
}

func execPrint(vm api.VM, frame api.Frame) {
	value := vm.GetStack().PopValue()
	fmt.Println(value)
}

func execRet(vm api.VM, frame api.Frame) {
	vm.GetCallStack().PopFrame()
}
