package stack

import (
	"focal-vm/internal/bytecode/constants"
	"focal-vm/internal/vm/runtime"
	"focal-vm/public/runtimeapi"
)

var STACK_SIZE int32 = 256
var CALL_STACK_SIZE int32 = 256

type Stack struct {
	ptr   int32
	stack []runtimeapi.Value
}

func NewStack() runtimeapi.Stack {
	return &Stack{ptr: -1, stack: make([]runtimeapi.Value, STACK_SIZE)}
}

func (s *Stack) GetPointer() int32 {
	return s.ptr
}

func (s *Stack) GetTopValue() runtimeapi.Value {
	return s.stack[s.ptr]
}

func (s *Stack) PushValue(o runtimeapi.Value) int32 {
	s.ptr++
	s.stack[s.ptr] = o
	return s.ptr
}

func (s *Stack) PushConst(o constants.Constant) int32 {
	s.ptr++
	s.stack[s.ptr] = runtime.ConstantToValue(o)
	return s.ptr
}

func (s *Stack) PopValue() runtimeapi.Value {
	o := s.stack[s.ptr]
	s.ptr--
	return o
}
