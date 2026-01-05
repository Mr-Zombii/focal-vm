package opload

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/util"
	"focal-vm/internal/vm/runtime"
	"focal-vm/internal/vm/runtime/opload/opbool"
	"focal-vm/public/runtimeapi"
)

func install_scope_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_LOAD_CONST] = execLOAD_CONST

	opcodeMap[opcodes.OP_DEFINE_GLOBAL] = execDEFINE_GLOBAL
	opcodeMap[opcodes.OP_DEFINE_LOCAL] = execDEFINE_LOCAL

	opcodeMap[opcodes.OP_LOCAL_EXISTS] = execLOCAL_EXISTS
	opcodeMap[opcodes.OP_OWNS_LOCAL] = execOWNS_LOCAL

	opcodeMap[opcodes.OP_SET_LOCAL] = execSET_LOCAL
	opcodeMap[opcodes.OP_GET_LOCAL] = execGET_LOCAL
}

/*
[stack-in]: (empty)

[stack-out]:
└─> constantValue
*/
func execLOAD_CONST(vm runtimeapi.VM, frame runtimeapi.Frame) {
	ptr := frame.GetPtr()
	code := *frame.GetCode()
	stack := vm.GetValueStack()

	flags := util.ReadU8LE(code, ptr)
	ptr++

	width := int32(flags&0x3) + 1

	constantIdx := util.ReadVariableLEU32(code, ptr, width)
	constant := frame.GetConstantPool().GetConstant(int32(constantIdx))
	ptr += width

	stack.Push(runtime.ConstantToValue(constant))
	frame.SetPtr(ptr)
}

/*
[stack-in]:
└─> globalName

[stack-out]: (empty)
*/
func execDEFINE_GLOBAL(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	nameValue := stack.Pop()
	checkType(vm, nameValue, runtimeapi.ValueTagUTF8String)

	name := nameValue.GetRawValue().(string)
	vm.GetScope().DefineLocal(name)
}

/*
[stack-in]:
└─> localName

[stack-out]: (empty)
*/
func execDEFINE_LOCAL(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	nameValue := stack.Pop()
	checkType(vm, nameValue, runtimeapi.ValueTagUTF8String)

	name := nameValue.GetRawValue().(string)
	frame.GetScope().DefineLocal(name)
}

/*
[stack-in]:
└─> localName

[stack-out]:
└─> booleanValue
*/
func execLOCAL_EXISTS(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	nameValue := stack.Pop()
	checkType(vm, nameValue, runtimeapi.ValueTagUTF8String)

	name := nameValue.GetRawValue().(string)

	stack.Push(opbool.ToBoolValue(frame.GetScope().HasLocal(name)))
}

/*
[stack-in]:
└─> localName

[stack-out]:
└─> booleanValue
*/
func execOWNS_LOCAL(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	nameValue := stack.Pop()
	checkType(vm, nameValue, runtimeapi.ValueTagUTF8String)

	name := nameValue.GetRawValue().(string)

	stack.Push(opbool.ToBoolValue(frame.GetScope().OwnsLocal(name)))
}

/*
[stack-in]:
├─> localName
└─> localValue

[stack-out]: (empty)
*/
func execSET_LOCAL(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	nameValue := stack.Pop()
	localValue := stack.Pop()
	checkType(vm, nameValue, runtimeapi.ValueTagUTF8String)

	name := nameValue.GetRawValue().(string)

	err := frame.GetScope().SetLocal(name, localValue)
	if err != nil {
		vm.Panic(err.Error())
	}
}

/*
[stack-in]:
└─> localName

[stack-out]:
└─> localValue
*/
func execGET_LOCAL(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	nameValue := stack.Pop()
	checkType(vm, nameValue, runtimeapi.ValueTagUTF8String)

	name := nameValue.GetRawValue().(string)

	value, err := frame.GetScope().GetLocal(name)
	if err != nil {
		vm.Panic(err.Error())
	}
	stack.Push(value)
}
