package vm

import (
	"fmt"
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/bytecode/constants"
	"focal-vm/internal/bytecode/spec"
	"focal-vm/internal/erroring"
	"focal-vm/internal/util"
	"focal-vm/internal/vm/gc"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/internal/vm/runtime"
	"focal-vm/internal/vm/runtime/allocator"
	"focal-vm/internal/vm/runtime/builtins"
	"focal-vm/internal/vm/runtime/opload"
	"focal-vm/public/runtimeapi"
	"os"
	"plugin"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/mitchellh/go-wordwrap"
)

type VM struct {
	stack             *util.Stack[rtvalue.RTValue]
	callStack         *util.Stack[runtimeapi.Frame]
	modMap            map[string]*spec.BCModule
	opcodeMap         []runtimeapi.OpcodeImpl
	currentFrame      runtimeapi.Frame
	scope             runtimeapi.Scope
	plugins           map[string]*plugin.Plugin
	moduleCollection  runtimeapi.ModuleCollection
	garbageCollector  *gc.GarbageCollector
	allocator         allocator.Allocator
	haltCallback      func()
	firstModuleLoaded bool
}

func NewVM() runtimeapi.VM {
	heap := allocator.NewAllocator(0)

	vm := &VM{
		stack:             util.NewStack[rtvalue.RTValue](256),
		callStack:         util.NewStack[runtimeapi.Frame](256),
		modMap:            map[string]*spec.BCModule{},
		scope:             runtime.NewScope(),
		plugins:           map[string]*plugin.Plugin{},
		moduleCollection:  NewModuleCollection(),
		garbageCollector:  gc.NewGarbageCollector(heap),
		allocator:         heap,
		haltCallback:      func() {},
		firstModuleLoaded: false,
	}

	opload.InstallOpcodes(vm)

	return vm
}

func (vm *VM) GetValueStack() *util.Stack[rtvalue.RTValue] {
	return vm.stack
}

func (vm *VM) GetCallStack() *util.Stack[runtimeapi.Frame] {
	return vm.callStack
}

func (vm *VM) InstallOpcodeMap(opcodeMap []runtimeapi.OpcodeImpl) {
	vm.opcodeMap = opcodeMap
}

func (vm *VM) GetLoadedModules() map[string]*spec.BCModule {
	return vm.modMap
}

func (vm *VM) GetOpcodeMap() []runtimeapi.OpcodeImpl {
	return vm.opcodeMap
}

func (vm *VM) LoadModule(moduleName string) *spec.BCModule {
	if mod, ok := vm.modMap[moduleName]; ok {
		return mod
	}
	module, err := vm.moduleCollection.SearchForModule(moduleName)
	if err != nil {
		panic("Could not find module named \"" + moduleName + "\"")
	}

	vm.modMap[moduleName] = module
	if !vm.firstModuleLoaded {
		vm.firstModuleLoaded = true
		builtins.Register(vm, vm.scope, module.GetTypePool())
	}
	return module
}

func (vm *VM) GetRTValuePool() *rtvalue.RTValuePool {
	return vm.garbageCollector.GetMainPool()
}

type _VmErrorHandler struct {
	vm *VM
}

func (e *_VmErrorHandler) Panic(title string, args ...interface{}) {
	argStr := fmt.Sprint(args...)
	outStr := ""

	a := strings.Split(argStr, "\n")
	aLen := len(a) - 1

	for i, v := range a {
		outStr += "    " + v
		if i < aLen {
			outStr += "\n"
		}
	}

	e.vm.Panic(title + ":\n" + outStr)
}

