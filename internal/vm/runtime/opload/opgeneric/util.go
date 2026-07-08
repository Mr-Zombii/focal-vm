package opgeneric

import (
	"fmt"
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/vm/rtvalue"
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
	targetType bctypes.BCType, supplier func(*rtvalue.RTValuePool, T) rtvalue.RTValue,
) runtimeapi.OpcodeImpl {
	return func(vm runtimeapi.VM, f runtimeapi.Frame) {
		stack := vm.GetValueStack()
		rtpool := vm.GetRTValuePool()

		aValue := stack.Pop()

		switch aValue.GetTag() {
		case rtvalue.RTValueTag_I8:
			stack.Push(supplier(rtpool, T(aValue.(*rtvalue.RTValueI8).GetValue())))
		case rtvalue.RTValueTag_I16:
			stack.Push(supplier(rtpool, T(aValue.(*rtvalue.RTValueI16).GetValue())))
		case rtvalue.RTValueTag_I32:
			stack.Push(supplier(rtpool, T(aValue.(*rtvalue.RTValueI32).GetValue())))
		case rtvalue.RTValueTag_I64:
			stack.Push(supplier(rtpool, T(aValue.(*rtvalue.RTValueI64).GetValue())))
		case rtvalue.RTValueTag_F32:
			stack.Push(supplier(rtpool, T(aValue.(*rtvalue.RTValueF32).GetValue())))
		case rtvalue.RTValueTag_F64:
			stack.Push(supplier(rtpool, T(aValue.(*rtvalue.RTValueF64).GetValue())))
		default:
			vm.Panic(fmt.Sprintf("Unhandled conversion from type \"%s\" tag to type \"%s\"", aValue.GetType(), targetType))
		}

		aValue.DecRefCount()
	}
}
