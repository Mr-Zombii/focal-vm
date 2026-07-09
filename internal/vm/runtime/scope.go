package runtime

import (
	"fmt"
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/util"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/public/runtimeapi"
)

type Scope struct {
	parent        runtimeapi.Scope
	definedLocals map[string]bool
	localTypes    map[string]bctypes.BCType
	locals        map[string]rtvalue.RTValue
}

func NewScope() runtimeapi.Scope {
	return &Scope{definedLocals: map[string]bool{}, localTypes: map[string]bctypes.BCType{}, locals: map[string]rtvalue.RTValue{}}
}

func (s *Scope) NewChildScope() runtimeapi.Scope {
	scope := NewScope()
	scope.(*Scope).parent = s
	return scope
}

func (s *Scope) GetParent() runtimeapi.Scope {
	return s.parent
}

func (s *Scope) DefineAndSet(n string, v rtvalue.RTValue) {
	s.DefineLocal(n, v.GetType())
	s.SetLocal(n, v)
}

func (s *Scope) DefineLocal(n string, t bctypes.BCType) {
	s.definedLocals[n] = false
	s.localTypes[n] = t
	s.locals[n] = nil
}

func (s *Scope) HasLocal(name string) bool {
	if _, ok := s.definedLocals[name]; ok {
		return true
	}

	if s.parent == nil {
		return false
	}

	scopeStack := util.NewStack[runtimeapi.Scope](256)
	scopeStack.Push(s.parent)

	for scopeStack.GetPointer() != -1 {
		scope := scopeStack.Pop().(*Scope)
		if _, ok := scope.definedLocals[name]; ok {
			return true
		}
		if scope.parent != nil {
			scopeStack.Push(scope.GetParent())
		}
	}

	return false
}

func (s *Scope) OwnsLocal(name string) bool {
	_, ok := s.definedLocals[name]
	return ok
}

func (s *Scope) GetLocal(name string) (rtvalue.RTValue, error) {
	if _, ok := s.definedLocals[name]; ok {
		return s.locals[name], nil
	}

	if s.parent == nil {
		return nil, fmt.Errorf("Cannot get local '%s' as is not defined in parent nor local scope", name)
	}

	scopeStack := util.NewStack[runtimeapi.Scope](256)
	scopeStack.Push(s.parent)

	for scopeStack.GetPointer() != -1 {
		scope := scopeStack.Pop().(*Scope)
		if _, ok := scope.definedLocals[name]; ok {
			return scope.locals[name], nil
		}
		if scope.parent != nil {
			scopeStack.Push(scope.GetParent())
		}
	}

	return nil, fmt.Errorf("Cannot get local '%s' as is not defined in parent nor local scope", name)
}

func (s *Scope) SetLocal(name string, v rtvalue.RTValue) error {
	if _, ok := s.definedLocals[name]; ok {
		if a, okk := s.locals[name]; okk {
			if a != nil {
				a.DecRefCount()
			}
		}
		s.locals[name] = v
		return nil
	}

	if s.parent == nil {
		return fmt.Errorf("Cannot set local '%s' as is not defined in parent nor local scope", name)
	}

	scopeStack := util.NewStack[runtimeapi.Scope](256)
	scopeStack.Push(s.parent)

	for scopeStack.GetPointer() != -1 {
		scope := scopeStack.Pop().(*Scope)
		if _, ok := scope.definedLocals[name]; ok {
			if a, okk := scope.locals[name]; okk {
				if a != nil {
					a.DecRefCount()
				}
			}
			scope.locals[name] = v
			return nil
		}
		if scope.parent != nil {
			scopeStack.Push(scope.GetParent())
		}
	}

	return fmt.Errorf("Cannot set local '%s' as is not defined in parent nor local scope", name)
}

func (s *Scope) Reset() {
	for key := range s.locals {
		v := s.locals[key]
		if v != nil {
			v.DecRefCount()
		}
		delete(s.locals, key)
	}
	for key := range s.definedLocals {
		delete(s.definedLocals, key)
	}
	for key := range s.localTypes {
		delete(s.localTypes, key)
	}
}

func (s *Scope) Visit(visitor func(runtimeapi.Scope, string, rtvalue.RTValue)) {
	scopeStack := util.NewStack[runtimeapi.Scope](256)
	scopeStack.Push(s)

	for scopeStack.GetPointer() != -1 {
		scope := scopeStack.Pop().(*Scope)
		for n, v := range scope.locals {
			visitor(scope, n, v)
		}
		if scope.parent != nil {
			scopeStack.Push(scope.GetParent())
		}
	}
}