func (vm *VM) Run(moduleName string) {
	//defer func() { // Error Handler
	//	if v := recover(); v != nil {
	//		vm.Panic(fmt.Sprint(v))
	//		return
	//	}
	//}()
	erroring.GlobalErrorHandler = &_VmErrorHandler{
		vm: vm,
	}
	mod, ok := vm.modMap[moduleName]
	if !ok {
		panic("Tried to load main function from non-existent module \"" + moduleName + "\"!")
	}
	tpool := mod.GetTypePool()
	cpool := mod.GetConstantPool()
	mainFnSignature := "function[array[string]][]"
	mainFnType, err := tpool.GetTypeFromSignature(mainFnSignature)
	if err != nil {
		vm.Panic(fmt.Sprintf("module '%s' could not find a function with the signature '%s', please make a function named 'main' with the signature '%s' or launch a different module!", moduleName, mainFnSignature, mainFnSignature))
	}
	fun, err := mod.GetFunction("main", mainFnType.(*bctypes.FunctionType))
	if err != nil {
		vm.Panic(err.Error())
		return
	}
	frame := runtime.NewFrame(nil, vm.scope, mod, fun)
	args := os.Args[3:len(os.Args)]
	rtV, _ := tpool.GetOrCreateUTFStringType()
	rtpool := vm.GetRTValuePool()

	value := rtpool.CreateArray(rtV, int32(len(args)))
	backing := value.GetValue()
	for i, v := range args {
		backing[i] = rtpool.GetOrMakeRTValueString(v)
	}
	scope := frame.GetScope()
	paramName := cpool.ExpectConstant(fun.GetParamNameIndexes()[0], constants.ConstantTagUTF8String).(*constants.ConstantUTF8String).GetValue()
	scope.DefineAndSet(paramName, value)

	vm.callStack.Push(frame)

	for vm.callStack.GetPointer() != -1 {
		vm.currentFrame = vm.callStack.GetTop()

		ptr := vm.currentFrame.GetPtr()
		code := *vm.currentFrame.GetCode()
		vm.currentFrame.SetPtr(ptr + 1)

		opcode := code[ptr]
		opcodeImpl := vm.opcodeMap[opcode]
		if opcodeImpl == nil {
			panic("Unimplemented opcode " + strconv.Itoa(int(opcode)))
		}
		opcodeImpl(vm, vm.currentFrame)
		//vm.garbageCollector.Collect(vm)
	}

	//vm.garbageCollector.PrintGens()
	vm.Panic("Exit")
	vm.Halt(0)
}

func (vm *VM) GetScope() runtimeapi.Scope {
	return vm.scope
}

func (vm *VM) ResetStackPointers() {
	for vm.GetValueStack().GetPointer() != -1 {
		vm.GetValueStack().Pop()
	}
	for vm.GetCallStack().GetPointer() != -1 {
		vm.GetCallStack().Pop()
	}
}

var errorOut = ""

func printlnToBuf(strings ...string) {
	printToBuf(strings...)
	printToBuf("\n")
}

func printToBuf(strings ...string) {
	out := ""
	for _, v := range strings {
		out += v
	}
	errorOut += out
}

func (vm *VM) printBox(title string, contents string, boxChars string) (string, int32) {
	boxRunes := []rune(boxChars)
	cornerTL := string(boxRunes[0])
	cornerTR := string(boxRunes[1])
	cornerBL := string(boxRunes[2])
	cornerBR := string(boxRunes[3])
	hbar := string(boxRunes[4])
	vbar := string(boxRunes[5])

	contentsBuffer := ""
	lines := strings.Split(strings.ReplaceAll(contents, "\t", strings.Repeat(" ", 4)), "\n")
	longestLine := utf8.RuneCountInString(title)
	for _, line := range lines {
		if utf8.RuneCountInString(line) > longestLine {
			longestLine = utf8.RuneCountInString(line)
		}
	}
	longestLine += 2

	leftLen := (longestLine) / 2
	rightLen := longestLine - leftLen

	for _, line := range lines {
		lineOut := line
		pad := (longestLine - 2) - utf8.RuneCountInString(line)

		lineOut += strings.Repeat(" ", pad)
		contentsBuffer += vbar + " " + lineOut + " " + vbar + "\n"
	}
	boxBuffer := ""

	titleLen := utf8.RuneCountInString(title)
	titleLenL := titleLen / 2
	titleLenR := titleLen - titleLenL
	if leftLen > rightLen {
		leftLen -= titleLenL
		rightLen -= titleLenR
	} else {
		rightLen -= titleLenR
		leftLen -= titleLenL
	}

	blankLine := vbar
	for range longestLine {
		blankLine += " "
	}
	blankLine += vbar + "\n"

	boxBuffer += cornerTL
	for range leftLen {
		boxBuffer += hbar
	}
	boxBuffer += title
	for range rightLen {
		boxBuffer += hbar
	}
	boxBuffer += cornerTR + "\n"

	boxBuffer += blankLine
	boxBuffer += contentsBuffer
	boxBuffer += blankLine

	boxBuffer += cornerBL
	for range longestLine {
		boxBuffer += hbar
	}
	boxBuffer += cornerBR + "\n"

	return boxBuffer, int32(longestLine)
}

