package stack

import (
	"focal-vm/public/runtimeapi"
)

type CallStack struct {
	ptr   int32
	stack []runtimeapi.Frame
}

func NewCallStack() runtimeapi.CallStack {
	return &CallStack{ptr: -1, stack: make([]runtimeapi.Frame, CALL_STACK_SIZE)}
}

func (s *CallStack) GetPointer() int32 {
	return s.ptr
}

func (s *CallStack) GetTopFrame() runtimeapi.Frame {
	return s.stack[s.ptr]
}

func (s *CallStack) PushFrame(o runtimeapi.Frame) int32 {
	s.ptr++
	s.stack[s.ptr] = o
	return s.ptr
}

func (s *CallStack) PopFrame() runtimeapi.Frame {
	o := s.stack[s.ptr]
	s.ptr--
	return o
}
