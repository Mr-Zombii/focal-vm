package api

type ValueTag uint8

type Value interface {
	GetTag() ValueTag
	String() string
}

type CallableValue interface {
	GetTag() ValueTag
	String() string
	Call(VM)
}