func (vm *VM) Panic(message string) {
	errorOut = ""

	var stackFrames []runtimeapi.Frame
	for vm.callStack.GetPointer() != -1 {
		frame := vm.callStack.Pop()
		stackFrames = append(stackFrames, frame)
	}

	var stackValues []rtvalue.RTValue
	for vm.stack.GetPointer() != -1 {
		value := vm.stack.Pop()
		stackValues = append(stackValues, value)
	}

	printlnToBuf("┌[Allocator]")
	//printToBuf("├")
	printToBuf("└")

	printToBuf("──> { Backing Memory Size: " + strconv.Itoa(vm.allocator.GetTotalBackingSize()) + " }")

	printlnToBuf()
	printlnToBuf()

	lastFrameIdx := len(stackFrames) - 1
	printlnToBuf("┌[Call-Stack]")
	for i := range stackFrames {
		frame := stackFrames[lastFrameIdx-i]
		if i != lastFrameIdx {
			printToBuf("├")
		} else {
			printToBuf("└")
		}

		printToBuf("──> { Idx: #"+strconv.Itoa(i)+" Module: \"", frame.GetModuleName(), "\", Function: \"", frame.GetFunctionName(), "\" }")

		if i != lastFrameIdx {
			printlnToBuf()
		}
	}
	if len(stackFrames) == 0 {
		printlnToBuf("└─(Empty Stack)")
	}

	printlnToBuf()

	lastValueIdx := len(stackValues) - 1
	printlnToBuf("┌[Value-Stack]")
	for i := range stackValues {
		value := stackValues[lastValueIdx-i]
		if i != lastValueIdx {
			printToBuf("├")
		} else {
			printToBuf("└")
		}

		if value == nil {
			printToBuf("─> { Idx: #" + strconv.Itoa(i) + ", Nil Stack Value }")
		} else {
			if value.GetTag() == 0 {
				printToBuf("─> { Idx: #" + strconv.Itoa(i) + ", Broken Stack Value }")
			} else {
				printToBuf("─> { Idx: #"+strconv.Itoa(i)+", Type: "+value.GetType().String()+" Value: ", strconv.Quote(value.String()), " }")
				value.DecRefCount()
			}
		}

		if i != lastValueIdx {
			printlnToBuf()
		}
	}
	if len(stackValues) == 0 {
		printToBuf("└──(Empty Stack)")
	}

	box1, width1 := vm.printBox("[ Panic Dump ]", errorOut, "╔╗╚╝═║")
	box2, width2 := vm.printBox("[ Panic Message ]", wordwrap.WrapString(message, 125), "╭╮╰╯─│")

	if width1 > width2 {
		box2 = vm.indentStr(vm.getCenter(width2, width1), box2)
	} else if width2 > width1 {
		box1 = vm.indentStr(vm.getCenter(width1, width2), box1)
	}

	print(box1)
	print(box2)

	vm.Halt(-1)
}

func (vm *VM) getCenter(width1 int32, width2 int32) int32 {
	leftLen := (width1) / 2
	offs := (width2 / 2) - leftLen
	return offs
}

func (vm *VM) indentStr(identWidth int32, v string) string {
	strBuffer := ""
	lines := strings.Split(v, "\n")
	lineCount := len(lines)
	for i, line := range lines {
		if strings.TrimRight(line, " ") == "" {
			continue
		}
		strBuffer += strings.Repeat(" ", int(identWidth)) + line
		if i != lineCount-1 {
			strBuffer += "\n"
		}
	}
	return strBuffer
}

func (vm *VM) Halt(exitCode int32) {
	vm.ResetStackPointers()
	vm.haltCallback()
	os.Exit(int(exitCode))
}

func (vm *VM) LoadPlugin(pluginName string) *plugin.Plugin {
	if p, ok := vm.plugins[pluginName]; ok {
		return p
	}

	p, err := plugin.Open(pluginName)
	if err != nil {
		vm.Panic(err.Error())
	}
	vm.plugins[pluginName] = p
	return p
}

func (vm *VM) GetLoadedPlugins() map[string]*plugin.Plugin {
	return vm.plugins
}

func (vm *VM) SetStopCallback(f func()) {
	vm.haltCallback = f
}

func (vm *VM) GetModuleCollection() runtimeapi.ModuleCollection {
	return vm.moduleCollection
}

func (vm *VM) freePool() {
	v := vm.GetRTValuePool().GetValues()
	for _, a := range v {
		if vm.allocator.IsInvalidOrFree(v) {
			erroring.GlobalErrorHandler.Panic("RTValue Lifetime Error", "RTValue that was already free has been kept and freed twice!!")
			return
		}
		a.OnFree()
		vm.allocator.Free(a)
	}
}
