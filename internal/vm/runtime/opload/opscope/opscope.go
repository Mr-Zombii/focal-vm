package opscope

import (
	"focal-vm/public/runtimeapi"
)

func Install_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	Install_local_scope_instructions(opcodeMap)
	Install_global_scope_instructions(opcodeMap)
}
