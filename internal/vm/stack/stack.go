package stack

import (
	"focal-lang/internal/bytecode/constants"
	"focal-lang/internal/vm/api"
	"focal-lang/internal/vm/runtime"
)

var STACK_SIZE int32 = 256
var CALL_STACK_SIZE int32 = 256

type Stack struct {
	ptr   int32
	stack []api.Value
}

func NewStack() api.Stack {
	return &Stack{ptr: -1, stack: make([]api.Value, STACK_SIZE)}
}

func (s *Stack) GetPointer() int32 {
	return s.ptr
}

func (s *Stack) GetTopValue() api.Value {
	return s.stack[s.ptr]
}

func (s *Stack) PushValue(o api.Value) int32 {
	s.ptr++
	s.stack[s.ptr] = o
	return s.ptr
}

func (s *Stack) PushConst(o constants.Constant) int32 {
	s.ptr++
	s.stack[s.ptr] = runtime.ConstantToValue(o)
	return s.ptr
}

func (s *Stack) PopValue() api.Value {
	o := s.stack[s.ptr]
	s.ptr--
	return o
}
