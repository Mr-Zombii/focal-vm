package opcore

import "focal-vm/public/runtimeapi"

func Install_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	Install_control_flow(opcodeMap)
	Install_load_instructions(opcodeMap)
	Install_stack_instructions(opcodeMap)
	Install_call_instructions(opcodeMap)
}
