package opint

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/public/runtimeapi"
	"math/bits"
)

func Install_int_bitwise_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_RSH] = _instruction_rsh
	opcodeMap[opcodes.OP_LSH] = _instruction_lsh
	opcodeMap[opcodes.OP_RRT] = _instruction_rrt
	opcodeMap[opcodes.OP_LRT] = _instruction_lrt
	opcodeMap[opcodes.OP_BNOT] = _instruction_bnot
	opcodeMap[opcodes.OP_BOR] = _instruction_bor
	opcodeMap[opcodes.OP_BAND] = _instruction_band
	opcodeMap[opcodes.OP_BXOR] = _instruction_bxor
}

func _instruction_rsh(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_int_instruction(vm, false, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueI8(a >> b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueI16(a >> b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueI32(a >> b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueI64(a >> b))
	})
}

func _instruction_lsh(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_int_instruction(vm, false, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueI8(a << b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueI16(a << b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueI32(a << b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueI64(a << b))
	})
}

func _instruction_rrt(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_int_instruction(vm, false, func(a, b int8) {
		ua := bits.Reverse8(uint8(a))
		o := bits.Reverse8(bits.RotateLeft8(ua, int(b)))
		stack.Push(rtpool.GetOrMakeRTValueI8(int8(o)))
	}, func(a, b int16) {
		ua := bits.Reverse16(uint16(a))
		o := bits.Reverse16(bits.RotateLeft16(ua, int(b)))
		stack.Push(rtpool.GetOrMakeRTValueI16(int16(o)))
	}, func(a, b int32) {
		ua := bits.Reverse32(uint32(a))
		o := bits.Reverse32(bits.RotateLeft32(ua, int(b)))
		stack.Push(rtpool.GetOrMakeRTValueI32(int32(o)))
	}, func(a, b int64) {
		ua := bits.Reverse64(uint64(a))
		o := bits.Reverse64(bits.RotateLeft64(ua, int(b)))
		stack.Push(rtpool.GetOrMakeRTValueI64(int64(o)))
	})
}

func _instruction_lrt(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_int_instruction(vm, false, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueI8(int8(bits.RotateLeft8(uint8(a), int(b)))))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueI16(int16(bits.RotateLeft16(uint16(a), int(b)))))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueI32(int32(bits.RotateLeft32(uint32(a), int(b)))))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueI64(int64(bits.RotateLeft64(uint64(a), int(b)))))
	})
}

func _instruction_bnot(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	aValue := stack.Pop()
	CheckInt(vm, aValue)

	switch aValue.GetTag() {
	case rtvalue.RTValueTag_I8:
		a := aValue.(*rtvalue.RTValueI8).GetValue()
		stack.Push(rtpool.GetOrMakeRTValueI8(^a))
	case rtvalue.RTValueTag_I16:
		a := aValue.(*rtvalue.RTValueI16).GetValue()
		stack.Push(rtpool.GetOrMakeRTValueI16(^a))
	case rtvalue.RTValueTag_I32:
		a := aValue.(*rtvalue.RTValueI32).GetValue()
		stack.Push(rtpool.GetOrMakeRTValueI32(^a))
	case rtvalue.RTValueTag_I64:
		a := aValue.(*rtvalue.RTValueI64).GetValue()
		stack.Push(rtpool.GetOrMakeRTValueI64(^a))
	default:
		panic("unhandled default case")
	}

	aValue.DecRefCount()
}

func _instruction_bor(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_int_instruction(vm, false, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueI8(a | b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueI16(a | b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueI32(a | b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueI64(a | b))
	})
}

func _instruction_band(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_int_instruction(vm, false, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueI8(a & b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueI16(a & b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueI32(a & b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueI64(a & b))
	})
}

func _instruction_bxor(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()

	_int_instruction(vm, false, func(a, b int8) {
		stack.Push(rtpool.GetOrMakeRTValueI8(a ^ b))
	}, func(a, b int16) {
		stack.Push(rtpool.GetOrMakeRTValueI16(a ^ b))
	}, func(a, b int32) {
		stack.Push(rtpool.GetOrMakeRTValueI32(a ^ b))
	}, func(a, b int64) {
		stack.Push(rtpool.GetOrMakeRTValueI64(a ^ b))
	})
}
