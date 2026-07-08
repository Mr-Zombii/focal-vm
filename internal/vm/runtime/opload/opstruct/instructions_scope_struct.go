package opstruct

import (
	"fmt"
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/util"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/internal/vm/runtime/opload/opint"
	"focal-vm/public/runtimeapi"
)

func Install_struct_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_STRUCT_NEW] = _instruction_struct_new
	opcodeMap[opcodes.OP_STRUCT_SET_FIELD] = _instruction_struct_set_field
	opcodeMap[opcodes.OP_STRUCT_GET_FIELD] = _instruction_struct_get_field
}

func _instruction_struct_new(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	rtpool := vm.GetRTValuePool()
	tpool := frame.GetTypePool()

	ptr := frame.GetPtr()
	code := *frame.GetCode()

	flags := util.ReadU8LE(code, ptr)
	ptr++

	width := int32(flags&0x3) + 1

	structTypeIdx := util.ReadVariableLEI32(code, ptr, width)
	structType := tpool.ExpectType(structTypeIdx, bctypes.BCTYPE_STRUCT).(*bctypes.StructType)
	ptr += width
	frame.SetPtr(ptr)

	structValue := rtpool.CreateStruct(structType)
	stack.Push(structValue)
}

/*
[stack-in]:
├─> structValue
├─> fieldValue
└─> fieldIdx

[stack-out]: (empty)
*/
func _instruction_struct_set_field(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()

	fieldIdxVal := stack.Pop()
	opint.CheckInt(vm, fieldIdxVal)

	fieldVal := stack.Pop()

	structValue := stack.Pop()
	CheckStruct(vm, structValue)

	rtStruct := structValue.(*rtvalue.RTValueStruct)
	switch fieldIdxVal.GetTag() {
	case rtvalue.RTValueTag_I8:
		rtStruct.SetField(int32(fieldIdxVal.(*rtvalue.RTValueI8).GetValue()), fieldVal)
	case rtvalue.RTValueTag_I16:
		rtStruct.SetField(int32(fieldIdxVal.(*rtvalue.RTValueI16).GetValue()), fieldVal)
	case rtvalue.RTValueTag_I32:
		rtStruct.SetField(fieldIdxVal.(*rtvalue.RTValueI32).GetValue(), fieldVal)
	default:
		vm.Panic(fmt.Sprintf("Unsupported type for indexing into structs: %s", fieldIdxVal.GetType()))
	}

	fieldIdxVal.DecRefCount()
	structValue.DecRefCount()
}

/*
[stack-in]:
├─> structValue
└─> fieldIdx

[stack-out]:
└─> fieldValue
*/
func _instruction_struct_get_field(vm runtimeapi.VM, _ runtimeapi.Frame) {
	stack := vm.GetValueStack()

	fieldIdxVal := stack.Pop()
	opint.CheckInt(vm, fieldIdxVal)

	structValue := stack.Pop()
	CheckStruct(vm, structValue)

	rtStruct := structValue.(*rtvalue.RTValueStruct)

	var fieldValue rtvalue.RTValue
	switch fieldIdxVal.GetTag() {
	case rtvalue.RTValueTag_I8:
		fieldValue = rtStruct.GetField(int32(fieldIdxVal.(*rtvalue.RTValueI8).GetValue()))
	case rtvalue.RTValueTag_I16:
		fieldValue = rtStruct.GetField(int32(fieldIdxVal.(*rtvalue.RTValueI16).GetValue()))
	case rtvalue.RTValueTag_I32:
		fieldValue = rtStruct.GetField(fieldIdxVal.(*rtvalue.RTValueI32).GetValue())
	default:
		vm.Panic(fmt.Sprintf("Unsupported type for indexing into structs: %s", fieldIdxVal.GetType()))
	}
	fieldIdxVal.DecRefCount()
	structValue.DecRefCount()

	stack.Push(fieldValue)
	fieldValue.IncRefCount()
}
