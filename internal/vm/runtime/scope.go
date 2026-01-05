package runtime

import (
	"fmt"
	"focal-vm/internal/util"
	"focal-vm/public/runtimeapi"
)

type Scope struct {
	parent        runtimeapi.Scope
	definedLocals map[string]bool
	locals        map[string]runtimeapi.Value
}

func NewScope() runtimeapi.Scope {
	return &Scope{definedLocals: map[string]bool{}, locals: map[string]runtimeapi.Value{}}
}

func (s *Scope) NewChildScope() runtimeapi.Scope {
	scope := NewScope()
	scope.(*Scope).parent = s
	return scope
}

func (s *Scope) GetParent() runtimeapi.Scope {
	return s.parent
}

func (s *Scope) DefineAndSet(n string, v runtimeapi.Value) {
	s.DefineLocal(n)
	s.SetLocal(n, v)
}

func (s *Scope) DefineLocal(n string) {
	s.definedLocals[n] = false
	s.locals[n] = nil
}

func (s *Scope) HasLocal(name string) bool {
	scopeStack := util.NewStack[runtimeapi.Scope](256)
	scopeStack.Push(s)

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

func (s *Scope) GetLocal(name string) (runtimeapi.Value, error) {
	scopeStack := util.NewStack[runtimeapi.Scope](256)
	scopeStack.Push(s)

	for scopeStack.GetPointer() != -1 {
		scope := scopeStack.Pop().(*Scope)
		if _, ok := scope.definedLocals[name]; ok {
			return scope.locals[name], nil
		}
		if scope.parent != nil {
			scopeStack.Push(scope.GetParent())
		}
	}

	return nil, fmt.Errorf("Cannot get local '%v' as is not defined in parent nor local scope", name)
}

func (s *Scope) SetLocal(name string, v runtimeapi.Value) error {
	scopeStack := util.NewStack[runtimeapi.Scope](256)
	scopeStack.Push(s)

	for scopeStack.GetPointer() != -1 {
		scope := scopeStack.Pop().(*Scope)
		if _, ok := scope.definedLocals[name]; ok {
			scope.locals[name] = v
		}
		if scope.parent != nil {
			scopeStack.Push(scope.GetParent())
		}
	}

	return fmt.Errorf("Cannot set local '%v' as is not defined in parent nor local scope", name)
}

func (s *Scope) Reset() {
	for key := range s.locals {
		delete(s.locals, key)
	}
	for key := range s.definedLocals {
		delete(s.definedLocals, key)
	}
}

func (s *Scope) Visit(visitor func(string, runtimeapi.Value)) {
	for k, value := range s.locals {
		if value != nil {
			visitor(k, value)
		}
	}
}
