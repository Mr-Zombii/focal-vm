package runtime

import (
	"focal-lang/internal/bytecode/constants"
	"focal-lang/internal/bytecode/spec"
	"focal-lang/internal/vm/api"
)

type Frame struct {
	parent api.Frame
	module *spec.BCModule
	fun    *spec.BCFunction
	cpool  *constants.ConstantPool
	locals map[string]api.Value
	code   []uint8
	ptr    int32

	scope api.Scope
}

func NewFrame(parentScope api.Scope, module *spec.BCModule, fn *spec.BCFunction) api.Frame {
	frame := &Frame{locals: map[string]api.Value{}}
	frame.fun = fn
	frame.module = module
	frame.cpool = module.GetConstantPool()
	frame.scope = parentScope.NewChildScope()
	frame.LoadFn(fn)
	return frame
}

func (f *Frame) NewChildFrame(module *spec.BCModule, fn *spec.BCFunction) api.Frame {
	frame := NewFrame(f.scope, module, fn).(*Frame)
	frame.parent = f
	return frame
}

func (f *Frame) LoadFn(fn *spec.BCFunction) {
	f.code = fn.GetCode()
	f.ptr = 0
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

func (f *Frame) GetFunction() *spec.BCFunction {
	return f.fun
}

func (f *Frame) GetModule() *spec.BCModule {
	return f.module
}

func (f *Frame) GetScope() api.Scope {
	return f.scope
}
