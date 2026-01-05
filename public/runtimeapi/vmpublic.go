package runtimeapi

import (
	"focal-vm/internal/bytecode/spec"
	"focal-vm/internal/util"
	"plugin"
)

type OpcodeImpl func(VM, Frame)

type VM interface {
	GetLoadedModules() map[string]*spec.BCModule
	GetOpcodeMap() []OpcodeImpl
	InstallOpcodeMap([]OpcodeImpl)
	LoadModule(string) *spec.BCModule
	Run(string)
	GetCallStack() *util.Stack[Frame]
	GetValueStack() *util.Stack[Value]
	GetScope() Scope
	GetLoadedPlugins() map[string]*plugin.Plugin
	LoadPlugin(string) *plugin.Plugin
	Panic(string)
	Halt(int32)
	SetStopCallback(f func())
	GetModuleCollection() ModuleCollection
}
