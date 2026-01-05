package opbool

import (
	"fmt"
	"focal-vm/internal/vm/runtime"
	"focal-vm/public/runtimeapi"
)

func CheckBool(vm runtimeapi.VM, value runtimeapi.Value) {
	if value.GetTag() == runtimeapi.ValueTagBoolean {
		vm.Panic(fmt.Sprintf("Stack value should be a boolean type, not type %v", value.GetTag()))
	}
}

func ToBoolValue(v bool) runtimeapi.Value {
	if v {
		return runtime.BOOLEAN_VALUE_TRUE
	}
	return runtime.BOOLEAN_VALUE_FALSE
}

/*
[stack-in]:
├─> intValue A
└─> intValue B

[stack-out]:
└─> intValue
*/
func _bool_instruction(vm runtimeapi.VM, action func(a bool, b bool) bool) {
	stack := vm.GetValueStack()

	aValue := stack.Pop()
	bValue := stack.Pop()

	CheckBool(vm, aValue)
	CheckBool(vm, bValue)

	a := aValue.GetRawValue().(bool)
	b := bValue.GetRawValue().(bool)

	stack.Push(ToBoolValue(action(a, b)))
}
