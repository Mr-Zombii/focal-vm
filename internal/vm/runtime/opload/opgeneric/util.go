package opgeneric

import (
	"fmt"
	"focal-vm/internal/vm/runtime"
	"focal-vm/public/runtimeapi"
)

/*
[stack-in]:
├─> floatValue A
└─> floatValue B

[stack-out]:
└─> floatValue
*/
func _conv_instruction[T int8 | int16 | int32 | int64 | float32 | float64](
	targetType runtimeapi.ValueTag, supplier func(T) runtimeapi.Value,
) runtimeapi.OpcodeImpl {
	return func(vm runtimeapi.VM, _ runtimeapi.Frame) {
		stack := vm.GetValueStack()

		aValue := stack.Pop()

		if runtime.ValueIsFloat(aValue) || runtime.ValueIsInteger(aValue) {
			stack.Push(supplier(T(aValue.GetRawValue())))
			return
		}
		vm.Panic(fmt.Sprintf("Unhandled conversion from type \"%v\" tag to type \"%v\"", aValue.GetTag(), targetType))
	}
}
