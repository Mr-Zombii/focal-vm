package oparray

import (
	"fmt"
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/util"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/internal/vm/runtime/opload/opint"
	"focal-vm/public/runtimeapi"
)

func Install_array_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_ARRAY_NEW] = _instruction_array_new
	opcodeMap[opcodes.OP_ARRAY_STORE] = _instruction_array_store
	opcodeMap[opcodes.OP_ARRAY_LOAD] = _instruction_array_load
}

/*
[stack-in]:
└─> arraySize

[stack-out]:
└─> arrayValue
*/
func _instruction_array_new(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()
	tpool := frame.GetTypePool()

	lengthValue := stack.Pop()
	opint.CheckInt(vm, lengthValue)

	var length int32
	switch lengthValue.GetTag() {
	case rtvalue.RTValueTag_I8:
		length = int32(lengthValue.(*rtvalue.RTValueI8).GetValue())
	case rtvalue.RTValueTag_I16:
		length = int32(lengthValue.(*rtvalue.RTValueI16).GetValue())
	case rtvalue.RTValueTag_I32:
		length = lengthValue.(*rtvalue.RTValueI32).GetValue()
	default:
		vm.Panic(fmt.Sprintf("Cannot use type \"%s\" for new array length!", lengthValue.GetType()))
	}
	lengthValue.DecRefCount()

	ptr := frame.GetPtr()
	code := *frame.GetCode()

	flags := util.ReadU8LE(code, ptr)
	ptr++

	width := int32(flags&0x3) + 1

	elemTypeIdx := util.ReadVariableLEI32(code, ptr, width)
	elemType := tpool.GetType(elemTypeIdx)
	ptr += width
	frame.SetPtr(ptr)

	stack.Push(rtpool.CreateArray(elemType, length))
}

/*
[stack-in]:
├─> array
└─> index

[stack-out]:
└─> arrayElement
*/
func _instruction_array_load(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_array_instruction(vm, func(array *rtvalue.RTValueArray) {
		indexValue := stack.Pop()
		opint.CheckInt(vm, indexValue)

		var index int32
		switch indexValue.GetTag() {
		case rtvalue.RTValueTag_I8:
			index = int32(indexValue.(*rtvalue.RTValueI8).GetValue())
		case rtvalue.RTValueTag_I16:
			index = int32(indexValue.(*rtvalue.RTValueI16).GetValue())
		case rtvalue.RTValueTag_I32:
			index = indexValue.(*rtvalue.RTValueI32).GetValue()
		default:
			vm.Panic(fmt.Sprintf("Cannot use type \"%s\" for array index!", indexValue.GetType()))
		}
		indexValue.DecRefCount()

		if int32(len(array.GetValue())-1) < index || index < 0 {
			vm.Panic(fmt.Sprintf("Index %d out of bounds for array size %d", index, array.GetLength()))
		}

		stack.Push(array.GetValue()[index])
	})
}

/*
[stack-in]:
├─> array
├─> index
└─> value

[stack-out]:
*/
func _instruction_array_store(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_array_instruction(vm, func(array *rtvalue.RTValueArray) {
		indexValue := stack.Pop()
		opint.CheckInt(vm, indexValue)

		elementValue := stack.Pop()
		if !elementValue.GetType().Equals(array.GetType()) {
			vm.Panic(fmt.Sprintf("Cannot store value of type \"%s\" for in array of type \"%s\"", elementValue.GetType(), array.GetType()))
			return
		}

		var index int32
		switch indexValue.GetTag() {
		case rtvalue.RTValueTag_I8:
			index = int32(indexValue.(*rtvalue.RTValueI8).GetValue())
		case rtvalue.RTValueTag_I16:
			index = int32(indexValue.(*rtvalue.RTValueI16).GetValue())
		case rtvalue.RTValueTag_I32:
			index = indexValue.(*rtvalue.RTValueI32).GetValue()
		default:
			vm.Panic(fmt.Sprintf("Cannot use type \"%s\" for new array index!", indexValue.GetType()))
		}
		indexValue.DecRefCount()

		if int32(len(array.GetValue())-1) < index || index < 0 {
			vm.Panic(fmt.Sprintf("Index %d out of bounds for array size %d", index, array.GetLength()))
		}

		array.GetValue()[index] = elementValue
	})
}
