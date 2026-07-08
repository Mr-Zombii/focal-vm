package runtime

import (
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/bytecode/constants"
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/bytecode/spec"
	"focal-vm/public/runtimeapi"
)

type Frame struct {
	parent       runtimeapi.Frame
	caller       runtimeapi.Frame
	moduleName   string
	functionName string
	cpool        *constants.ConstantPool
	tpool        *bctypes.TypePool
	code         []uint8
	ptr          int32

	scope runtimeapi.Scope
}

func NewPseudoFrame(caller runtimeapi.Frame, parentScope runtimeapi.Scope, moduleName string, functionName string) *Frame {
	frame := &Frame{}
	frame.caller = caller
	frame.scope = parentScope.NewChildScope()
	frame.moduleName = moduleName
	frame.functionName = functionName
	frame.code = []uint8{
		uint8(opcodes.OP_RET),
	}
	return frame
}

func NewFrame(caller runtimeapi.Frame, parentScope runtimeapi.Scope, module *spec.BCModule, fn *spec.BCFunction) runtimeapi.Frame {
	frame := &Frame{}
	frame.caller = caller
	frame.functionName = fn.GetName()
	frame.moduleName = module.GetName()
	frame.cpool = module.GetConstantPool()
	frame.tpool = module.GetTypePool()
	frame.scope = parentScope.NewChildScope()
	frame.LoadFn(fn)
	return frame
}

func (f *Frame) NewChildFrame(module *spec.BCModule, fn *spec.BCFunction) runtimeapi.Frame {
	frame := NewFrame(f, f.scope, module, fn).(*Frame)
	frame.parent = f
	return frame
}

func (f *Frame) LoadFn(fn *spec.BCFunction) {
	f.code = fn.GetCode()
	f.ptr = 0
	f.scope.Reset()
	f.functionName = fn.GetName()
	f.moduleName = fn.GetModule().GetName()
	f.cpool = fn.GetModule().GetConstantPool()
	f.tpool = fn.GetModule().GetTypePool()
}

func (f *Frame) GetCode() *[]uint8 {
	return &f.code
}

func (f *Frame) SetPtr(ptr int32) {
	f.ptr = ptr
}

func (f *Frame) GetPtr() int32 {
	return f.ptr
}

func (f *Frame) GetConstantPool() *constants.ConstantPool {
	return f.cpool
}

func (f *Frame) GetTypePool() *bctypes.TypePool {
	return f.tpool
}

func (f *Frame) GetModuleName() string {
	return f.moduleName
}

func (f *Frame) GetFunctionName() string {
	return f.functionName
}

func (f *Frame) GetScope() runtimeapi.Scope {
	return f.scope
}

func (f *Frame) GetCaller() runtimeapi.Frame {
	return f.caller
}
