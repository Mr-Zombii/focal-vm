package opload

import (
	"fmt"
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/public/runtimeapi"
)

func InstallOpcodes(vm runtimeapi.VM) {
	opcodeMap := make([]runtimeapi.OpcodeImpl, opcodes.OPCODE_COUNT)
	installScope(opcodeMap)
	installCall(opcodeMap)
	installStack(opcodeMap)

	opcodeMap[opcodes.OP_PRINT] = execPrint
	opcodeMap[opcodes.OP_RET] = execRet
	vm.InstallOpcodeMap(opcodeMap)
}

func execPrint(vm runtimeapi.VM, frame runtimeapi.Frame) {
	value := vm.GetStack().PopValue()
	fmt.Println(value)
}

func execRet(vm runtimeapi.VM, frame runtimeapi.Frame) {
	vm.GetCallStack().PopFrame()
}
