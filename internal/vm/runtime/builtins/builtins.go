package builtins

import (
	"focal-vm/internal/vm/runtime"
	"focal-vm/internal/vm/runtime/ffi"
	"focal-vm/public/runtimeapi"
)

func Register(vm runtimeapi.VM) {
	RegisterLunno(vm)

	pluginloader := runtime.NewNativeFunction(func(v runtimeapi.VM) {
		pluginName := v.GetStack().PopValue().(*runtime.UTF8StringValue).GetValue()
		loadedPlugin := v.LoadPlugin(pluginName)
		fnSymbol := v.GetStack().PopValue().(*runtime.UTF8StringValue).GetValue()
		lookup, err := loadedPlugin.Lookup(fnSymbol)
		if err != nil {
			v.Panic("Error loading plugin: " + err.Error())
			return
		}
		ffi.CallForeignFunction(v, lookup)
	})
	vm.GetScope().SetLocal("_builtin_load_plugin", pluginloader)
}
