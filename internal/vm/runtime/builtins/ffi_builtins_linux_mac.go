//go:build darwin || freebsd || linux || netbsd || android

package builtins

import (
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/vm/runtime/ffi"
	"focal-vm/public/runtimeapi"
	"syscall"

	"github.com/ebitengine/purego"
)

func _builtin_load_foreign_linux_mac(vm runtimeapi.VM, foreign_object_name string) uintptr {
	ptr, err := purego.Dlopen(foreign_object_name, purego.RTLD_LAZY|purego.RTLD_LOCAL)
	if err != nil {
		vm.Panic(err.Error())
	}

	return ptr
}

func _builtin_call_sym_foreign_linux_mac(vm runtimeapi.VM, pointer uintptr, procName string, arguments []uintptr) uintptr {
	symbolPtr, err := purego.Dlsym(pointer, procName)
	if err != nil {
		vm.Panic(err.Error())
	}
	ret, _, errno := purego.SyscallN(symbolPtr, arguments...)
	if errno != 0 {
		vm.Panic(syscall.Errno(errno).Error())
	}

	return ret
}

func _builtin_free_foreign_linux_mac(vm runtimeapi.VM, pointer uintptr) {
	err := purego.Dlclose(pointer)
	if err != nil {
		vm.Panic(err.Error())
	}
}

func RegisterForeign(vm runtimeapi.VM, scope runtimeapi.Scope, tpool *bctypes.TypePool) {
	scope.DefineAndSet("_builtin_load_foreign", ffi.NewBuiltinFunction(vm, tpool, _builtin_load_foreign_linux_mac))
	scope.DefineAndSet("_builtin_call_sym_foreign", ffi.NewBuiltinFunction(vm, tpool, _builtin_call_sym_foreign_linux_mac))
	scope.DefineAndSet("_builtin_free_foreign", ffi.NewBuiltinFunction(vm, tpool, _builtin_free_foreign_linux_mac))
}
