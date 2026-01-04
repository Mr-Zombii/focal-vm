package runtimeapi

import (
	"focal-vm/internal/bytecode/spec"
	"plugin"
)

type OpcodeImpl func(VM, Frame)

type VM interface {
	GetLoadedModules() map[string]*spec.BCModule
	GetOpcodeMap() []OpcodeImpl
	InstallOpcodeMap([]OpcodeImpl)
	LoadModule(string) *spec.BCModule
	Run(string)
	GetCallStack() CallStack
	GetStack() Stack
	GetScope() Scope
	GetLoadedPlugins() map[string]*plugin.Plugin
	LoadPlugin(string) *plugin.Plugin
	Panic(string)
	Halt()
}
