package opscope

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/util"
	"focal-vm/public/runtimeapi"
)

func Install_local_scope_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_DEFINE_LOCAL] = _instruction_define_local

	opcodeMap[opcodes.OP_LOCAL_EXISTS] = _instruction_local_exists
	opcodeMap[opcodes.OP_OWNS_LOCAL] = _instruction_owns_local

	opcodeMap[opcodes.OP_SET_LOCAL] = _instruction_set_local
	opcodeMap[opcodes.OP_GET_LOCAL] = _instruction_get_local
}

/*
[stack-in]:
└─> localName

[stack-out]: (empty)
*/
func _instruction_define_local(_ runtimeapi.VM, frame runtimeapi.Frame) {
	scope := frame.GetScope()
	tpool := frame.GetTypePool()
	_local_instruction(frame, func(flags uint8, localName string) {
		ptr := frame.GetPtr()
		code := *frame.GetCode()

		width := int32((flags>>2)&0x3) + 1

		typeIdx := util.ReadVariableLEI32(code, ptr, width)
		ptr += width
		frame.SetPtr(ptr)
		localType := tpool.GetType(typeIdx)

		scope.DefineLocal(localName, localType)
	})
}

/*
[stack-in]:
└─> localName

[stack-out]:
└─> booleanValue
*/
func _instruction_local_exists(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()
	scope := frame.GetScope()
	_local_instruction(frame, func(flags uint8, localName string) {
		stack.Push(rtpool.GetOrMakeRTValueBool(scope.HasLocal(localName)))
	})
}

/*
[stack-in]:
└─> localName

[stack-out]:
└─> booleanValue
*/
func _instruction_owns_local(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()
	scope := frame.GetScope()
	_local_instruction(frame, func(flags uint8, localName string) {
		stack.Push(rtpool.GetOrMakeRTValueBool(scope.OwnsLocal(localName)))
	})
}

/*
[stack-in]:
├─> localName
└─> localValue

[stack-out]: (empty)
*/
func _instruction_set_local(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	scope := frame.GetScope()
	_local_instruction(frame, func(flags uint8, localName string) {
		localValue := stack.Pop()

		err := scope.SetLocal(localName, localValue)
		if err != nil {
			vm.Panic(err.Error())
		}
	})
}

/*
[stack-in]:
└─> localName

[stack-out]:
└─> localValue
*/
func _instruction_get_local(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	scope := frame.GetScope()
	_local_instruction(frame, func(flags uint8, localName string) {
		value, err := scope.GetLocal(localName)
		if err != nil {
			vm.Panic(err.Error())
		}
		value.IncRefCount()
		stack.Push(value)
	})
}
