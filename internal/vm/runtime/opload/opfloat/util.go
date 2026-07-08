package opfloat

import (
	"fmt"
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/public/runtimeapi"
)

func CheckFloat(vm runtimeapi.VM, value rtvalue.RTValue) {
	if _, ok := value.GetType().(*bctypes.FloatType); !ok {
		vm.Panic(fmt.Sprintf("Stack value should be a float type, not type %s", value.GetType()))
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

	is32Bit := aValue.GetType().GetTag() == bctypes.BCTYPE_F32

	if aValue.GetTag() != bValue.GetTag() {
		vm.Panic("OP_FADD: stack values a & b must be the same bit width, cannot mix 32 & 64")
	}

	if is32Bit {
		a := aValue.(*rtvalue.RTValueF32).GetValue()
		b := bValue.(*rtvalue.RTValueF32).GetValue()

		action32(a, b)
		aValue.DecRefCount()
		bValue.DecRefCount()
		return
	}

	a := aValue.(*rtvalue.RTValueF64).GetValue()
	b := bValue.(*rtvalue.RTValueF64).GetValue()

	action64(a, b)
	aValue.DecRefCount()
	bValue.DecRefCount()
}
