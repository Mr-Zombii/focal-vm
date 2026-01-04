package opload

import (
	"focal-vm/internal/bytecode/constants"
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/util"
	"focal-vm/public/runtimeapi"
)

func installScope(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_CLOAD] = execCLOAD

	opcodeMap[opcodes.OP_VLOAD] = execVLOAD
	opcodeMap[opcodes.OP_VSTORE] = execVSTORE
	opcodeMap[opcodes.OP_DGLOBAL] = execDGLOBAL
	opcodeMap[opcodes.OP_DLOCAL] = execDLOCAL
}

func execDGLOBAL(vm runtimeapi.VM, frame runtimeapi.Frame) {
	defineLocal(vm.GetScope(), frame)
}

func execDLOCAL(vm runtimeapi.VM, frame runtimeapi.Frame) {
	defineLocal(frame.GetScope(), frame)
}

func defineLocal(scope runtimeapi.Scope, frame runtimeapi.Frame) {
	ptr := frame.GetPtr()
	code := *frame.GetCode()

	flags := util.ReadU8LE(code, ptr)
	ptr++

	width := int32(flags&0x3) + 1

	constantIdx := util.ReadVariableLE32(code, ptr, width)
	constant := frame.GetConstantPool().GetConstant(int32(constantIdx))
	ptr += width
	frame.SetPtr(ptr)

	localName := constant.(*constants.ConstantUTF8String).GetValue()
	scope.DefineLocal(localName)
}

func execCLOAD(vm runtimeapi.VM, frame runtimeapi.Frame) {
	ptr := frame.GetPtr()
	code := *frame.GetCode()

	flags := util.ReadU8LE(code, ptr)
	ptr++

	width := int32(flags&0x3) + 1

	constantIdx := util.ReadVariableLE32(code, ptr, width)
	constant := frame.GetConstantPool().GetConstant(int32(constantIdx))
	ptr += width
	frame.SetPtr(ptr)

	vm.GetStack().PushConst(constant)
}

func execVLOAD(vm runtimeapi.VM, frame runtimeapi.Frame) {
	ptr := frame.GetPtr()
	code := *frame.GetCode()

	flags := util.ReadU8LE(code, ptr)
	ptr++

	width := int32(flags&0x3) + 1

	constantIdx := util.ReadVariableLE32(code, ptr, width)
	constant := frame.GetConstantPool().GetConstant(int32(constantIdx))
	ptr += width
	frame.SetPtr(ptr)

	localName := constant.(*constants.ConstantUTF8String).GetValue()
	vm.GetStack().PushValue(frame.GetScope().GetLocal(localName))
}

func execVSTORE(vm runtimeapi.VM, frame runtimeapi.Frame) {
	ptr := frame.GetPtr()
	code := *frame.GetCode()

	flags := util.ReadU8LE(code, ptr)
	ptr++

	width := int32(flags&0x3) + 1

	constantIdx := util.ReadVariableLE32(code, ptr, width)
	constant := frame.GetConstantPool().GetConstant(int32(constantIdx))
	ptr += width
	frame.SetPtr(ptr)

	localName := constant.(*constants.ConstantUTF8String).GetValue()
	frame.GetScope().SetLocal(localName, vm.GetStack().PopValue())
}
