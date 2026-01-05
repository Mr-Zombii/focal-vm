package opbool

import (
	"focal-vm/public/runtimeapi"
)

func Install_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	Install_bool_logical_instructions(opcodeMap)
}
