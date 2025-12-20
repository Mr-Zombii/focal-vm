package runtime

import (
	"fmt"
	"focal-lang/internal/bytecode/constants"
	"focal-lang/internal/bytecode/opcodes"
	"focal-lang/internal/util"
	"focal-lang/internal/vm/api"
)

func InstallOpcodes(vm api.VM) {
	opcodeMap := make([]api.OpcodeImpl, opcodes.OPCODE_COUNT)
	opcodeMap[opcodes.OP_CLOAD1] = makeLoadConstant(1)
	opcodeMap[opcodes.OP_CLOAD2] = makeLoadConstant(2)
	opcodeMap[opcodes.OP_CLOAD3] = makeLoadConstant(3)
	opcodeMap[opcodes.OP_CLOAD4] = makeLoadConstant(4)
	opcodeMap[opcodes.OP_PRINT] = execPrint
	opcodeMap[opcodes.OP_CALL11] = makeCall(1, 1)
	opcodeMap[opcodes.OP_CALL12] = makeCall(1, 2)
	opcodeMap[opcodes.OP_CALL13] = makeCall(1, 3)
	opcodeMap[opcodes.OP_CALL14] = makeCall(1, 4)
	opcodeMap[opcodes.OP_CALL21] = makeCall(2, 1)
	opcodeMap[opcodes.OP_CALL22] = makeCall(2, 2)
	opcodeMap[opcodes.OP_CALL23] = makeCall(2, 3)
	opcodeMap[opcodes.OP_CALL24] = makeCall(2, 4)
	opcodeMap[opcodes.OP_CALL31] = makeCall(3, 1)
	opcodeMap[opcodes.OP_CALL32] = makeCall(3, 2)
	opcodeMap[opcodes.OP_CALL33] = makeCall(3, 3)
	opcodeMap[opcodes.OP_CALL34] = makeCall(3, 4)
	opcodeMap[opcodes.OP_CALL41] = makeCall(4, 1)
	opcodeMap[opcodes.OP_CALL42] = makeCall(4, 2)
	opcodeMap[opcodes.OP_CALL43] = makeCall(4, 3)
	opcodeMap[opcodes.OP_CALL44] = makeCall(4, 4)
	opcodeMap[opcodes.OP_RET] = execRet
	vm.InstallOpcodeMap(opcodeMap)
}

func makeLoadConstant(width int32) api.OpcodeImpl {
	return func(v api.VM, f api.Frame) {
		execLoadConstant(v, f, width)
	}
}

func execLoadConstant(vm api.VM, frame api.Frame, width int32) {
	ptr := frame.GetPtr()
	code := *frame.GetCode()

	constantIdx := util.ReadVariableLE32(code, ptr, width)
	constant := frame.GetConstantPool().GetConstant(int32(constantIdx))
	ptr += width
	frame.SetPtr(ptr)

	vm.GetStack().PushConst(constant)
}

func makeCall(widthA int32, widthB int32) api.OpcodeImpl {
	return func(v api.VM, f api.Frame) {
		execCall(v, f, widthA, widthB)
	}
}

func execCall(vm api.VM, frame api.Frame, widthA int32, widthB int32) {
	ptr := frame.GetPtr()
	code := *frame.GetCode()

	modNameId := util.ReadVariableLE32(code, ptr, widthA)
	ptr += widthA
	funNameId := util.ReadVariableLE32(code, ptr, widthB)
	ptr += widthB
	frame.SetPtr(ptr)

	modName := frame.GetConstantPool().GetConstant(int32(modNameId))
	funName := frame.GetConstantPool().GetConstant(int32(funNameId))

	modNameStr := modName.(*constants.ConstantUTF8String).GetValue()
	funNameStr := funName.(*constants.ConstantUTF8String).GetValue()

	mod := vm.LoadModule(modNameStr)
	fn := mod.GetFunction(funNameStr)

	callFrame := NewFrame(mod, fn)
	callStack := vm.GetCallStack()

	callStack.PushFrame(callFrame)
}

func execPrint(vm api.VM, frame api.Frame) {
	value := vm.GetStack().PopValue()
	fmt.Println(value)
}

func execRet(vm api.VM, frame api.Frame) {
	vm.GetCallStack().PopFrame()
}
