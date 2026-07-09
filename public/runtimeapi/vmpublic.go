package runtimeapi

import (
	"focal-vm/internal/bytecode/spec"
	"focal-vm/internal/util"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/internal/vm/runtime/allocator"
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
	GetValueStack() *util.Stack[rtvalue.RTValue]
	GetRTValuePool() *rtvalue.RTValuePool
	GetScope() Scope
	GetLoadedPlugins() map[string]*plugin.Plugin
	LoadPlugin(string) *plugin.Plugin
	Panic(string)
	Halt(int32)
	SetStopCallback(f func())
	GetModuleCollection() ModuleCollection
	GetAllocator() allocator.Allocator
}
