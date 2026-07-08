package opbool

import (
	"fmt"
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/public/runtimeapi"
)

func CheckBool(vm runtimeapi.VM, value rtvalue.RTValue) {
	if value.GetType().GetTag() != bctypes.BCTYPE_BOOLEAN {
		vm.Panic(fmt.Sprintf("Stack value should be a boolean type, not type %s", value.GetType()))
	}
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
	rtpool := vm.GetRTValuePool()

	aValue := stack.Pop()
	bValue := stack.Pop()

	CheckBool(vm, aValue)
	CheckBool(vm, bValue)

	a := aValue.(*rtvalue.RTValueBool).GetValue()
	b := bValue.(*rtvalue.RTValueBool).GetValue()

	stack.Push(rtpool.GetOrMakeRTValueBool(action(a, b)))
	aValue.DecRefCount()
	bValue.DecRefCount()
}
