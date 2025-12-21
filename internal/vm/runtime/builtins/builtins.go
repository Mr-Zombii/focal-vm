package builtins

import (
	"fmt"
	"focal-lang/internal/vm/api"
	"focal-lang/internal/vm/runtime"
)

func Register(vm api.VM) {
	value := runtime.NewNativeFunction(func(v api.VM) {
		value := v.GetStack().PopValue()
		caller := v.GetCallStack().GetTopFrame()
		callerModule := caller.GetModule()

		fnName := caller.GetFunction().GetName()
		fmt.Println("Called from \"" + fnName + "\" in module \"" + callerModule.GetName() + "\"")
		fmt.Println(value)
	})
	vm.GetScope().SetLocal("print", value)
}
