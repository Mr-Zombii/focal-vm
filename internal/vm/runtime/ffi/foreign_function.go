package ffi

import (
	"focal-vm/internal/vm/runtime"
	"focal-vm/public/runtimeapi"
	"reflect"
	goruntime "runtime"
)

type ForeignFunctionValue struct {
	function interface{}
}

func NewForeignFunction(fn interface{}) *ForeignFunctionValue {
	return &ForeignFunctionValue{function: fn}
}

func (v *ForeignFunctionValue) Call(vm runtimeapi.VM) {
	CallForeignFunction(vm, v.function)
}

func (v *ForeignFunctionValue) GetTag() runtimeapi.ValueTag {
	return runtimeapi.ValueTagForeignFunction
}

func (v *ForeignFunctionValue) GetFunction() interface{} {
	return v.function
}

func (v *ForeignFunctionValue) String() string {
	return "Foreign"
}

func (v *ForeignFunctionValue) GetRawValue() interface{} {
	return v.function
}

func CallForeignFunction(vm runtimeapi.VM, rawFn interface{}) {
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
	var arguments []reflect.Value

	paramIdxStart := 0
	if len(fnParamTypes) != 0 && fnParamTypes[0].String() == "runtimeapi.VM" {
		arguments = append(arguments, reflect.ValueOf(vm))
		paramIdxStart++
	}

	var err error
	for i := paramIdxStart; i < fnParamCount; i++ {
		paramType := fnParamTypes[i]
		if vm.GetStack().GetPointer() == -1 && fnType.IsVariadic() {
			arguments = append(arguments, reflect.Zero(paramType))
			continue
		}
		value := vm.GetStack().PopValue()
		var arg reflect.Value
		if paramType.String() == "runtimeapi.Value" {
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

	vm.GetCallStack().PushFrame(runtime.NewPseudoFrame(vm.GetCallStack().GetTopFrame(), vm.GetScope(), "{ Native-Function }", goruntime.FuncForPC(fnValue.Pointer()).Name()))

	var returnValues []reflect.Value
	if fnType.IsVariadic() {
		returnValues = fnValue.CallSlice(arguments)
	} else {
		returnValues = fnValue.Call(arguments)
	}
	for i := range returnValues {
		returnValue := returnValues[i]
		runtimeValue, err := ReflectionValueToRuntimeValue(returnValue)
		if err != nil {
			vm.Panic("Error converting native value to stack value!: " + err.Error())
		}
		vm.GetStack().PushValue(runtimeValue)
	}
}
