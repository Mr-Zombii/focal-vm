package opload

import (
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/vm/runtime"
	"focal-vm/public/runtimeapi"
)

func installArray(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_NEWARRAY] = execNEWARRAY
	opcodeMap[opcodes.OP_ALOAD] = execALOAD
	opcodeMap[opcodes.OP_ASTORE] = execASTORE
}

func execNEWARRAY(vm runtimeapi.VM, frame runtimeapi.Frame) {
	//ptr := frame.GetPtr()
	//code := *frame.GetCode()

	//flags := util.ReadU8LE(code, ptr)
	//ptr++

	//widthA := int32(flags&0x3) + 1
	//typeId := util.ReadVariableLE32(code, ptr, widthA)
	//ptr += widthA
	//frame.SetPtr(ptr)

	size := vm.GetStack().PopValue().(*runtime.Int32Value)
	backing := make([]runtimeapi.Value, size.GetValue())
	vm.GetStack().PushValue(runtime.NewArrayValue(backing))
}

func execALOAD(vm runtimeapi.VM, frame runtimeapi.Frame) {
	idx := vm.GetStack().PopValue().(*runtime.Int32Value).GetValue()
	arr := vm.GetStack().PopValue().(*runtime.ArrayValue).GetValue()

	vm.GetStack().PushValue(arr[idx])
}

func execASTORE(vm runtimeapi.VM, frame runtimeapi.Frame) {
	val := vm.GetStack().PopValue()
	idx := vm.GetStack().PopValue().(*runtime.Int32Value).GetValue()
	arr := vm.GetStack().PopValue().(*runtime.ArrayValue).GetValue()

	arr[idx] = val
}
