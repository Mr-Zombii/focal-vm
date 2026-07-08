package runtimeapi

import (
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/vm/rtvalue"
)

type Scope interface {
	HasLocal(string) bool
	OwnsLocal(name string) bool

	GetLocal(string) (rtvalue.RTValue, error)
	SetLocal(string, rtvalue.RTValue) error
	DefineAndSet(string, rtvalue.RTValue)

	DefineLocal(string, bctypes.BCType)

	NewChildScope() Scope
	GetParent() Scope
	Reset()

	Visit(func(Scope, string, rtvalue.RTValue))
}
