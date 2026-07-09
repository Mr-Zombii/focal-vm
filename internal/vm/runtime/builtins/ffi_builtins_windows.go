//go:build windows

package builtins

import (
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/vm/runtime/ffi"
	"focal-vm/public/runtimeapi"
	"syscall"
)

func _builtin_load_foreign_windows(vm runtimeapi.VM, foreign_object_name string) uintptr {
	foreign_object_name = foreign_object_name + ".dll"

	handle, err := syscall.LoadLibrary(foreign_object_name)
	if err != nil {
		vm.Panic(err.Error())
	}

	return uintptr(handle)
}
func _builtin_call_sym_foreign_windows(vm runtimeapi.VM, pointer uintptr, procName string, arguments []uintptr) uintptr {
	handle := syscall.Handle(pointer)
	ptr, err := syscall.GetProcAddress(handle, procName)
	if err != nil {
		vm.Panic(err.Error())
	}
	ret, _, err := syscall.SyscallN(ptr, arguments...)
	if err != nil && err.Error() != "The operation completed successfully." {
		vm.Panic(err.Error())
	}

	return ret
}

func _builtin_free_foreign_windows(vm runtimeapi.VM, pointer uintptr) {
	err := syscall.FreeLibrary(syscall.Handle(pointer))
	if err != nil {
		vm.Panic(err.Error())
	}
}

func RegisterForeign(vm runtimeapi.VM, scope runtimeapi.Scope, tpool *bctypes.TypePool) {
	scope.DefineAndSet("_builtin_load_foreign", ffi.NewBuiltinFunction(vm, tpool, _builtin_load_foreign_windows))
	scope.DefineAndSet("_builtin_call_sym_foreign", ffi.NewBuiltinFunction(vm, tpool, _builtin_call_sym_foreign_windows))
	scope.DefineAndSet("_builtin_free_foreign", ffi.NewBuiltinFunction(vm, tpool, _builtin_free_foreign_windows))
	return
}
