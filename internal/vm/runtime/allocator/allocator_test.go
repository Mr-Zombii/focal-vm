package allocator

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestName(t *testing.T) {
	allocator := NewAllocator(8)

	ptrA := (*int32)(allocator.Alloc(4))
	*ptrA = 10
	ptrB := (*int32)(allocator.Alloc(4))
	*ptrB = 20

	allocator.Free(ptrA)
	allocator.Free(ptrB)

	ptrC := (*int64)(allocator.Alloc(8))
	*ptrC = 30

	fmt.Println(ptrC)

	if unsafe.Pointer(ptrC) != unsafe.Pointer(ptrA) {
		allocator.Free(ptrC)
		t.Fail()
	}

	allocator.Free(ptrC)
}
