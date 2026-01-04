package main

import (
	"fmt"
	"math/big"
	"unsafe"
)

/* =======================
   Value definitions
   ======================= */

type Value interface{}

type (
	VFloat   struct{ V float64 }
	VInt     struct{ V *big.Int }
	VBool    struct{ V bool }
	VByte    struct{ V byte }
	VHeapRef struct{ ID int }
	VArray   struct{ V []Value }
	VTuple   struct{ V []Value }
	VUnit    struct{}
	VModule  struct{ V map[string]Value }
	VNative  struct{ Fn func([]Value) Value }
	VClosure struct{ Name string }
	VThunk   struct{ Fn func() Value }
	VRange   struct{ ID int }
)

/* =======================
   Heap objects
   ======================= */

type HeapObj interface{}

type (
	HString  struct{ V string }
	HArray   struct{ V []float64 }
	HBytes   struct{ V []byte }
	HClosure struct {
		Name string
		Code []Opcode
		Env  []EnvEntry
	}
	HRange struct {
		Current *big.Int
		Step    *big.Int
		End     *big.Int // nil if none
	}
)

type Opcode struct{} // placeholder

type EnvEntry struct {
	Name string
	Val  Value
	Mut  bool
}

/* =======================
   Generations
   ======================= */

type Generation struct {
	Objs   map[int]HeapObj
	Marked map[int]bool
}

/* =======================
   Globals
   ======================= */

var (
	internedStrings    = map[string]int{}
	newInternedStrings = map[string]int{}

	allocationCount     = 0
	allocationThreshold = 10_000

	youngGen = Generation{
		Objs:   make(map[int]HeapObj),
		Marked: make(map[int]bool),
	}
	oldGen = Generation{
		Objs:   make(map[int]HeapObj),
		Marked: make(map[int]bool),
	}

	nextID = 0
)

/* =======================
   Utilities
   ======================= */

func allocID() int {
	id := nextID
	nextID++
	return id
}

func objSize(v interface{}) int {
	return int(unsafe.Sizeof(v))
}

func heapMemory() (int, int) {
	sum := func(m map[int]HeapObj) int {
		total := 0
		for _, o := range m {
			total += objSize(o)
		}
		return total
	}
	return sum(youngGen.Objs), sum(oldGen.Objs)
}

func printHeapMemory() {
	y, o := heapMemory()
	fmt.Printf("Heap memory: young=%dB, old=%dB\n", y, o)
}

func heapSize() (int, int) {
	return len(youngGen.Objs), len(oldGen.Objs)
}

func printHeapStats() {
	y, o := heapSize()
	fmt.Printf("Heap: young=%d, old=%d\n", y, o)
}

/* =======================
   Allocation
   ======================= */

func allocInYoung(obj HeapObj) Value {
	id := allocID()
	youngGen.Objs[id] = obj
	return VHeapRef{ID: id}
}

func findHeapObj(id int) (HeapObj, bool) {
	if o, ok := youngGen.Objs[id]; ok {
		return o, true
	}
	o, ok := oldGen.Objs[id]
	return o, ok
}

/* =======================
   Accessors
   ======================= */

func getString(id int) (string, bool) {
	if o, ok := findHeapObj(id); ok {
		if s, ok := o.(HString); ok {
			return s.V, true
		}
	}
	return "", false
}

func getBytes(id int) ([]byte, bool) {
	if o, ok := findHeapObj(id); ok {
		if b, ok := o.(HBytes); ok {
			return b.V, true
		}
	}
	return nil, false
}

func getArray(id int) ([]float64, bool) {
	if o, ok := findHeapObj(id); ok {
		if a, ok := o.(HArray); ok {
			return a.V, true
		}
	}
	return nil, false
}

