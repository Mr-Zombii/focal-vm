package runtime

import "focal-lang/internal/vm/api"

type Scope struct {
	parent api.Scope
	locals map[string]api.Value
}

func NewScope() api.Scope {
	return &Scope{locals: map[string]api.Value{}}
}

func (s *Scope) NewChildScope() api.Scope {
	scope := NewScope()
	scope.(*Scope).parent = s
	return scope
}

func (s *Scope) GetParent() api.Scope {
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

func (s *Scope) GetLocal(n string) api.Value {
	v, ok := s.locals[n]
	if !ok && s.GetParent() != nil {
		return s.GetParent().GetLocal(n)
	}
	return v
}

func (s *Scope) setLocalInternal(n string, v api.Value) {
	_, ok := s.locals[n]
	if !ok && s.parent != nil {
		s.GetParent().(*Scope).setLocalInternal(n, v)
		return
	}
	s.locals[n] = v
}

func (s *Scope) SetLocal(n string, v api.Value) {
	if s.HasLocal(n) {
		s.setLocalInternal(n, v)
		return
	}

	s.locals[n] = v
}
