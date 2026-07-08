package builtins

import (
	"focal-vm/public/runtimeapi"
)

//func _builtin_get_module_name(vm runtimeapi.VM) string {
//	return vm.GetCallStack().GetTop().GetCaller().GetModuleName()
//}
//
//func _builtin_load_module_bytes(vm runtimeapi.VM, moduleName string) []byte {
//	bytes, err := vm.GetModuleCollection().SearchForModuleBytes(moduleName)
//	if err != nil {
//		vm.Panic(err.Error())
//	}
//
//	return bytes
//}
//
//func _builtin_define_module(vm runtimeapi.VM, moduleName string, moduleBytes []byte) {
//	if _builtin_module_defined(vm, moduleName) {
//		vm.Panic(fmt.Sprintf("Module \"%v\" already defined in the VM!", moduleName))
//	}
//
//	reader := bcio.NewReader(moduleBytes)
//	module := reader.ReadModule()
//	vm.GetLoadedModules()[moduleName] = module
//}
//
//func _builtin_instance_module(vm runtimeapi.VM, moduleName string) runtimeapi.Value {
//	if !_builtin_module_defined(vm, moduleName) {
//		vm.Panic(fmt.Sprintf("Failed to instance module \"%v\", module was not defined!", moduleName))
//	}
//
//	module := vm.GetLoadedModules()[moduleName]
//	return moduleToObject(vm.GetScope().NewChildScope(), module)
//}
//
//func _builtin_unregistered_define_and_instance_module(vm runtimeapi.VM, moduleName string, moduleBytes []byte) runtimeapi.Value {
//	reader := bcio.NewReader(moduleBytes)
//	module := reader.ReadModule()
//	return moduleToObject(vm.GetScope().NewChildScope(), module)
//}
//
//func _builtin_module_defined(vm runtimeapi.VM, moduleName string) bool {
//	_, ok := vm.GetLoadedModules()[moduleName]
//	return ok
//}
//
//func _builtin_registered_define_and_instance_module(vm runtimeapi.VM, moduleName string, moduleBytes []byte) runtimeapi.Value {
//	if _builtin_module_defined(vm, moduleName) {
//		vm.Panic(fmt.Sprintf("Module \"%v\" already defined in the VM!", moduleName))
//		return nil
//	}
//
//	reader := bcio.NewReader(moduleBytes)
//	module := reader.ReadModule()
//	vm.GetLoadedModules()[moduleName] = module
//	return moduleToObject(vm.GetScope().NewChildScope(), module)
//}
//
//func _builtin_get_module(vm runtimeapi.VM, moduleName string) runtimeapi.Value {
//	module := vm.GetLoadedModules()[moduleName]
//	return moduleToObject(vm.GetScope().NewChildScope(), module)
//}
//
//func moduleToObject(scope runtimeapi.Scope, module *spec.BCModule) *runtime.ScopeValue {
//	moduleObject := runtime.NewScopeValue(scope)
//	moduleScope := moduleObject.GetScope()
//
//	moduleScope.DefineAndSet("this_module", moduleObject)
//
//	moduleScope.DefineAndSet("this_module_name", runtime.NewUTF8StringValue(module.GetName()))
//	moduleScope.DefineAndSet("this_module_vmregistered", runtime.BOOLEAN_VALUE_TRUE)
//
//	for _, fn := range module.GetFunctions() {
//		moduleScope.DefineLocal(fn.GetName())
//		moduleScope.SetLocal(fn.GetName(), runtime.NewFunction(moduleScope, fn))
//	}
//
//	return moduleObject
//}

func RegisterFocal(scope runtimeapi.Scope) {
	//scope.DefineAndSet("_builtin_load_module_bytes", ffi.NewForeignFunction(_builtin_load_module_bytes))
	//scope.DefineAndSet("_builtin_define_module", ffi.NewForeignFunction(_builtin_define_module))
	//scope.DefineAndSet("_builtin_get_module", ffi.NewForeignFunction(_builtin_get_module))
}
