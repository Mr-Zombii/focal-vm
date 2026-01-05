package runtimeapi

type ValueTag uint8

const (
	ValueTagUnknown ValueTag = iota

	ValueTagBoolean

	ValueTagInt8
	ValueTagInt16
	ValueTagInt32
	ValueTagInt64

	ValueTagFloat32
	ValueTagFloat64

	ValueTagUTF8String

	ValueTagArray

	ValueTagFunction
	ValueTagNativeFunction
	ValueTagForeignFunction
	ValueTagScope
)

type Value interface {
	GetTag() ValueTag
	String() string
	GetRawValue() interface{}
}

type CallableValue interface {
	GetTag() ValueTag
	String() string
	Call(VM)
}
