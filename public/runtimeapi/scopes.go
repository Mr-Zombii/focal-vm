package runtimeapi

type Scope interface {
	HasLocal(string) bool
	OwnsLocal(name string) bool

	GetLocal(string) (Value, error)
	SetLocal(string, Value) error
	DefineAndSet(string, Value)

	DefineLocal(string)

	NewChildScope() Scope
	GetParent() Scope
	Reset()

	Visit(func(string, Value))
}
