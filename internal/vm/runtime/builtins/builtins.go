package builtins

import (
	"fmt"
	"focal-lang/internal/vm/api"
	"focal-lang/internal/vm/runtime"
)

func Register(vm api.VM) {
	value := runtime.NewNativeFunction(func(v api.VM) {
		value := v.GetStack().PopValue()
		fmt.Println(value)
		fmt.Println(value)
	})
	vm.GetScope().SetLocal("print", value)
}
