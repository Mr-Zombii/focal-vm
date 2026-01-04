package runtimeapi

type ValueTag uint8

const (
	ValueTagUnknown ValueTag = iota

	ValueTagBoolean

	ValueTagInt8
	ValueTagInt16
	ValueTagInt32
	ValueTagInt64
	ValueTagInt128 // Unused

	ValueTagFloat8  // Unused
	ValueTagFloat16 // Unused
	ValueTagFloat32
	ValueTagFloat64
	ValueTagFloat128 // Unused

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
