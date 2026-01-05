package opint

import (
	"focal-vm/public/runtimeapi"
)

func Install_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	Install_int_arithmatic_instructions(opcodeMap)
	Install_int_comparison_instructions(opcodeMap)
	Install_int_bitwise_instructions(opcodeMap)
}
