package opint

import (
	"fmt"
	"focal-vm/internal/vm/runtime"
	"focal-vm/public/runtimeapi"
)

func CheckInt(vm runtimeapi.VM, value runtimeapi.Value) {
	if runtime.ValueIsInteger(value) {
		vm.Panic(fmt.Sprintf("Stack value should be an integer type, not type %v", value.GetTag()))
	}
}

/*
[stack-in]:
├─> intValue A
└─> intValue B

[stack-out]:
└─> intValue
*/
func _int_instruction(vm runtimeapi.VM, strict bool, action8 func(a int8, b int8), action16 func(a int16, b int16), action32 func(a int32, b int32), action64 func(a int64, b int64)) {
	stack := vm.GetValueStack()

	aValue := stack.Pop()
	bValue := stack.Pop()

	CheckInt(vm, aValue)
	CheckInt(vm, bValue)

	if aValue.GetTag() != bValue.GetTag() && strict {
		vm.Panic("OP_FADD: stack values a & b must be the same bit width, cannot mix 32 & 64")
	}

	switch aValue.GetTag() {
	case runtimeapi.ValueTagInt8:
		a := aValue.GetRawValue().(int8)
		b := bValue.GetRawValue().(int8)
		action8(a, b)
	case runtimeapi.ValueTagInt16:
		a := aValue.GetRawValue().(int16)
		b := bValue.GetRawValue().(int16)
		action16(a, b)
	case runtimeapi.ValueTagInt32:
		a := aValue.GetRawValue().(int32)
		b := bValue.GetRawValue().(int32)
		action32(a, b)
	case runtimeapi.ValueTagInt64:
		a := aValue.GetRawValue().(int64)
		b := bValue.GetRawValue().(int64)
		action64(a, b)
	}
}
