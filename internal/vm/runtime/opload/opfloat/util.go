package opfloat

import (
	"fmt"
	"focal-vm/internal/vm/runtime"
	"focal-vm/public/runtimeapi"
)

func CheckFloat(vm runtimeapi.VM, value runtimeapi.Value) {
	if runtime.ValueIsFloat(value) {
		vm.Panic(fmt.Sprintf("Stack value should be a float type, not type %v", value.GetTag()))
	}
}

/*
[stack-in]:
├─> floatValue A
└─> floatValue B

[stack-out]:
└─> floatValue
*/
func _float_instruction(vm runtimeapi.VM, action32 func(a float32, b float32), action64 func(a float64, b float64)) {
	stack := vm.GetValueStack()

	aValue := stack.Pop()
	bValue := stack.Pop()

	CheckFloat(vm, aValue)
	CheckFloat(vm, bValue)

	is32Bit := aValue.GetTag() == runtimeapi.ValueTagFloat32

	if aValue.GetTag() != bValue.GetTag() {
		vm.Panic("OP_FADD: stack values a & b must be the same bit width, cannot mix 32 & 64")
	}

	if is32Bit {
		a := aValue.GetRawValue().(float32)
		b := bValue.GetRawValue().(float32)

		action32(a, b)
		return
	}

	a := aValue.GetRawValue().(float64)
	b := bValue.GetRawValue().(float64)

	action64(a, b)
}
