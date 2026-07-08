package opscope

import (
	"focal-vm/internal/bytecode/constants"
	"focal-vm/internal/util"
	"focal-vm/public/runtimeapi"
)

func _local_instruction(frame runtimeapi.Frame, action func(flags uint8, localName string)) {
	cpool := frame.GetConstantPool()

	ptr := frame.GetPtr()
	code := *frame.GetCode()

	flags := util.ReadU8LE(code, ptr)
	ptr++

	width := int32(flags&0x3) + 1

	constantIdx := util.ReadVariableLEI32(code, ptr, width)
	ptr += width
	frame.SetPtr(ptr)

	localName := cpool.ExpectConstant(constantIdx, constants.ConstantTagUTF8String).(*constants.ConstantUTF8String).GetValue()

	action(flags, localName)
}
