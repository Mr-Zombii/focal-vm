package runtime

import (
	"focal-vm/public/runtimeapi"
)

type Scope struct {
	parent runtimeapi.Scope
	locals map[string]runtimeapi.Value
}

func (s *Scope) Reset() {
	for key := range s.locals {
		delete(s.locals, key)
	}
}

func NewScope() runtimeapi.Scope {
	return &Scope{locals: map[string]runtimeapi.Value{}}
}

func (s *Scope) NewChildScope() runtimeapi.Scope {
	scope := NewScope()
	scope.(*Scope).parent = s
	return scope
}

func (s *Scope) GetParent() runtimeapi.Scope {
	return s.parent
}

func (s *Scope) DefineLocal(n string) {
	s.locals[n] = nil
}

func (s *Scope) HasLocal(n string) bool {
	_, ok := s.locals[n]
	if !ok && s.parent != nil {
		return s.parent.HasLocal(n)
	}
	return ok
}

func (s *Scope) GetLocal(n string) runtimeapi.Value {
	v, ok := s.locals[n]
	if !ok && s.GetParent() != nil {
		return s.GetParent().GetLocal(n)
	}
	return v
}

func (s *Scope) setLocalInternal(n string, v runtimeapi.Value) {
	_, ok := s.locals[n]
	if !ok && s.parent != nil {
		s.GetParent().(*Scope).setLocalInternal(n, v)
		return
	}
	s.locals[n] = v
}

func (s *Scope) SetLocal(n string, v runtimeapi.Value) {
	if s.HasLocal(n) {
		s.setLocalInternal(n, v)
		return
	}

	s.locals[n] = v
}
