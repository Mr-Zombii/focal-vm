package builtins

import (
	"focal-vm/public/runtimeapi"
)

func Register(scope runtimeapi.Scope) {
	// RegisterLunno(scope)
	// RegisterFocal(scope)

	//pluginloader := ffi.NewForeignFunction(func(v runtimeapi.VM, pluginNameValue runtimeapi.Value, fnSymbolValue runtimeapi.Value) {
	//	pluginName := pluginNameValue.(*runtime.UTF8StringValue).GetValue()
	//	loadedPlugin := v.LoadPlugin(pluginName)
	//	fnSymbol := fnSymbolValue.(*runtime.UTF8StringValue).GetValue()
	//	lookup, err := loadedPlugin.Lookup(fnSymbol)
	//	if err != nil {
	//		v.Panic("Error loading plugin: " + err.Error())
	//		return
	//	}
	//	ffi.CallForeignFunction(v, lookup)
	//})
	//scope.DefineAndSet("_builtin_load_plugin", pluginloader)
}
