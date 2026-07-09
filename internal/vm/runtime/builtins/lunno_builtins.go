package builtins

import (
	"bufio"
	"fmt"
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/internal/vm/runtime/ffi"
	"focal-vm/public/runtimeapi"
	"math"
	"os"
	"strings"
)

func _builtin_print(stackValue rtvalue.RTValue) {
	if stackValue == nil {
		print("nil")
		return
	}
	print(stackValue.String())
}

var bufioReader = bufio.NewReader(os.Stdin)

func _builtin_read_line(vm runtimeapi.VM) string {
	line, _, err := bufioReader.ReadLine()
	if err != nil {
		vm.Panic(err.Error())
	}
	return string(line)
}

func _builtin_panic(vm runtimeapi.VM, v string) {
	fmt.Print(v)
	vm.Halt(0)
}

func _builtin_floor(v float64) int {
	return int(math.Floor(v))
}

func _builtin_ceil(v float64) int {
	return int(math.Ceil(v))
}

func _builtin_str_concat(a string, b string) string {
	return a + b
}

func _builtin_strlen(v string) int {
	return len(v)
}

func _builtin_substr(v string, start, length int) string {
	return v[start : start+length]
}

func _builtin_char_at(v string, b byte) int {
	return strings.IndexByte(v, b)
}

func _builtin_to_upper(v string) string {
	return strings.ToUpper(v)
}

func _builtin_to_lower(v string) string {
	return strings.ToLower(v)
}

func RegisterLunno(vm runtimeapi.VM, scope runtimeapi.Scope, tpool *bctypes.TypePool) {
	scope.DefineAndSet("_builtin_print", ffi.NewBuiltinFunction(vm, tpool, _builtin_print))
	//scope.DefineAndSet("_builtin_read_line", ffi.NewBuiltinFunction(vm, tpool, _builtin_read_line))
	//scope.DefineAndSet("_builtin_panic", ffi.NewBuiltinFunction(vm, tpool, _builtin_panic))

	scope.DefineAndSet("_builtin_floor", ffi.NewBuiltinFunction(vm, tpool, _builtin_floor))
	scope.DefineAndSet("_builtin_ceil", ffi.NewBuiltinFunction(vm, tpool, _builtin_ceil))

	scope.DefineAndSet("_builtin_str_concat", ffi.NewBuiltinFunction(vm, tpool, _builtin_str_concat))
	scope.DefineAndSet("_builtin_strlen", ffi.NewBuiltinFunction(vm, tpool, _builtin_strlen))
	scope.DefineAndSet("_builtin_substr", ffi.NewBuiltinFunction(vm, tpool, _builtin_substr))
	scope.DefineAndSet("_builtin_char_at", ffi.NewBuiltinFunction(vm, tpool, _builtin_char_at))
	scope.DefineAndSet("_builtin_to_upper", ffi.NewBuiltinFunction(vm, tpool, _builtin_to_upper))
	scope.DefineAndSet("_builtin_to_lower", ffi.NewBuiltinFunction(vm, tpool, _builtin_to_lower))
}
