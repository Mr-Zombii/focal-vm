package api

import (
	"focal-lang/internal/bytecode/constants"
	"focal-lang/internal/bytecode/spec"
)

type Frame interface {
	NewChildFrame(*spec.BCModule, *spec.BCFunction) Frame
	LoadFn(*spec.BCFunction)
	GetCode() *[]uint8
	GetPtr() int32
	SetPtr(int32)
	GetConstantPool() *constants.ConstantPool
	GetFunction() *spec.BCFunction
	GetModule() *spec.BCModule
	GetScope() Scope
}
