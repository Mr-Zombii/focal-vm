package stack

import (
	"focal-lang/internal/vm/api"
)

type CallStack struct {
	ptr   int32
	stack []api.Frame
}

func NewCallStack() api.CallStack {
	return &CallStack{ptr: -1, stack: make([]api.Frame, CALL_STACK_SIZE)}
}

func (s *CallStack) GetPointer() int32 {
	return s.ptr
}

func (s *CallStack) GetTopFrame() api.Frame {
	return s.stack[s.ptr]
}

func (s *CallStack) PushFrame(o api.Frame) int32 {
	s.ptr++
	s.stack[s.ptr] = o
	return s.ptr
}

func (s *CallStack) PopFrame() api.Frame {
	o := s.stack[s.ptr]
	s.ptr--
	return o
}
