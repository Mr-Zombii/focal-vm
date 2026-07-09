package builtins

import (
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/public/runtimeapi"
)

func Register(vm runtimeapi.VM, scope runtimeapi.Scope, tpool *bctypes.TypePool) {
	RegisterLunno(vm, scope, tpool)
	RegisterForeign(vm, scope, tpool)
}
