package opscope

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/util"
	"focal-vm/public/runtimeapi"
)

func Install_global_scope_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_DEFINE_GLOBAL] = _instruction_define_global

	opcodeMap[opcodes.OP_GLOBAL_EXISTS] = _instruction_global_exists
	opcodeMap[opcodes.OP_OWNS_GLOBAL] = _instruction_owns_global

	opcodeMap[opcodes.OP_SET_GLOBAL] = _instruction_set_global
	opcodeMap[opcodes.OP_GET_GLOBAL] = _instruction_get_global
}

/*
[stack-in]:
└─> globalName

[stack-out]: (empty)
*/
func _instruction_define_global(_ runtimeapi.VM, frame runtimeapi.Frame) {
	scope := frame.GetScope()
	tpool := frame.GetTypePool()
	_local_instruction(frame, func(flags uint8, globalName string) {
		ptr := frame.GetPtr()
		code := *frame.GetCode()

		width := int32((flags>>2)&0x3) + 1

		typeIdx := util.ReadVariableLEI32(code, ptr, width)
		ptr += width
		frame.SetPtr(ptr)
		localType := tpool.GetType(typeIdx)

		scope.DefineLocal(globalName, localType)
	})
}

/*
[stack-in]:
└─> globalName

[stack-out]:
└─> booleanValue
*/
func _instruction_global_exists(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()
	scope := frame.GetScope()
	_local_instruction(frame, func(flags uint8, globalName string) {
		stack.Push(rtpool.GetOrMakeRTValueBool(scope.HasLocal(globalName)))
	})
}

/*
[stack-in]:
└─> globalName

[stack-out]:
└─> booleanValue
*/
func _instruction_owns_global(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()
	scope := frame.GetScope()
	_local_instruction(frame, func(flags uint8, globalName string) {
		stack.Push(rtpool.GetOrMakeRTValueBool(scope.OwnsLocal(globalName)))
	})
}

/*
[stack-in]:
├─> globalName
└─> globalValue

[stack-out]: (empty)
*/
func _instruction_set_global(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	scope := frame.GetScope()
	_local_instruction(frame, func(flags uint8, globalName string) {
		globalValue := stack.Pop()

		err := scope.SetLocal(globalName, globalValue)
		if err != nil {
			vm.Panic(err.Error())
		}
	})
}

/*
[stack-in]:
└─> globalName

[stack-out]:
└─> globalValue
*/
func _instruction_get_global(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	scope := frame.GetScope()
	_local_instruction(frame, func(flags uint8, globalName string) {
		value, err := scope.GetLocal(globalName)
		if err != nil {
			vm.Panic(err.Error())
		}
		value.IncRefCount()
		stack.Push(value)
	})
}