func getValue(id int) (Value, bool) {
	o, ok := findHeapObj(id)
	if !ok {
		return nil, false
	}

	switch v := o.(type) {
	case HString:
		arr := make([]Value, len(v.V))
		for i := range v.V {
			arr[i] = VByte{V: v.V[i]}
		}
		return VArray{V: arr}, true

	case HBytes:
		arr := make([]Value, len(v.V))
		for i := range v.V {
			arr[i] = VByte{V: v.V[i]}
		}
		return VArray{V: arr}, true

	case HArray:
		arr := make([]Value, len(v.V))
		for i, f := range v.V {
			arr[i] = VFloat{V: f}
		}
		return VArray{V: arr}, true

	case HClosure:
		return VClosure{Name: v.Name}, true

	case HRange:
		return VRange{ID: id}, true
	}

	return nil, false
}

/* =======================
   GC Mark & Sweep
   ======================= */

func mark(id int, gen *Generation) {
	if !gen.Marked[id] {
		gen.Marked[id] = true
	}
}

func markValue(v Value) {
	stack := []Value{v}

	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		switch x := n.(type) {
		case VHeapRef:
			if _, ok := youngGen.Objs[x.ID]; ok {
				mark(x.ID, &youngGen)
			} else if _, ok := oldGen.Objs[x.ID]; ok {
				mark(x.ID, &oldGen)
			}

		case VRange:
			if _, ok := youngGen.Objs[x.ID]; ok {
				mark(x.ID, &youngGen)
			} else if _, ok := oldGen.Objs[x.ID]; ok {
				mark(x.ID, &oldGen)
			}

		case VArray:
			stack = append(stack, x.V...)

		case VTuple:
			stack = append(stack, x.V...)
		}
	}
}

func markEnv(env []EnvEntry) {
	for _, e := range env {
		markValue(e.Val)
	}
}

func sweep(gen *Generation) {
	for id := range gen.Objs {
		if !gen.Marked[id] {
			delete(gen.Objs, id)
		}
	}
	gen.Marked = map[int]bool{}
}

/* =======================
   Promotion & GC
   ======================= */

func markAndPromote(roots []Value, env []EnvEntry) {
	youngGen.Marked = map[int]bool{}
	oldGen.Marked = map[int]bool{}

	for _, r := range roots {
		markValue(r)
	}
	markEnv(env)

	for _, id := range newInternedStrings {
		mark(id, &youngGen)
		mark(id, &oldGen)
	}

	for id, obj := range youngGen.Objs {
		if youngGen.Marked[id] {
			oldGen.Objs[id] = obj
			delete(youngGen.Objs, id)
		}
	}

	newInternedStrings = map[string]int{}
	sweep(&youngGen)
	sweep(&oldGen)
}

func maybeCollectGC(roots []Value, env []EnvEntry) {
	allocationCount++
	if allocationCount >= allocationThreshold {
		markAndPromote(roots, env)
		allocationCount = 0
	}
}

/* =======================
   Alloc helpers
   ======================= */

func allocStringWithGC(stack []Value, env []EnvEntry, s string) Value {
	if id, ok := internedStrings[s]; ok {
		return VHeapRef{ID: id}
	}

	maybeCollectGC(stack, env)
	id := allocID()

	youngGen.Objs[id] = HString{V: s}
	internedStrings[s] = id
	newInternedStrings[s] = id

	return VHeapRef{ID: id}
}

func allocBytesWithGC(stack []Value, env []EnvEntry, b []byte) Value {
	maybeCollectGC(stack, env)
	id := allocID()
	youngGen.Objs[id] = HBytes{V: b}
	return VHeapRef{ID: id}
}

func allocArrayWithGC(stack []Value, env []EnvEntry, a []float64) Value {
	maybeCollectGC(stack, env)
	return allocInYoung(HArray{V: a})
}

func allocRange(current, step, end *big.Int) Value {
	id := allocID()
	youngGen.Objs[id] = HRange{
		Current: current,
		Step:    step,
		End:     end,
	}
	return VHeapRef{ID: id}
}

func resetHeap() {
	youngGen.Objs = map[int]HeapObj{}
	youngGen.Marked = map[int]bool{}
	oldGen.Objs = map[int]HeapObj{}
	oldGen.Marked = map[int]bool{}
	newInternedStrings = map[string]int{}
	nextID = 0
}
