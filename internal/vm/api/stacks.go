package api

import "focal-lang/internal/bytecode/constants"

type CallStack interface {
	GetPointer() int32
	GetTopFrame() Frame
	PushFrame(Frame) int32
	PopFrame() Frame
}

type Stack interface {
	GetPointer() int32
	GetTopValue() Value
	PushValue(Value) int32
	PushConst(constants.Constant) int32
	PopValue() Value
}
