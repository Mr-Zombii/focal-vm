package opload

import (
	"fmt"
	"focal-vm/internal/bytecode/constants"
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/bytecode/spec"
	"focal-vm/internal/util"
	"focal-vm/internal/vm/runtime"
	"focal-vm/public/runtimeapi"
	"reflect"
)

func install_call_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_CALL] = execCALL
	opcodeMap[opcodes.OP_TCALL] = execTCALL
	opcodeMap[opcodes.OP_LOAD_STATIC_FUNCTION] = execLOAD_STATIC_FUNCTION
}

func execLOAD_STATIC_FUNCTION(vm runtimeapi.VM, frame runtimeapi.Frame) {
	ptr := frame.GetPtr()
	code := *frame.GetCode()

	flags := util.ReadU8LE(code, ptr)
	ptr++

	widthA := int32(flags&0x3) + 1
	widthB := int32((flags>>2)&0x3) + 1

	modNameId := util.ReadVariableLEU32(code, ptr, widthA)
	ptr += widthA
	funNameId := util.ReadVariableLEU32(code, ptr, widthB)
	ptr += widthB
	frame.SetPtr(ptr)

	modName := frame.GetConstantPool().GetConstant(int32(modNameId))
	funName := frame.GetConstantPool().GetConstant(int32(funNameId))

	if modName.GetTag() != constants.ConstantTagUTF8String {
		vm.Panic("OP_LOAD_STATIC_FUNCTION: expected module-name to be a \"ConstantUTF8String\", not \"" + reflect.TypeOf(modName).Name())
	}
	if funName.GetTag() != constants.ConstantTagUTF8String {
		vm.Panic("OP_LOAD_STATIC_FUNCTION: expected function-name to be a \"ConstantUTF8String\", not \"" + reflect.TypeOf(funName).Name())
	}

	modNameStr := modName.(*constants.ConstantUTF8String).GetValue()
	funNameStr := funName.(*constants.ConstantUTF8String).GetValue()

	callingModuleName := frame.GetModuleName()

	mod := vm.LoadModule(modNameStr)
	fn := mod.GetFunction(funNameStr)

	if callingModuleName != modNameStr && fn.GetModifier()&spec.BCFunctionModPrivate != 0 {
		vm.Panic(fmt.Sprintf("Function access violation, you cannot access the function \"%v\" from module \"%v\", it can only be accessed from its own module \"%v\"!", funNameStr, callingModuleName, modNameStr))
	}

	parentScope := vm.GetScope()
	if fn.GetModifier()&spec.BCFunctionModSubFunc != 0 {
		parentScope = frame.GetScope()
	}

	fnValue := runtime.NewFunction(parentScope, fn)
	vm.GetValueStack().Push(fnValue)
}

func execTCALL(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	fnValue := stack.Pop()

	checkType(vm, fnValue, runtimeapi.ValueTagFunction)

	fn := fnValue.(*runtime.FunctionValue)
	frame.LoadFn(fn.GetFunction())
}

func execCALL(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	fnValue := stack.Pop()

	checkType(vm, fnValue, runtimeapi.ValueTagFunction, runtimeapi.ValueTagForeignFunction, runtimeapi.ValueTagNativeFunction)

	callable := fnValue.(runtimeapi.CallableValue)
	callable.Call(vm)
}
