package opmath

import (
	"focal-vm/public/runtimeapi"
)

func Install_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	Install_arithmatic_instructions(opcodeMap)
	Install_relational_instructions(opcodeMap)
}
