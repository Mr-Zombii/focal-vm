package opcore

import "focal-vm/public/runtimeapi"

func Install_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	Install_control_flow(opcodeMap)
}
