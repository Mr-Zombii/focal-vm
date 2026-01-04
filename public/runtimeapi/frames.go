package runtimeapi

import (
	"focal-vm/internal/bytecode/constants"
	"focal-vm/internal/bytecode/spec"
)

type Frame interface {
	NewChildFrame(*spec.BCModule, *spec.BCFunction) Frame
	LoadFn(*spec.BCFunction)
	GetCode() *[]uint8
	GetPtr() int32
	SetPtr(int32)
	GetConstantPool() *constants.ConstantPool

	GetFunctionName() string
	GetModuleName() string
	GetScope() Scope
	GetCaller() Frame
}
