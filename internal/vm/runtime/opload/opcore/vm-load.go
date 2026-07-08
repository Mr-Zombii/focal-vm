package opcore

import (
	"fmt"
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/bytecode/constants"
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/bytecode/spec"
	"focal-vm/internal/util"
	"focal-vm/public/runtimeapi"
)

func Install_load_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_LOAD_CONST] = _instruction_load_const
	opcodeMap[opcodes.OP_LOAD_STATIC_FUNCTION] = _instruction_load_static_function
}

/*
[stack-in]: (empty)

[stack-out]:
└─> constantValue
*/
func _instruction_load_const(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	cpool := frame.GetConstantPool()
	rtpool := vm.GetRTValuePool()

	ptr := frame.GetPtr()
	code := *frame.GetCode()

	flags := util.ReadU8LE(code, ptr)
	ptr++

	width := int32(flags&0x3) + 1

	constantIdx := util.ReadVariableLEU32(code, ptr, width)
	constant := cpool.GetConstant(int32(constantIdx))
	ptr += width
	frame.SetPtr(ptr)

	value := rtpool.CreateOrGetFromConstant(constant)
	stack.Push(value)
}

func _instruction_load_static_function(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	cpool := frame.GetConstantPool()
	tpool := frame.GetTypePool()
	rtpool := vm.GetRTValuePool()

	ptr := frame.GetPtr()
	code := *frame.GetCode()

	flags := util.ReadU8LE(code, ptr)
	ptr++

	widthA := int32(flags&0x3) + 1
	widthB := int32((flags>>2)&0x3) + 1
	widthC := int32((flags>>4)&0x3) + 1

	modNameId := util.ReadVariableLEI32(code, ptr, widthA)
	ptr += widthA
	funNameId := util.ReadVariableLEI32(code, ptr, widthB)
	ptr += widthB
	typeId := util.ReadVariableLEI32(code, ptr, widthC)
	ptr += widthC
	frame.SetPtr(ptr)

	modName := cpool.ExpectConstant(modNameId, constants.ConstantTagUTF8String).(*constants.ConstantUTF8String).GetValue()
	fnName := cpool.ExpectConstant(funNameId, constants.ConstantTagUTF8String).(*constants.ConstantUTF8String).GetValue()
	fnType := tpool.ExpectType(typeId, bctypes.BCTYPE_FUNCTION).(*bctypes.FunctionType)

	callingModuleName := frame.GetModuleName()

	mod := vm.LoadModule(modName)

	fn, err := mod.GetFunction(fnName, fnType)
	if err != nil {
		vm.Panic(err.Error())
		return
	}

	if callingModuleName != modName && fn.GetModifier()&spec.BCFunctionModPrivate != 0 {
		vm.Panic(fmt.Sprintf("Function access violation, you cannot access the function \"%s\" from module \"%s\", it can only be accessed from its own module \"%s\"!", fnName, callingModuleName, modName))
		return
	}

	//parentScope := vm.GetScope()
	//if fn.GetModifier()&spec.BCFunctionModSubFunc != 0 {
	//	parentScope = frame.GetScope()
	//}

	fnValue := rtpool.CreateVMFunction(fn)
	stack.Push(fnValue)
}
