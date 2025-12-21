package opload

import (
	"focal-lang/internal/bytecode/constants"
	"focal-lang/internal/bytecode/opcodes"
	"focal-lang/internal/util"
	"focal-lang/internal/vm/api"
)

func installScope(opcodeMap []api.OpcodeImpl) {
	opcodeMap[opcodes.OP_CLOAD] = execCLOAD

	opcodeMap[opcodes.OP_VLOAD] = execVLOAD
	opcodeMap[opcodes.OP_VSTORE] = execVSTORE
	opcodeMap[opcodes.OP_DGLOBAL] = execDGLOBAL
	opcodeMap[opcodes.OP_DLOCAL] = execDLOCAL
}

func execDGLOBAL(vm api.VM, frame api.Frame) {
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
	vm.GetScope().DefineLocal(localName)
}

func execDLOCAL(vm api.VM, frame api.Frame) {
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
	frame.GetScope().DefineLocal(localName)
}

func execCLOAD(vm api.VM, frame api.Frame) {
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

func execVLOAD(vm api.VM, frame api.Frame) {
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

func execVSTORE(vm api.VM, frame api.Frame) {
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
