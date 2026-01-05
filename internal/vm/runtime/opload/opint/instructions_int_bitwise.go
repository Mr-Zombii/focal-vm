package opint

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/vm/runtime"
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

func _instruction_rsh(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_int_instruction(vm, false, func(a, b int8) {
		stack.Push(runtime.NewInt8Value(a >> b))
	}, func(a, b int16) {
		stack.Push(runtime.NewInt16Value(a >> b))
	}, func(a, b int32) {
		stack.Push(runtime.NewInt32Value(a >> b))
	}, func(a, b int64) {
		stack.Push(runtime.NewInt64Value(a >> b))
	})
}

func _instruction_lsh(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_int_instruction(vm, false, func(a, b int8) {
		stack.Push(runtime.NewInt8Value(a << b))
	}, func(a, b int16) {
		stack.Push(runtime.NewInt16Value(a << b))
	}, func(a, b int32) {
		stack.Push(runtime.NewInt32Value(a << b))
	}, func(a, b int64) {
		stack.Push(runtime.NewInt64Value(a << b))
	})
}

func _instruction_rrt(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_int_instruction(vm, false, func(a, b int8) {
		ua := bits.Reverse8(uint8(a))
		o := bits.Reverse8(bits.RotateLeft8(ua, int(b)))
		stack.Push(runtime.NewInt8Value(int8(o)))
	}, func(a, b int16) {
		ua := bits.Reverse16(uint16(a))
		o := bits.Reverse16(bits.RotateLeft16(ua, int(b)))
		stack.Push(runtime.NewInt16Value(int16(o)))
	}, func(a, b int32) {
		ua := bits.Reverse32(uint32(a))
		o := bits.Reverse32(bits.RotateLeft32(ua, int(b)))
		stack.Push(runtime.NewInt32Value(int32(o)))
	}, func(a, b int64) {
		ua := bits.Reverse64(uint64(a))
		o := bits.Reverse64(bits.RotateLeft64(ua, int(b)))
		stack.Push(runtime.NewInt64Value(int64(o)))
	})
}

func _instruction_lrt(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_int_instruction(vm, false, func(a, b int8) {
		stack.Push(runtime.NewInt8Value(int8(bits.RotateLeft8(uint8(a), int(b)))))
	}, func(a, b int16) {
		stack.Push(runtime.NewInt16Value(int16(bits.RotateLeft16(uint16(a), int(b)))))
	}, func(a, b int32) {
		stack.Push(runtime.NewInt32Value(int32(bits.RotateLeft32(uint32(a), int(b)))))
	}, func(a, b int64) {
		stack.Push(runtime.NewInt64Value(int64(bits.RotateLeft64(uint64(a), int(b)))))
	})
}

func _instruction_bnot(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	aValue := stack.Pop()
	CheckInt(vm, aValue)

	switch aValue.GetTag() {
	case runtimeapi.ValueTagInt8:
		a := aValue.GetRawValue().(int8)
		stack.Push(runtime.NewInt8Value(^a))
	case runtimeapi.ValueTagInt16:
		a := aValue.GetRawValue().(int16)
		stack.Push(runtime.NewInt16Value(^a))
	case runtimeapi.ValueTagInt32:
		a := aValue.GetRawValue().(int32)
		stack.Push(runtime.NewInt32Value(^a))
	case runtimeapi.ValueTagInt64:
		a := aValue.GetRawValue().(int64)
		stack.Push(runtime.NewInt64Value(^a))
	}
}

func _instruction_bor(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_int_instruction(vm, false, func(a, b int8) {
		stack.Push(runtime.NewInt8Value(a | b))
	}, func(a, b int16) {
		stack.Push(runtime.NewInt16Value(a | b))
	}, func(a, b int32) {
		stack.Push(runtime.NewInt32Value(a | b))
	}, func(a, b int64) {
		stack.Push(runtime.NewInt64Value(a | b))
	})
}

func _instruction_band(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_int_instruction(vm, false, func(a, b int8) {
		stack.Push(runtime.NewInt8Value(a & b))
	}, func(a, b int16) {
		stack.Push(runtime.NewInt16Value(a & b))
	}, func(a, b int32) {
		stack.Push(runtime.NewInt32Value(a & b))
	}, func(a, b int64) {
		stack.Push(runtime.NewInt64Value(a & b))
	})
}

func _instruction_bxor(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_int_instruction(vm, false, func(a, b int8) {
		stack.Push(runtime.NewInt8Value(a ^ b))
	}, func(a, b int16) {
		stack.Push(runtime.NewInt16Value(a ^ b))
	}, func(a, b int32) {
		stack.Push(runtime.NewInt32Value(a ^ b))
	}, func(a, b int64) {
		stack.Push(runtime.NewInt64Value(a ^ b))
	})
}
