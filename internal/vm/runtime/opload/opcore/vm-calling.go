package opcore

import (
	"fmt"
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/bytecode/constants"
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/bytecode/spec"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/internal/vm/runtime"
	"focal-vm/internal/vm/runtime/ffi"
	"focal-vm/public/runtimeapi"
)

func Install_call_instructions(opcodeMap []runtimeapi.OpcodeImpl) {
	opcodeMap[opcodes.OP_CALL] = _call_instruction(false)
	opcodeMap[opcodes.OP_TCALL] = _call_instruction(true)
}

func _call_instruction(tail bool) runtimeapi.OpcodeImpl {
	return func(vm runtimeapi.VM, callerFrame runtimeapi.Frame) {
		stack := vm.GetValueStack()
		fnValue := stack.Pop()

		CheckFunction(vm, fnValue)

		if v, ok := fnValue.(*rtvalue.RTValueGOFunction); ok {
			if tail {
				vm.Panic("Cannot do a tail call on a native function!")
				return
			}
			ffi.CallBuiltinFunction(vm, callerFrame.GetTypePool(), v.GetFunction())
			v.DecRefCount()
			return
		}

		var targetFrame runtimeapi.Frame

		byteFunction := fnValue.(*rtvalue.RTValueVMFunction).GetFunction()
		byteModule := byteFunction.GetModule()

		cpool := byteModule.GetConstantPool()
		tpool := byteModule.GetTypePool()

		fnType := tpool.ExpectType(byteFunction.GetTypeIdx(), bctypes.BCTYPE_FUNCTION).(*bctypes.FunctionType)

		if tail {
			targetFrame = callerFrame
			targetFrame.LoadFn(byteFunction)
		} else {
			var targetParentScope runtimeapi.Scope

			if byteFunction.GetModifier()&spec.BCFunctionModSubFunc != 0 {
				targetParentScope = callerFrame.GetScope()
			} else {
				targetParentScope = vm.GetScope()
			}

			targetFrame = runtime.NewFrame(callerFrame, targetParentScope, byteModule, byteFunction)
		}

		scope := targetFrame.GetScope()

		paramIndexes := byteFunction.GetParamNameIndexes()
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
					byteModule.GetName(), byteFunction.GetName(), fnType))
				return
			}

			scope.DefineLocal(paramName, paramType)
			err := scope.SetLocal(paramName, stackValue)
			if err != nil {
				vm.Panic(err.Error())
				return
			}
		}

		if !tail {
			vm.GetCallStack().Push(targetFrame)
		}

		fnValue.DecRefCount()
	}
}
