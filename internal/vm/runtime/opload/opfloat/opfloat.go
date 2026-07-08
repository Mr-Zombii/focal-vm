package opfloat

import (
	"focal-vm/public/runtimeapi"
)

func Install_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	Install_float_arithmatic_instructions(opcodeMap)
	Install_float_relational_instructions(opcodeMap)
}
