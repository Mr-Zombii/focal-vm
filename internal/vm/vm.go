package vm

import (
	"focal-lang/internal/bytecode/bcio"
	"focal-lang/internal/bytecode/spec"
	"focal-lang/internal/vm/api"
	"focal-lang/internal/vm/runtime"
	"focal-lang/internal/vm/runtime/builtins"
	"focal-lang/internal/vm/runtime/opload"
	"focal-lang/internal/vm/stack"
	"io"
	"os"
)

type VM struct {
	stack        api.Stack
	callStack    api.CallStack
	modMap       map[string]*spec.BCModule
	opcodeMap    []api.OpcodeImpl
	currentFrame api.Frame
	scope        api.Scope
}

func NewVM() api.VM {
	vm := &VM{
		stack:     stack.NewStack(),
		callStack: stack.NewCallStack(),
		modMap:    map[string]*spec.BCModule{},
		scope:     runtime.NewScope(),
	}

	opload.InstallOpcodes(vm)
	builtins.Register(vm)

	return vm
}

func (vm *VM) GetStack() api.Stack {
	return vm.stack
}

func (vm *VM) GetCallStack() api.CallStack {
	return vm.callStack
}

func (vm *VM) InstallOpcodeMap(opcodeMap []api.OpcodeImpl) {
	vm.opcodeMap = opcodeMap
}

func (vm *VM) GetLoadedModules() map[string]*spec.BCModule {
	return vm.modMap
}

func (vm *VM) GetOpcodeMap() []api.OpcodeImpl {
	return vm.opcodeMap
}

func (vm *VM) LoadModule(moduleName string) *spec.BCModule {
	mod, ok := vm.modMap[moduleName]
	if ok {
		return mod
	}
	f, exists := os.OpenFile(moduleName+".fbc", os.O_RDONLY, 0)
	if exists != nil {
		panic("Could not find module named \"" + moduleName + "\"")
	}

	in, _ := io.ReadAll(f)
	f.Close()

	reader := bcio.NewReader(in)
	module := reader.ReadModule()
	vm.AddModule(module)
	return module
}

func (vm *VM) AddModule(module *spec.BCModule) {
	vm.modMap[module.GetName()] = module
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
	frame := runtime.NewFrame(vm.scope, mod, fun)
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

func (vm *VM) GetScope() api.Scope {
	return vm.scope
}
