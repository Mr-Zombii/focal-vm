package opgeneric

import (
	"focal-vm/public/runtimeapi"
)

func Install_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	Install_equality_instructions(opcodeMap)
	Install_conversion_instructions(opcodeMap)
}
