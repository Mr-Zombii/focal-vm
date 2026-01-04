package main

import (
	"fmt"
	"math/big"
	"os"
)

/* =======================
   Test Harness
   ======================= */

func runTest(name string, f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[ERROR] %s: unexpected panic: %v\n", name, r)
			os.Exit(1)
		}
	}()

	f()
	fmt.Printf("Passed: %s\n", name)
}

/* =======================
   Tests
   ======================= */

func main() {
	fmt.Println("GC Testing")

	runTest("interned string allocation", func() {
		resetHeap()
		var stack []Value
		var env []EnvEntry

		v1 := allocStringWithGC(stack, env, "hello")
		v2 := allocStringWithGC(stack, env, "hello")

		r1, ok1 := v1.(VHeapRef)
		r2, ok2 := v2.(VHeapRef)
		if !ok1 || !ok2 {
			panic("Expected VHeapRef for string")
		}
		if r1.ID != r2.ID {
			panic("Interned string allocated multiple times")
		}
	})

	runTest("array allocation and retrieval", func() {
		resetHeap()
		var stack []Value
		var env []EnvEntry

		arr := []float64{1.0, 2.0, 3.0}
		v := allocArrayWithGC(stack, env, arr)

		ref, ok := v.(VHeapRef)
		if !ok {
			panic("Expected VHeapRef for array")
		}

		a, ok := getArray(ref.ID)
		if !ok || len(a) != 3 {
			panic("Failed to retrieve array from heap")
		}
	})

	runTest("range allocation", func() {
		resetHeap()

		r := allocRange(
			big.NewInt(0),
			big.NewInt(1),
			big.NewInt(10),
		)

		ref, ok := r.(VHeapRef)
		if !ok {
			panic("Expected VHeapRef for range")
		}

		v, ok := getValue(ref.ID)
		if !ok {
			panic("Range not correctly allocated")
		}

		if _, ok := v.(VRange); !ok {
			panic("Expected VRange value")
		}
	})

	runTest("young to old promotion", func() {
		resetHeap()
		var stack []Value

		env := []EnvEntry{
			{
				Name: "x",
				Val:  VInt{V: big.NewInt(1)},
				Mut:  false,
			},
		}

		v := allocStringWithGC(stack, env, "gc_test")

		for i := 0; i < allocationThreshold; i++ {
			_ = allocStringWithGC(stack, env, "dummy")
		}

		markAndPromote([]Value{v}, env)

		ref, ok := v.(VHeapRef)
		if !ok {
			panic("Expected VHeapRef for string")
		}

		if _, ok := oldGen.Objs[ref.ID]; !ok {
			panic("Object not promoted to old generation after GC")
		}
	})

	runTest("get_string returns correct value", func() {
		resetHeap()
		var stack []Value
		var env []EnvEntry

		v := allocStringWithGC(stack, env, "teststr")

		ref, ok := v.(VHeapRef)
		if !ok {
			panic("Expected VHeapRef")
		}

		s, ok := getString(ref.ID)
		if !ok || s != "teststr" {
			panic("get_string failed")
		}
	})

	runTest("bytes allocation and retrieval", func() {
		resetHeap()
		var stack []Value
		var env []EnvEntry

		b := []byte{'a', 'b', 'c'}
		v := allocBytesWithGC(stack, env, b)

		ref, ok := v.(VHeapRef)
		if !ok {
			panic("Expected VHeapRef")
		}

		arr, ok := getBytes(ref.ID)
		if !ok || string(arr) != "abc" {
			panic("get_bytes failed")
		}
	})

	runTest("array to VArray conversion", func() {
		resetHeap()
		var stack []Value
		var env []EnvEntry

		arr := []float64{1.1, 2.2}
		v := allocArrayWithGC(stack, env, arr)

		ref, ok := v.(VHeapRef)
		if !ok {
			panic("Expected VHeapRef")
		}

		val, ok := getValue(ref.ID)
		if !ok {
			panic("getValue failed")
		}

		va, ok := val.(VArray)
		if !ok || len(va.V) != 2 {
			panic("Array to VArray conversion failed")
		}

		f1 := va.V[0].(VFloat).V
		f2 := va.V[1].(VFloat).V
		if f1 != 1.1 || f2 != 2.2 {
			panic("Array values incorrect")
		}
	})

	runTest("marking nested VArray", func() {
		resetHeap()
		var stack []Value
		var env []EnvEntry

		inner := allocStringWithGC(stack, env, "inner")
		v := allocArrayWithGC(stack, env, []float64{1.0, 2.0})

		markValue(VArray{V: []Value{inner, v}})
	})

	runTest("range allocation with no end", func() {
		resetHeap()

		v := allocRange(
			big.NewInt(0),
			big.NewInt(1),
			nil,
		)

		ref, ok := v.(VHeapRef)
		if !ok {
			panic("Expected VHeapRef")
		}

		val, ok := getValue(ref.ID)
		if !ok {
			panic("getValue failed")
		}

		if _, ok := val.(VRange); !ok {
			panic("Range with nil end failed")
		}
	})

	runTest("large allocation stress test", func() {
		resetHeap()
		var stack []Value
		var env []EnvEntry

		num := 500_000
		refs := make([]Value, 0, num)

		for i := 1; i <= num; i++ {
			s := fmt.Sprintf("str_%d", i)
			v := allocStringWithGC(stack, env, s)
			refs = append(refs, v)
		}

		for _, v := range refs {
			markValue(v)
		}

		markAndPromote(refs, env)
	})

	fmt.Println("All GC tests passed.")
}
