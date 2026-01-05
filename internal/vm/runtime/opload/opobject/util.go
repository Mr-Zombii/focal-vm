package opobject

import (
	"fmt"
	"focal-vm/internal/vm/runtime"
	"focal-vm/public/runtimeapi"
)

func CheckObject(vm runtimeapi.VM, value runtimeapi.Value) {
	if value.GetTag() == runtimeapi.ValueTagScope {
		vm.Panic(fmt.Sprintf("Stack value should be an object type, not type %v", value.GetTag()))
	}
}

func CheckUtf(vm runtimeapi.VM, value runtimeapi.Value) {
	if value.GetTag() == runtimeapi.ValueTagScope {
		vm.Panic(fmt.Sprintf("Stack value should be a string type, not type %v", value.GetTag()))
	}
}

/*
[stack-in]:
├─> object A
└─> fieldName B

[stack-out]:
└─> intValue
*/
func _object_instruction(vm runtimeapi.VM, action func(object *runtime.ScopeValue, fieldName string)) {
	stack := vm.GetValueStack()

	aValue := stack.Pop()
	bValue := stack.Pop()

	CheckObject(vm, aValue)
	CheckUtf(vm, bValue)

	a := aValue.GetRawValue().(*runtime.ScopeValue)
	b := bValue.GetRawValue().(string)

	action(a, b)
}
