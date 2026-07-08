package opint

import (
	"fmt"
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/public/runtimeapi"
)

func CheckInt(vm runtimeapi.VM, value rtvalue.RTValue) {
	if _, ok := value.GetType().(*bctypes.IntegerType); !ok {
		vm.Panic(fmt.Sprintf("Stack value should be an integer type, not type %s", value.GetType()))
	}
}

/*
[stack-in]:
├─> intValue A
└─> intValue B

[stack-out]:
└─> intValue | boolValue
*/
func _int_instruction(vm runtimeapi.VM, strict bool, action8 func(a int8, b int8), action16 func(a int16, b int16), action32 func(a int32, b int32), action64 func(a int64, b int64)) {
	stack := vm.GetValueStack()

	aValue := stack.Pop()
	bValue := stack.Pop()

	CheckInt(vm, aValue)
	CheckInt(vm, bValue)

	if aValue.GetTag() != bValue.GetTag() && strict {
		vm.Panic("INT_OP: stack values a & b must be the same bit width, cannot mix 32 & 64")
	}

	switch aValue.GetTag() {
	case rtvalue.RTValueTag_I8:
		a := aValue.(*rtvalue.RTValueI8).GetValue()
		b := bValue.(*rtvalue.RTValueI8).GetValue()
		action8(a, b)
	case rtvalue.RTValueTag_I16:
		a := aValue.(*rtvalue.RTValueI16).GetValue()
		b := bValue.(*rtvalue.RTValueI16).GetValue()
		action16(a, b)
	case rtvalue.RTValueTag_I32:
		a := aValue.(*rtvalue.RTValueI32).GetValue()
		b := bValue.(*rtvalue.RTValueI32).GetValue()
		action32(a, b)
	case rtvalue.RTValueTag_I64:
		a := aValue.(*rtvalue.RTValueI64).GetValue()
		b := bValue.(*rtvalue.RTValueI64).GetValue()
		action64(a, b)
	default:
		panic("unhandled default case")
	}

	aValue.DecRefCount()
	bValue.DecRefCount()
}
