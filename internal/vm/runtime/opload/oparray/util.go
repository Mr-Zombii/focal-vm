package oparray

import (
	"fmt"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/public/runtimeapi"
)

func CheckArray(vm runtimeapi.VM, value rtvalue.RTValue) {
	if value.GetTag() != rtvalue.RTValueTag_ARRAY {
		vm.Panic(fmt.Sprintf("Stack value should be an array type, not type %s", value.GetType()))
	}
}

/*
[stack-in]:
├─> array A

[stack-out]:
└─> intValue
*/
func _array_instruction(vm runtimeapi.VM, action func(array *rtvalue.RTValueArray)) {
	stack := vm.GetValueStack()

	aValue := stack.Pop()
	CheckArray(vm, aValue)

	a := aValue.(*rtvalue.RTValueArray)
	action(a)
	aValue.DecRefCount()
}
