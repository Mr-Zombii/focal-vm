package util

type ExpandingStack[T any] struct {
	ptr   int32
	stack []T
}

func NewExpandingStack[T any](size int32) *ExpandingStack[T] {
	return &ExpandingStack[T]{ptr: -1, stack: make([]T, size)}
}

func (s *ExpandingStack[T]) GetPointer() int32 {
	return s.ptr
}

func (s *ExpandingStack[T]) GetTop() T {
	return s.stack[s.ptr]
}

func (s *ExpandingStack[T]) Push(o T) int32 {
	s.ptr++
	if int(s.ptr) >= len(s.stack) {
		s.stack = append(s.stack, o)
	} else {
		s.stack[s.ptr] = o
	}
	return s.ptr
}

func (s *ExpandingStack[T]) Pop() T {
	o := s.stack[s.ptr]
	s.ptr--
	return o
}
