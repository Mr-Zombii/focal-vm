package oparray

import (
	"fmt"
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/vm/runtime"
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

	sizeValue := stack.Pop()
	opint.CheckInt(vm, sizeValue)

	size := runtime.GetValueAsInt(sizeValue)

	backingArray := make([]runtimeapi.Value, size)
	stack.Push(runtime.NewArrayValue(backingArray))
}

/*
[stack-in]:
├─> array
└─> index

[stack-out]:
└─> arrayElement
*/
func _instruction_array_load(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_array_instruction(vm, func(array *runtime.ArrayValue) {
		indexValue := stack.Pop()
		opint.CheckInt(vm, indexValue)

		index := runtime.GetValueAsInt(indexValue)

		if len(array.GetValue())-1 < index || index < 0 {
			vm.Panic(fmt.Sprintf("Index %v out of bounds for array size %v", index, array.GetLength()))
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
func _instruction_array_store(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	_array_instruction(vm, func(array *runtime.ArrayValue) {
		indexValue := stack.Pop()
		elementValue := stack.Pop()
		opint.CheckInt(vm, indexValue)

		index := runtime.GetValueAsInt(indexValue)

		if len(array.GetValue())-1 < index || index < 0 {
			vm.Panic(fmt.Sprintf("Index %v out of bounds for array size %v", index, array.GetLength()))
		}

		array.GetValue()[index] = elementValue
	})
}
