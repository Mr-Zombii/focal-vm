package builtins

import (
	"bufio"
	"fmt"
	"focal-vm/internal/vm/runtime/ffi"
	"focal-vm/public/runtimeapi"
	"math"
	"os"
	"strings"
)

func _builtin_print(stackValue runtimeapi.Value) {
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
	vm.Halt()
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

func RegisterLunno(vm runtimeapi.VM) {
	vm.GetScope().SetLocal("_builtin_print", ffi.NewForeignFunction(_builtin_print))
	vm.GetScope().SetLocal("_builtin_read_line", ffi.NewForeignFunction(_builtin_read_line))
	vm.GetScope().SetLocal("_builtin_panic", ffi.NewForeignFunction(_builtin_panic))

	vm.GetScope().SetLocal("_builtin_floor", ffi.NewForeignFunction(_builtin_floor))
	vm.GetScope().SetLocal("_builtin_ceil", ffi.NewForeignFunction(_builtin_ceil))

	vm.GetScope().SetLocal("_builtin_str_concat", ffi.NewForeignFunction(_builtin_str_concat))
	vm.GetScope().SetLocal("_builtin_strlen", ffi.NewForeignFunction(_builtin_strlen))
	vm.GetScope().SetLocal("_builtin_substr", ffi.NewForeignFunction(_builtin_substr))
	vm.GetScope().SetLocal("_builtin_char_at", ffi.NewForeignFunction(_builtin_char_at))
	vm.GetScope().SetLocal("_builtin_to_upper", ffi.NewForeignFunction(_builtin_to_upper))
	vm.GetScope().SetLocal("_builtin_to_lower", ffi.NewForeignFunction(_builtin_to_lower))
}
