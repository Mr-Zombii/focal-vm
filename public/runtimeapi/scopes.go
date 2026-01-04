package runtimeapi

type Scope interface {
	HasLocal(string) bool
	GetLocal(string) Value
	SetLocal(string, Value)

	DefineLocal(string)

	NewChildScope() Scope
	GetParent() Scope
	Reset()
}
