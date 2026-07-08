package opmath

import (
	"fmt"
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/public/runtimeapi"
)

type Number = float32

func CheckNumber(vm runtimeapi.VM, value rtvalue.RTValue) {
	_, ok := value.GetType().(*bctypes.IntegerType)
	_, ok2 := value.GetType().(*bctypes.FloatType)
	if !(ok || ok2) {
		vm.Panic(fmt.Sprintf("Stack value should be an integer or float type, not type %s", value.GetType()))
	}
}

func _number_instruction(
	vm runtimeapi.VM, strict bool,
	action8 func(a int8, b int8),
	action16 func(a int16, b int16),
	action32 func(a int32, b int32),
	action64 func(a int64, b int64),
	actionF32 func(a float32, b float32),
	actionF64 func(a float64, b float64),
) {
	stack := vm.GetValueStack()

	aValue := stack.Pop()
	bValue := stack.Pop()

	CheckNumber(vm, aValue)
	CheckNumber(vm, bValue)

	if aValue.GetTag() != bValue.GetTag() && strict {
		vm.Panic("INT_OP: stack values a & b must be the same bit width, cannot mix 2 different bit widths nor mix float & int")
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
	case rtvalue.RTValueTag_F32:
		a := aValue.(*rtvalue.RTValueF32).GetValue()
		b := bValue.(*rtvalue.RTValueF32).GetValue()
		actionF32(a, b)
	case rtvalue.RTValueTag_F64:
		a := aValue.(*rtvalue.RTValueF64).GetValue()
		b := bValue.(*rtvalue.RTValueF64).GetValue()
		actionF64(a, b)
	default:
		panic("unhandled default case")
	}

	aValue.DecRefCount()
	bValue.DecRefCount()
}
