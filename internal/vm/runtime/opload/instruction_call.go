package opload

import (
	"fmt"
	"focal-vm/internal/bytecode/constants"
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/util"
	"focal-vm/internal/vm/runtime"
	"focal-vm/public/runtimeapi"
)

func installCall(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_CALL] = execCALL
	opcodeMap[opcodes.OP_FLOAD] = execFLOAD
}

func execFLOAD(vm runtimeapi.VM, frame runtimeapi.Frame) {
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

func execTCALL(vm runtimeapi.VM, frame runtimeapi.Frame) {
	fn := vm.GetStack().PopValue().(runtimeapi.CallableValue)
	if bf, ok := fn.(*runtime.FunctionValue); ok {
		frame.LoadFn(bf.GetFunction())
		return
	}
	vm.Panic(fmt.Sprintf("Expected Focal Function, not Native or FFI/Plugin function!, module: %s, function: %s", frame.GetModuleName(), frame.GetFunctionName()))
}

func execCALL(vm runtimeapi.VM, frame runtimeapi.Frame) {
	fn := vm.GetStack().PopValue().(runtimeapi.CallableValue)
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
