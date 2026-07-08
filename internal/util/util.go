package util

import (
	"unsafe"
)

var PointerSize = int32(unsafe.Sizeof(uintptr(0)))
