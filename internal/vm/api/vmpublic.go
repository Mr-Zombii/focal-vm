package api

import (
	"focal-lang/internal/bytecode/spec"
)

type OpcodeImpl func(VM, Frame)

type VM interface {
	GetLoadedModules() map[string]*spec.BCModule
	GetOpcodeMap() []OpcodeImpl
	InstallOpcodeMap([]OpcodeImpl)
	AddModule(*spec.BCModule)
	LoadModule(string) *spec.BCModule
	Run(string)
	GetCallStack() CallStack
	GetStack() Stack
	GetScope() Scope
}
