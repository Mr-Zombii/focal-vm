package vm

import (
	"fmt"
	"focal-vm/internal/bytecode/bcio"
	"focal-vm/internal/bytecode/spec"
	"focal-vm/internal/vm/runtime"
	"focal-vm/internal/vm/runtime/builtins"
	"focal-vm/internal/vm/runtime/opload"
	"focal-vm/internal/vm/stack"
	"focal-vm/public/runtimeapi"
	"io"
	"os"
	"plugin"
	"strconv"
	"strings"
	"unicode/utf8"
)

type VM struct {
	stack        runtimeapi.Stack
	callStack    runtimeapi.CallStack
	modMap       map[string]*spec.BCModule
	opcodeMap    []runtimeapi.OpcodeImpl
	currentFrame runtimeapi.Frame
	scope        runtimeapi.Scope
	plugins      map[string]*plugin.Plugin
}

func NewVM() runtimeapi.VM {
	vm := &VM{
		stack:     stack.NewStack(),
		callStack: stack.NewCallStack(),
		modMap:    map[string]*spec.BCModule{},
		scope:     runtime.NewScope(),
		plugins:   map[string]*plugin.Plugin{},
	}

	opload.InstallOpcodes(vm)
	builtins.Register(vm)

	return vm
}

func (vm *VM) GetStack() runtimeapi.Stack {
	return vm.stack
}

func (vm *VM) GetCallStack() runtimeapi.CallStack {
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

	fmt.Println((strings.ReplaceAll(moduleName, ".", "/")) + ".fbc")
	f, exists := os.OpenFile((strings.ReplaceAll(moduleName, ".", "/"))+".fbc", os.O_RDONLY, 0)
	if exists != nil {
		panic("Could not find module named \"" + moduleName + "\"")
	}

	in, _ := io.ReadAll(f)
	f.Close()

	reader := bcio.NewReader(in)
	module := reader.ReadModule()
	module.SetName(moduleName)
	vm.modMap[moduleName] = module
	return module
}

func (vm *VM) Run(moduleName string) {
	mod, ok := vm.modMap[moduleName]
	if !ok {
		panic("Tried to load main function from non-existent module \"" + moduleName + "\"!")
	}
	fun := mod.GetFunction("main")
	if fun == nil {
		panic("Function main does not exist in module \"" + moduleName + "\"!")
	}
	frame := runtime.NewFrame(nil, vm.scope, mod, fun)
	vm.callStack.PushFrame(frame)

	for vm.callStack.GetPointer() != -1 {
		vm.currentFrame = vm.callStack.GetTopFrame()

		ptr := vm.currentFrame.GetPtr()
		code := *vm.currentFrame.GetCode()
		vm.currentFrame.SetPtr(ptr + 1)

		opcode := code[ptr]
		opcodeImpl := vm.opcodeMap[opcode]
		opcodeImpl(vm, vm.currentFrame)
	}

}

func (vm *VM) GetScope() runtimeapi.Scope {
	return vm.scope
}

func (vm *VM) ClearStacks() {
	for vm.GetStack().GetPointer() != -1 {
		vm.GetStack().PopValue()
	}
	for vm.GetCallStack().GetPointer() != -1 {
		vm.GetCallStack().PopFrame()
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

func flushBuffer() {
	_, err := os.Stdout.Write([]byte(errorOut))
	if err != nil {
		panic(err)
	}
	errorOut = ""
}

func (vm *VM) PrintBox(title string, contents string, boxChars string, partialOffset int32) int32 {
	boxRunes := []rune(boxChars)
	cornerTL := string(boxRunes[0])
	cornerTR := string(boxRunes[1])
	cornerBL := string(boxRunes[2])
	cornerBR := string(boxRunes[3])
	hbar := string(boxRunes[4])
	vbar := string(boxRunes[5])

	contentsBuffer := ""
	lines := strings.Split(contents, "\n")
	longestLine := utf8.RuneCountInString(title)
	for _, line := range lines {
		if utf8.RuneCountInString(line) > longestLine {
			longestLine = utf8.RuneCountInString(line)
		}
	}
	longestLine += 2

	leftLen := (longestLine) / 2
	rightLen := longestLine - leftLen

	offs := (int(partialOffset) / 2) - leftLen
	offStr := ""
	for range offs {
		offStr += " "
	}

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		lineOut := line
		pad := (longestLine - 2) - utf8.RuneCountInString(line)

		for range pad {
			lineOut += " "
		}
		contentsBuffer += offStr + vbar + " " + lineOut + " " + vbar + "\n"
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

	blankLine := offStr + vbar
	for range longestLine {
		blankLine += " "
	}
	blankLine += vbar + "\n"

	boxBuffer += offStr + cornerTL
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

	boxBuffer += offStr + cornerBL
	for range longestLine {
		boxBuffer += hbar
	}
	boxBuffer += cornerBR + "\n"

	print(boxBuffer)
	return int32(longestLine)
}

func (vm *VM) Panic(message string) {
	var stackFrames []runtimeapi.Frame
	for vm.callStack.GetPointer() != -1 {
		frame := vm.callStack.PopFrame()
		stackFrames = append(stackFrames, frame)
	}

	var stackValues []runtimeapi.Value
	for vm.stack.GetPointer() != -1 {
		value := vm.stack.PopValue()
		stackValues = append(stackValues, value)
	}

	lastFrameIdx := len(stackFrames) - 1
	printlnToBuf("┌[Call-Stack]")
	for i := range stackFrames {
		frame := stackFrames[lastFrameIdx-i]
		if i != lastFrameIdx {
			printToBuf("├")
		} else {
			printToBuf("└")
		}

		printToBuf("──> { Idx: #"+strconv.Itoa(i)+" Module: \"", frame.GetModuleName(), "\", Function: \"", frame.GetFunctionName(), "\" }\n")
	}
	if len(stackFrames) == 0 {
		printlnToBuf("└─(Empty Stack)")
	}

	printlnToBuf("")

	lastValueIdx := len(stackValues) - 1
	printlnToBuf("┌[Value-Stack]")
	for i := range stackValues {
		value := stackValues[lastValueIdx-i]
		if i != lastValueIdx {
			printToBuf("├")
		} else {
			printToBuf("└")
		}

		printToBuf("─> { Idx: #"+strconv.Itoa(i)+" Value: ", value.String(), " }\n")
	}
	if len(stackValues) == 0 {
		printlnToBuf("└──(Empty Stack)")
	}

	width := vm.PrintBox("[ Panic Dump ]", errorOut, "╔╗╚╝═║", 0)
	vm.PrintBox("[ Panic Message ]", message, "╭╮╰╯─│", width)

	vm.ClearStacks()

	os.Exit(-1)
}

func (vm *VM) Halt() {
	vm.ClearStacks()
	os.Exit(0)
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
