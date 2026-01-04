package builtins

import (
	"focal-vm/internal/vm/runtime"
	"focal-vm/internal/vm/runtime/ffi"
	"focal-vm/public/runtimeapi"
)

func Register(vm runtimeapi.VM) {
	RegisterLunno(vm)

	pluginloader := ffi.NewForeignFunction(func(v runtimeapi.VM, pluginNameValue runtimeapi.Value, fnSymbolValue runtimeapi.Value) {
		pluginName := pluginNameValue.(*runtime.UTF8StringValue).GetValue()
		// loadedPlugin := v.LoadPlugin(pluginName)
		fnSymbol := fnSymbolValue.(*runtime.UTF8StringValue).GetValue()
		println(pluginName, fnSymbol)
		// lookup, err := loadedPlugin.Lookup(fnSymbol)
		// if err != nil {
		// 	v.Panic("Error loading plugin: " + err.Error())
		// 	return
		// }
		// ffi.CallForeignFunction(v, lookup)
	})
	vm.GetScope().SetLocal("_builtin_load_plugin", pluginloader)
}
