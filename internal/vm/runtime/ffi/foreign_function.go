package ffi

//type ForeignFunctionValue struct {
//	function  interface{}
//	reflected reflect.Value
//	name      string
//}
//
//func NewForeignFunction(fn interface{}) *ForeignFunctionValue {
//	reflection := reflect.ValueOf(fn)
//	return &ForeignFunctionValue{function: fn, reflected: reflection, name: goruntime.FuncForPC(reflection.Pointer()).Name()}
//}
//
//func (v *ForeignFunctionValue) Call(vm runtimeapi.VM) {
//	CallForeignFunction(vm, v.function)
//}
//
//func (v *ForeignFunctionValue) GetTag() runtimeapi.ValueTag {
//	return runtimeapi.ValueTagForeignFunction
//}
//
//func (v *ForeignFunctionValue) GetFunction() interface{} {
//	return v.function
//}
//
//func (v *ForeignFunctionValue) String() string {
//	return "ForeignFunction{ " + v.name + " }"
//}
//
//func (v *ForeignFunctionValue) GetName() string {
//	return v.name
//}
//
//func (v *ForeignFunctionValue) GetReflection() reflect.Value {
//	return v.reflected
//}
//
//func (v *ForeignFunctionValue) GetRawValue() interface{} {
//	return v.function
//}
//
//func CallForeignFunction(vm runtimeapi.VM, rawFn interface{}) {
//	fnValue := reflect.ValueOf(rawFn)
//	fnType := fnValue.Type()
//	if fnType.Kind() != reflect.Func {
//		vm.Panic("Cannot call Foreign function with non-function type!")
//	}
//
//	fnParamCount := fnType.NumIn()
//	fnParamTypes := make([]reflect.Type, fnParamCount)
//	for i := range fnParamCount {
//		fnParamTypes[i] = fnType.In(i)
//	}
//
//	fnReturnCount := fnType.NumOut()
//	fnReturnTypes := make([]reflect.Type, fnReturnCount)
//	for i := range fnReturnCount {
//		fnReturnTypes[i] = fnType.Out(i)
//	}
//
//	var arguments []reflect.Value
//
//	paramIdxStart := 0
//	if len(fnParamTypes) != 0 && fnParamTypes[0].String() == "runtimeapi.VM" {
//		arguments = append(arguments, reflect.ValueOf(vm))
//		paramIdxStart++
//	}
//
//	var err error
//	for i := paramIdxStart; i < fnParamCount; i++ {
//		paramType := fnParamTypes[i]
//		if vm.GetValueStack().GetPointer() == -1 && fnType.IsVariadic() {
//			arguments = append(arguments, reflect.Zero(paramType))
//			continue
//		}
//		value := vm.GetValueStack().Pop()
//		var arg reflect.Value
//		if paramType.String() == "runtimeapi.Value" {
//			arg = reflect.ValueOf(value)
//		} else {
//			arg, err = RuntimeValueToReflectionValue(value, paramType)
//			if err != nil {
//				vm.Panic("Error converting stack value to native value!: " + err.Error())
//				return
//			}
//		}
//		arguments = append(arguments, arg)
//
//	}
//
//	vm.GetCallStack().Push(runtime.NewPseudoFrame(vm.GetCallStack().GetTop(), vm.GetScope(), "{ Native-Function }", goruntime.FuncForPC(fnValue.Pointer()).Name()))
//
//	var returnValues []reflect.Value
//	if fnType.IsVariadic() {
//		returnValues = fnValue.CallSlice(arguments)
//	} else {
//		returnValues = fnValue.Call(arguments)
//	}
//	for i := range returnValues {
//		returnValue := returnValues[i]
//		returnType := fnReturnTypes[i]
//
//		var runtimeValue runtimeapi.Value
//		if returnType.String() == "runtimeapi.Value" {
//			runtimeValue = returnValue.Interface().(runtimeapi.Value)
//		} else {
//			runtimeValue, err = ReflectionValueToRuntimeValue(returnValue)
//			if err != nil {
//				vm.Panic("Error converting native value to stack value!: " + err.Error())
//				return
//			}
//		}
//
//		vm.GetValueStack().Push(runtimeValue)
//	}
//}
