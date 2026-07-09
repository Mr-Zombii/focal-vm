package builtins

import (
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/public/runtimeapi"
)

func RegisterFocal(vm runtimeapi.VM, scope runtimeapi.Scope, tpool *bctypes.TypePool) {
	//scope.DefineAndSet("_builtin_load_module_bytes", ffi.NewForeignFunction(_builtin_load_module_bytes))
	//scope.DefineAndSet("_builtin_define_module", ffi.NewForeignFunction(_builtin_define_module))
	//scope.DefineAndSet("_builtin_get_module", ffi.NewForeignFunction(_builtin_get_module))
}
