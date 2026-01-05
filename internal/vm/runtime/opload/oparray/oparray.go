package oparray

import "focal-vm/public/runtimeapi"

func Install_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	Install_array_instructions(opcodeMap)
}
