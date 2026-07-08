package opload

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/vm/runtime/opload/oparray"
	"focal-vm/internal/vm/runtime/opload/opbool"
	"focal-vm/internal/vm/runtime/opload/opcore"
	"focal-vm/internal/vm/runtime/opload/opfloat"
	"focal-vm/internal/vm/runtime/opload/opgeneric"
	"focal-vm/internal/vm/runtime/opload/opint"
	"focal-vm/internal/vm/runtime/opload/opobject"
	"focal-vm/internal/vm/runtime/opload/opscope"
	"focal-vm/internal/vm/runtime/opload/opstruct"
	"focal-vm/public/runtimeapi"
)

func InstallOpcodes(vm runtimeapi.VM) {
	opcodeMap := make([]runtimeapi.OpcodeImpl, opcodes.OPCODE_COUNT)
	//install_call_instructions(opcodeMap)

	opstruct.Install_instructions(opcodeMap)
	opobject.Install_instructions(opcodeMap)
	opcore.Install_instructions(opcodeMap)
	oparray.Install_instructions(opcodeMap)
	opbool.Install_instructions(opcodeMap)
	opfloat.Install_instructions(opcodeMap)
	opgeneric.Install_instructions(opcodeMap)
	opint.Install_instructions(opcodeMap)
	opobject.Install_instructions(opcodeMap)
	opscope.Install_instructions(opcodeMap)

	vm.InstallOpcodeMap(opcodeMap)
}
