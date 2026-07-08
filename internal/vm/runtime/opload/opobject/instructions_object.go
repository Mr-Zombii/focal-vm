package opobject

import (
	"focal-vm/public/runtimeapi"
)

func Install_object_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	//opcodeMap[opcodes.OP_OBJECT_NEW] = _instruction_new_object
	//opcodeMap[opcodes.OP_OBJECT_SET_FIELD] = _instruction_object_set_field
	//opcodeMap[opcodes.OP_OBJECT_GET_FIELD] = _instruction_object_get_field
	//opcodeMap[opcodes.OP_OBJECT_DEFINE_FIELD] = _instruction_object_define_field
	//opcodeMap[opcodes.OP_OBJECT_HAS_FIELD] = _instruction_object_has_field
}

//
///*
//[stack-in]: (empty)
//
//[stack-out]:
//└─> object
//*/
//func _instruction_new_object(vm runtimeapi.VM, _ runtimeapi.Frame) {
//	vm.GetValueStack().Push(runtime.NewScopeValue(vm.GetScope().NewChildScope()))
//}
//
///*
//[stack-in]:
//├─> object
//├─> fieldName
//└─> fieldValue
//
//[stack-out]: (empty)
//*/
//func _instruction_object_set_field(vm runtimeapi.VM, _ runtimeapi.Frame) {
//	stack := vm.GetValueStack()
//
//	_object_instruction(vm, func(object *runtime.ScopeValue, fieldName string) {
//		fieldValue := stack.Pop()
//
//		err := object.SetField(fieldName, fieldValue)
//		if err != nil {
//			vm.Panic(err.Error())
//		}
//	})
//}
//
///*
//[stack-in]:
//├─> object
//└─> fieldName
//
//[stack-out]:
//└─> fieldValue
//*/
//func _instruction_object_get_field(vm runtimeapi.VM, _ runtimeapi.Frame) {
//	stack := vm.GetValueStack()
//
//	_object_instruction(vm, func(object *runtime.ScopeValue, fieldName string) {
//		fieldValue, err := object.GetField(fieldName)
//		if err != nil {
//			vm.Panic(err.Error())
//		}
//
//		stack.Push(fieldValue)
//	})
//}
//
///*
//[stack-in]:
//├─> object
//└─> fieldName
//
//[stack-out]: (empty)
//*/
//func _instruction_object_define_field(vm runtimeapi.VM, _ runtimeapi.Frame) {
//	_object_instruction(vm, func(object *runtime.ScopeValue, fieldName string) {
//		object.DefineField(fieldName)
//	})
//}
//
///*
//[stack-in]:
//├─> object
//└─> fieldName
//
//[stack-out]:
//└─> booleanValue
//*/
//func _instruction_object_has_field(vm runtimeapi.VM, _ runtimeapi.Frame) {
//	stack := vm.GetValueStack()
//	rtpool := vm.GetRTValuePool()
//
//	_object_instruction(vm, func(object *runtime.ScopeValue, fieldName string) {
//		stack.Push(rtpool.GetOrMakeRTValueBool(object.HasField(fieldName)))
//	})
//}
