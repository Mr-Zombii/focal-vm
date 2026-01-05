package util

type Stack[T any] struct {
	ptr   int32
	stack []T
}

func NewStack[T any](size int32) *Stack[T] {
	return &Stack[T]{ptr: -1, stack: make([]T, size)}
}

func (s *Stack[T]) GetPointer() int32 {
	return s.ptr
}

func (s *Stack[T]) GetTop() T {
	return s.stack[s.ptr]
}

func (s *Stack[T]) Push(o T) int32 {
	s.ptr++
	s.stack[s.ptr] = o
	return s.ptr
}

func (s *Stack[T]) Pop() T {
	o := s.stack[s.ptr]
	s.ptr--
	return o
}
