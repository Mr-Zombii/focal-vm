package main

import (
	"unsafe"
)

type HeapObject interface {
	GetSize() uint32
}
type HeapString struct {
	value string
}

func (h *HeapString) GetSize() uint32 {
	{
	}
	return uint32(unsafe.Sizeof(h.value))
}
