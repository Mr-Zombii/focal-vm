package oparray

import (
	"fmt"
	"focal-vm/internal/vm/runtime"
	"focal-vm/public/runtimeapi"
)

func CheckArray(vm runtimeapi.VM, value runtimeapi.Value) {
	if value.GetTag() == runtimeapi.ValueTagArray {
		vm.Panic(fmt.Sprintf("Stack value should be an array type, not type %v", value.GetTag()))
	}
}

/*
[stack-in]:
├─> array A

[stack-out]:
└─> intValue
*/
func _array_instruction(vm runtimeapi.VM, action func(array *runtime.ArrayValue)) {
	stack := vm.GetValueStack()

	aValue := stack.Pop()
	CheckArray(vm, aValue)

	a := aValue.GetRawValue().(*runtime.ArrayValue)
	action(a)
}
