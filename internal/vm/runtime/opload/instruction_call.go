package opload

import (
	"fmt"
	"focal-vm/internal/bytecode/constants"
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/bytecode/spec"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/internal/vm/runtime"
	"focal-vm/internal/vm/runtime/opload/opcore"
	"focal-vm/public/runtimeapi"
)

func install_call_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_CALL] = execCALL
	opcodeMap[opcodes.OP_TCALL] = execTCALL
}

func execTCALL(vm runtimeapi.VM, frame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	fnValue := stack.Pop()

	opcore.CheckFunction(vm, fnValue)

	fn := fnValue.(*rtvalue.RTValueVMFunction)

	bcFn := fn.GetFunction()
	frame.LoadFn(bcFn)

	module := bcFn.GetModule()
	cpool := module.GetConstantPool()
	fnType := bcFn.GetType()

	scope := frame.GetScope()

	paramIndexes := bcFn.GetParamNameIndexes()
	paramTypeIndexes := fnType.GetParamTypeIndexes()
	for i := len(paramIndexes) - 1; i > -1; i-- {
		paramNameIndex := paramIndexes[i]
		paramTypeIndex := paramTypeIndexes[i]

		paramName := cpool.ExpectConstant(paramNameIndex, constants.ConstantTagUTF8String).(*constants.ConstantUTF8String).GetValue()
		paramType := fnType.GetTypePool().GetType(paramTypeIndex)

		stackValue := stack.Pop()
		if !stackValue.GetType().Equals(paramType) {
			vm.Panic(fmt.Sprintf("Function param type mismatch (Tail Call):\n"+
				"\tStack value \"%s\", is type \"%s\" when type \"%s\" is expected for fn param { name: \"%s\", idx: \"%d\" } \n"+
				"\n"+
				"\tCallerFnModule: \"%v\", CallerFnName: \"%v\"\n"+
				"\tFnModule: \"%s\", FnName: \"%s\", FnSignature: \"%s\"", stackValue, stackValue.GetType(), paramType, paramName, i,
				frame.GetModuleName(), frame.GetFunctionName(),
				module.GetName(), bcFn.GetName(), fnType))
			return
		}
		scope.DefineLocal(paramName, paramType)
		err := scope.SetLocal(paramName, stackValue)
		if err != nil {
			vm.Panic(err.Error())
			return
		}
	}
}

func execCALL(vm runtimeapi.VM, callerFrame runtimeapi.Frame) {
	stack := vm.GetValueStack()
	callStack := vm.GetCallStack()
	fnValue := stack.Pop()

	opcore.CheckFunction(vm, fnValue)

	fn := fnValue.(*rtvalue.RTValueVMFunction)
	bcFn := fn.GetFunction()

	parentScope := vm.GetScope()
	if bcFn.GetModifier()&spec.BCFunctionModSubFunc != 0 {
		parentScope = callerFrame.GetScope()
	}

	module := fn.GetFunction().GetModule()
	frame := runtime.NewFrame(callerFrame, parentScope, module, bcFn)
	scope := frame.GetScope()

	cpool := module.GetConstantPool()
	fnType := bcFn.GetType()

	paramIndexes := bcFn.GetParamNameIndexes()
	paramTypeIndexes := fnType.GetParamTypeIndexes()
	for i := len(paramIndexes) - 1; i > -1; i-- {
		paramNameIndex := paramIndexes[i]
		paramTypeIndex := paramTypeIndexes[i]

		paramName := cpool.ExpectConstant(paramNameIndex, constants.ConstantTagUTF8String).(*constants.ConstantUTF8String).GetValue()
		paramType := fnType.GetTypePool().GetType(paramTypeIndex)

		stackValue := stack.Pop()
		if !stackValue.GetType().Equals(paramType) {
			vm.Panic(fmt.Sprintf("Function param type mismatch:\n"+
				"\tStack value \"%s\", is type \"%s\" when type \"%s\" is expected for fn param { name: \"%s\", idx: \"%d\" } \n"+
				"\n"+
				"\tCallerFnModule: \"%s\", CallerFnName: \"%s\"\n"+
				"\tFnModule: \"%s\", FnName: \"%s\", FnSignature: \"%s\"", stackValue, stackValue.GetType(), paramType, paramName, i,
				callerFrame.GetModuleName(), callerFrame.GetFunctionName(),
				module.GetName(), bcFn.GetName(), fnType))
			return
		}
		scope.DefineLocal(paramName, paramType)
		err := scope.SetLocal(paramName, stackValue)
		if err != nil {
			vm.Panic(err.Error())
			return
		}
	}

	callStack.Push(frame)
}
