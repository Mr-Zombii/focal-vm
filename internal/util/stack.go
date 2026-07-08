package util

import (
	"runtime"
)

type Stack[T any] struct {
	ptr        int32
	stack      []T
	callerName string
	validNil   T
}

func NewStack[T any](size int32) *Stack[T] {
	pc, _, _, _ := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	stack := make([]T, size)
	return &Stack[T]{ptr: -1, stack: stack, callerName: details.Name(), validNil: stack[0]}
}

func (s *Stack[T]) GetInternalArray() []T {
	return s.stack
}

func (s *Stack[T]) GetPointer() int32 {
	return s.ptr
}

func (s *Stack[T]) GetTop() T {
	return s.stack[s.ptr]
}

func (s *Stack[T]) Push(o T) int32 {
	s.ptr++
	if s.ptr >= int32(len(s.stack)) {
		s.ptr = int32(len(s.stack) - 1)
		panic("stack overflow")
	}
	s.stack[s.ptr] = o
	return s.ptr
}

func (s *Stack[T]) Pop() T {
	if s.ptr < 0 {
		s.ptr = 0
		panic("stack underflow")
	}
	o := s.stack[s.ptr]
	s.stack[s.ptr] = s.validNil
	s.ptr--
	return o
}
