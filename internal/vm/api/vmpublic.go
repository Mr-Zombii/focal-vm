package api

import (
	"focal-lang/internal/bytecode/constants"
	"focal-lang/internal/bytecode/spec"
)

type OpcodeImpl func(VM, Frame)

type VM interface {
	GetLoadedModules() map[string]*spec.BCModule
	GetOpcodeMap() []OpcodeImpl
	InstallOpcodeMap([]OpcodeImpl)
	AddModule(module *spec.BCModule)
	LoadModule(name string) *spec.BCModule
	Run(moduleName string)
	GetCallStack() CallStack
	GetStack() Stack
}

type CallStack interface {
	GetPointer() int32
	GetTopFrame() Frame
	PushFrame(Frame) int32
	PopFrame() Frame
}

type Stack interface {
	GetPointer() int32
	GetTopValue() Value
	PushValue(Value) int32
	PushConst(constants.Constant) int32
	PopValue() Value
}

type ValueTag uint8

type Value interface {
	GetTag() ValueTag
	String() string
}

type Frame interface {
	NewChildFrame(module *spec.BCModule, fn *spec.BCFunction) Frame
	LoadFn(fn *spec.BCFunction)
	GetCode() *[]uint8
	GetPtr() int32
	SetPtr(int32)
	GetConstantPool() *constants.ConstantPool
	GetFunction() *spec.BCFunction
	GetModule() *spec.BCModule
	LocalExist(n string) bool
	GetLocal(n string) Value
	SetLocal(n string, v Value)
}
