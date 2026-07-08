package gc

import (
	"fmt"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/internal/vm/runtime/allocator"
	"focal-vm/public/runtimeapi"
)

type GarbageCollector struct {
	G0 *rtvalue.RTValuePool
	G1 *rtvalue.RTValuePool

	G0_ObjectThreshold int
	G1_ObjectThreshold int

	iterations int32
}

func NewGarbageCollector(allocator allocator.Allocator) *GarbageCollector {
	return &GarbageCollector{
		G0: rtvalue.NewRTValuePool(allocator),
		G1: rtvalue.NewRTValuePool(allocator),

		G0_ObjectThreshold: 40,
		G1_ObjectThreshold: 40,
	}
}

func (gc *GarbageCollector) GetMainPool() *rtvalue.RTValuePool {
	return gc.G0
}

func (gc *GarbageCollector) Collect(vm runtimeapi.VM) {
	gc.collectSingleGen(vm, gc.G0, 0)
	//gc.collect(vm, gc.G0, gc.G1, gc.G0_ObjectThreshold, 3)
	//gc.collectSingleGen(vm, gc.G1, gc.G1_ObjectThreshold)
}

func (gc *GarbageCollector) collect(vm runtimeapi.VM, gen *rtvalue.RTValuePool, nextGen *rtvalue.RTValuePool, objThresh int, survivalThresh int32) {
	if gen.Length() >= objThresh {
		gen.ResetMarks()
		globalScope := vm.GetScope()

		stack := vm.GetValueStack()
		rtvalues := stack.GetInternalArray()

		for i := 0; i < int(stack.GetPointer()+1); i++ {
			value := rtvalues[i]
			value.Mark()
			value.Walk(func(v rtvalue.RTValue) {
				if !v.IsMarked() {
					v.Mark()
				}
			})
		}

		globalScope.Visit(func(scope runtimeapi.Scope, s string, value rtvalue.RTValue) {
			if !value.IsMarked() {
				value.Mark()
			}
		})

		callstack := vm.GetCallStack()
		frames := callstack.GetInternalArray()
		for i := 0; i < int(callstack.GetPointer()+1); i++ {
			frame := frames[i]
			scope := frame.GetScope()
			scope.Visit(func(scope runtimeapi.Scope, s string, value rtvalue.RTValue) {
				if !value.IsMarked() {
					value.Mark()
				}
			})
		}

		gen.CleanAndTransfer(survivalThresh, nextGen)
	}
}

func (gc *GarbageCollector) collectSingleGen(vm runtimeapi.VM, gen *rtvalue.RTValuePool, objThresh int) {
	if gen.Length() >= objThresh {
		gen.ResetMarks()
		globalScope := vm.GetScope()

		stack := vm.GetValueStack()
		rtvalues := stack.GetInternalArray()

		for i := 0; i < int(stack.GetPointer()+1); i++ {
			value := rtvalues[i]
			value.Mark()
			value.Walk(func(v rtvalue.RTValue) {
				if !v.IsMarked() {
					v.Mark()
				}
			})
		}

		globalScope.Visit(func(scope runtimeapi.Scope, s string, value rtvalue.RTValue) {
			if !value.IsMarked() {
				value.Mark()
			}
		})

		callstack := vm.GetCallStack()
		frames := callstack.GetInternalArray()
		for i := 0; i < int(callstack.GetPointer()+1); i++ {
			frame := frames[i]
			scope := frame.GetScope()
			scope.Visit(func(scope runtimeapi.Scope, s string, value rtvalue.RTValue) {
				if !value.IsMarked() {
					value.Mark()
				}
			})
		}

		gen.Clean()
	}
}

func (gc *GarbageCollector) PrintGens() {
	fmt.Println("G0: ", gc.G0.Length(), "value(s) allocated", gc.G0)
	fmt.Println("G1: ", gc.G1.Length(), "value(s) allocated", gc.G1)
}

var objectThreshold = 50

func GarbageCollect(vm runtimeapi.VM, runAnyway bool) {
	rtpool := vm.GetRTValuePool()
	if runAnyway || (rtpool.Length() >= objectThreshold) {
		rtpool.ResetMarks()

		stack := vm.GetValueStack()
		rtvalues := stack.GetInternalArray()

		for i := 0; i < int(stack.GetPointer()+1); i++ {
			value := rtvalues[i]
			value.Mark()
			value.Walk(func(v rtvalue.RTValue) {
				if !v.IsMarked() {
					v.Mark()
				}
			})
		}

		//vm.GetScope().Visit(func(owner runtimeapi.Scope, name string, value rtvalue.RTValue) {
		//	value.IncRefCount()
		//})

		oldLen := rtpool.Length()
		rtpool.Clean()
		fmt.Println(rtpool)
		fmt.Println(oldLen, " -> ", rtpool.Length(), "value(s) allocated")
	}
}
