package opload

import (
	"fmt"
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/vm/runtime"
	"focal-vm/internal/vm/runtime/opload/oparray"
	"focal-vm/internal/vm/runtime/opload/opbool"
	"focal-vm/internal/vm/runtime/opload/opcore"
	"focal-vm/internal/vm/runtime/opload/opfloat"
	"focal-vm/internal/vm/runtime/opload/opgeneric"
	"focal-vm/internal/vm/runtime/opload/opint"
	"focal-vm/internal/vm/runtime/opload/opobject"
	"focal-vm/public/runtimeapi"
)

func InstallOpcodes(vm runtimeapi.VM) {
	opcodeMap := make([]runtimeapi.OpcodeImpl, opcodes.OPCODE_COUNT)
	install_scope_instructions(opcodeMap)
	install_call_instructions(opcodeMap)
	install_stack_instructions(opcodeMap)

	oparray.Install_instructions(opcodeMap)
	opbool.Install_instructions(opcodeMap)
	opcore.Install_instructions(opcodeMap)
	opfloat.Install_instructions(opcodeMap)
	opgeneric.Install_instructions(opcodeMap)
	opint.Install_instructions(opcodeMap)
	opobject.Install_instructions(opcodeMap)

	opcodeMap[opcodes.OP_RET] = execRet
	vm.InstallOpcodeMap(opcodeMap)
}

func checkType(vm runtimeapi.VM, value runtimeapi.Value, vtypes ...runtimeapi.ValueTag) {
	for _, vtype := range vtypes {
		if vtype == value.GetTag() {
			return
		}
	}

	if len(vtypes) == 1 {
		vm.Panic(fmt.Sprintf("Unexpected stack value type, %v, expected type %v", value.GetTag(), vtypes[0]))
		return
	}
	out := fmt.Sprintf("Unexpected stack value type, %v, expected one of the types of [ ", value.GetTag())
	for i, v := range vtypes {
		out += fmt.Sprint(v)
		if i != len(vtypes)-1 {
			out += ", "
		}
	}
	out += " ]"
	vm.Panic(out)
}

func checkInt(vm runtimeapi.VM, value runtimeapi.Value) {
	if runtime.ValueIsInteger(value) {
		vm.Panic(fmt.Sprintf("Stack value should be an integer type, not type %v", value.GetTag()))
	}
}

func execRet(vm runtimeapi.VM, frame runtimeapi.Frame) {
	vm.GetCallStack().Pop()
}
