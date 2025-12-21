package opload

import (
	"focal-lang/internal/bytecode/constants"
	"focal-lang/internal/bytecode/opcodes"
	"focal-lang/internal/util"
	"focal-lang/internal/vm/api"
	"focal-lang/internal/vm/runtime"
)

func installCall(opcodeMap []api.OpcodeImpl) {
	opcodeMap[opcodes.OP_CALL] = execCALL
	opcodeMap[opcodes.OP_FLOAD] = execFLOAD
}

func execFLOAD(vm api.VM, frame api.Frame) {
	ptr := frame.GetPtr()
	code := *frame.GetCode()

	flags := util.ReadU8LE(code, ptr)
	ptr++

	widthA := int32(flags&0x3) + 1
	widthB := int32((flags>>2)&0x3) + 1

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

	fnValue := runtime.NewFunction(vm.GetScope(), fn)
	vm.GetStack().PushValue(fnValue)

	//callFrame := runtime.NewFrame(vm.GetScope(), mod, fn)
	//callStack := vm.GetCallStack()
	//callStack.PushFrame(callFrame)
}

func execCALL(vm api.VM, frame api.Frame) {
	fn := vm.GetStack().PopValue().(api.CallableValue)
	fn.Call(vm)
	//ptr := frame.GetPtr()
	//code := *frame.GetCode()

	//flags := util.ReadU8LE(code, ptr)
	//ptr++

	//widthA := int32(flags&0x3) + 1
	//widthB := int32((flags>>2)&0x3) + 1

	//modNameId := util.ReadVariableLE32(code, ptr, widthA)
	//ptr += widthA
	//funNameId := util.ReadVariableLE32(code, ptr, widthB)
	//ptr += widthB
	//frame.SetPtr(ptr)

	//modName := frame.GetConstantPool().GetConstant(int32(modNameId))
	//funName := frame.GetConstantPool().GetConstant(int32(funNameId))

	//modNameStr := modName.(*constants.ConstantUTF8String).GetValue()
	//funNameStr := funName.(*constants.ConstantUTF8String).GetValue()

	//mod := vm.LoadModule(modNameStr)
	//fn := mod.GetFunction(funNameStr)

	//callFrame := runtime.NewFrame(vm.GetScope(), mod, fn)
	//callStack := vm.GetCallStack()

	//callStack.PushFrame(callFrame)
}
