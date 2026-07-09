package ffi

import (
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/internal/vm/runtime"
	"focal-vm/public/runtimeapi"
	"reflect"
	goruntime "runtime"
)

func NewBuiltinFunction(vm runtimeapi.VM, tpool *bctypes.TypePool, fn interface{}) *rtvalue.RTValueGOFunction {
	t, err := ReflectionTypeToRuntimeType(tpool, reflect.TypeOf(fn))
	if err != nil {
		vm.Panic(err.Error())
	}
	return vm.GetRTValuePool().CreateGOFunction(t, fn)
}

func CallBuiltinFunction(vm runtimeapi.VM, tpool *bctypes.TypePool, rawFn interface{}) {
	fnValue := reflect.ValueOf(rawFn)
	fnType := fnValue.Type()
	if fnType.Kind() != reflect.Func {
		vm.Panic("Cannot call Foreign function with non-function type!")
	}

	fnParamCount := fnType.NumIn()
	fnParamTypes := make([]reflect.Type, fnParamCount)
	for i := range fnParamCount {
		fnParamTypes[i] = fnType.In(i)
	}

	fnReturnCount := fnType.NumOut()
	fnReturnTypes := make([]reflect.Type, fnReturnCount)
	for i := range fnReturnCount {
		fnReturnTypes[i] = fnType.Out(i)
	}

	var arguments []reflect.Value

	paramIdxStart := 0
	if len(fnParamTypes) != 0 && fnParamTypes[0].String() == "runtimeapi.VM" {
		arguments = append(arguments, reflect.ValueOf(vm))
		paramIdxStart++
	}

	var err error
	for i := paramIdxStart; i < fnParamCount; i++ {
		paramType := fnParamTypes[i]
		if vm.GetValueStack().GetPointer() == -1 && fnType.IsVariadic() {
			arguments = append(arguments, reflect.Zero(paramType))
			continue
		}
		value := vm.GetValueStack().Pop()
		var arg reflect.Value
		if paramType.String() == "rtvalue.RTValue" {
			arg = reflect.ValueOf(value)
		} else {
			arg, err = RuntimeValueToReflectionValue(value, paramType)
			if err != nil {
				vm.Panic("Error converting stack value to native value!: " + err.Error())
				return
			}
		}
		arguments = append(arguments, arg)

	}

	vm.GetCallStack().Push(runtime.NewPseudoFrame(vm.GetCallStack().GetTop(), vm.GetScope(), "{ Native-Function }", goruntime.FuncForPC(fnValue.Pointer()).Name()))

	var returnValues []reflect.Value
	if fnType.IsVariadic() {
		returnValues = fnValue.CallSlice(arguments)
	} else {
		returnValues = fnValue.Call(arguments)
	}
	for i := range returnValues {
		returnValue := returnValues[i]
		returnType := fnReturnTypes[i]

		var runtimeValue rtvalue.RTValue
		if returnType.String() == "rtvalue.RTValue" {
			runtimeValue = returnValue.Interface().(rtvalue.RTValue)
		} else {
			runtimeValue, err = ReflectionValueToRuntimeValue(vm.GetRTValuePool(), tpool, returnValue)
			if err != nil {
				vm.Panic("Error converting native value to stack value!: " + err.Error())
				return
			}
		}

		vm.GetValueStack().Push(runtimeValue)
	}
}
