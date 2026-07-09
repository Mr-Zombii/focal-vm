package builtins

import (
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/vm/runtime/ffi"
	"focal-vm/public/runtimeapi"
	"math"
)

func _builtin_f32_to_u32_bits(v float32) uint32 {
	return math.Float32bits(v)
}

func _builtin_f64_to_u64_bits(v float64) uint64 {
	return math.Float64bits(v)
}

func _builtin_u32_bits_to_f32(v uint32) float32 {
	return math.Float32frombits(v)
}

func _builtin_u64_bits_to_f64(v uint64) float64 {
	return math.Float64frombits(v)
}

func RegisterBitpacking(vm runtimeapi.VM, scope runtimeapi.Scope, tpool *bctypes.TypePool) {
	scope.DefineAndSet("_builtin_f32_to_u32_bits", ffi.NewBuiltinFunction(vm, tpool, _builtin_f32_to_u32_bits))
	scope.DefineAndSet("_builtin_f64_to_u64_bits", ffi.NewBuiltinFunction(vm, tpool, _builtin_f64_to_u64_bits))
	scope.DefineAndSet("_builtin_u32_bits_to_f32", ffi.NewBuiltinFunction(vm, tpool, _builtin_u32_bits_to_f32))
	scope.DefineAndSet("_builtin_u64_bits_to_f64", ffi.NewBuiltinFunction(vm, tpool, _builtin_u64_bits_to_f64))
}
