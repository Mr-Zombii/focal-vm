package opload

import (
	"focal-lang/internal/bytecode/opcodes"
	"focal-lang/internal/vm/api"
	"focal-lang/internal/vm/runtime"
)

func installArray(opcodeMap []api.OpcodeImpl) {
	opcodeMap[opcodes.OP_NEWARRAY] = execNEWARRAY
	opcodeMap[opcodes.OP_ALOAD] = execALOAD
	opcodeMap[opcodes.OP_ASTORE] = execASTORE
}

func execNEWARRAY(vm api.VM, frame api.Frame) {
	//ptr := frame.GetPtr()
	//code := *frame.GetCode()

	//flags := util.ReadU8LE(code, ptr)
	//ptr++

	//widthA := int32(flags&0x3) + 1
	//typeId := util.ReadVariableLE32(code, ptr, widthA)
	//ptr += widthA
	//frame.SetPtr(ptr)

	size := vm.GetStack().PopValue().(*runtime.Int32Value)
	backing := make([]api.Value, size.GetValue())
	vm.GetStack().PushValue(runtime.NewArrayValue(backing))
}

func execALOAD(vm api.VM, frame api.Frame) {
	idx := vm.GetStack().PopValue().(*runtime.Int32Value).GetValue()
	arr := vm.GetStack().PopValue().(*runtime.ArrayValue).GetValue()

	vm.GetStack().PushValue(arr[idx])
}

func execASTORE(vm api.VM, frame api.Frame) {
	val := vm.GetStack().PopValue()
	idx := vm.GetStack().PopValue().(*runtime.Int32Value).GetValue()
	arr := vm.GetStack().PopValue().(*runtime.ArrayValue).GetValue()

	arr[idx] = val
}
