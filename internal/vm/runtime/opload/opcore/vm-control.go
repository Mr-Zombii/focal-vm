package opcore

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/util"
	"focal-vm/internal/vm/runtime/opload/opbool"
	"focal-vm/public/runtimeapi"
)

func Install_control_flow(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_HALT] = _instruction_halt
	opcodeMap[opcodes.OP_JUMP] = _instruction_jump
	opcodeMap[opcodes.OP_BRANCH] = _instruction_branch
	opcodeMap[opcodes.OP_NOP] = func(runtimeapi.VM, runtimeapi.Frame) {}
}

func _instruction_halt(vm runtimeapi.VM, _ runtimeapi.Frame) {
	vm.Halt(0)
}

func _instruction_jump(_ runtimeapi.VM, frame runtimeapi.Frame) {
	ptr := frame.GetPtr()
	code := *frame.GetCode()

	flags := util.ReadU8LE(code, ptr)
	ptr++

	widthA := int32(flags&0x3) + 1

	relativeJumpAddress := util.ReadVariableLEI32(code, ptr, widthA)
	ptr += widthA
	ptr += relativeJumpAddress
	frame.SetPtr(ptr)
}

func _instruction_branch(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()

	conditionValue := stack.Pop()
	opbool.CheckBool(vm, conditionValue)
	condition := conditionValue.GetRawValue().(bool)

	if condition {
		_instruction_jump(vm, frame)
	}
}
