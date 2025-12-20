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
}

func NewFrame(module *spec.BCModule, fn *spec.BCFunction) api.Frame {
	frame := &Frame{locals: map[string]api.Value{}}
	frame.fun = fn
	frame.module = module
	frame.cpool = module.GetConstantPool()
	frame.LoadFn(fn)
	return frame
}

func (f *Frame) NewChildFrame(module *spec.BCModule, fn *spec.BCFunction) api.Frame {
	frame := NewFrame(module, fn).(*Frame)
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

func (f *Frame) LocalExist(n string) bool {
	_, ok := f.locals[n]
	if !ok && f.parent != nil {
		return f.parent.LocalExist(n)
	}
	return ok
}

func (f *Frame) GetLocal(n string) api.Value {
	v, ok := f.locals[n]
	if !ok && f.parent != nil {
		return f.parent.GetLocal(n)
	}
	return v
}

func (f *Frame) setLocalInternal(n string, v api.Value) {
	_, ok := f.locals[n]
	if !ok && f.parent != nil {
		f.parent.(*Frame).setLocalInternal(n, v)
		return
	}
	f.locals[n] = v
}

func (f *Frame) SetLocal(n string, v api.Value) {
	if f.LocalExist(n) {
		f.setLocalInternal(n, v)
		return
	}

	f.locals[n] = v
}
