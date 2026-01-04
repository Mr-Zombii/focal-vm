package main

import (
	"fmt"
	"focal-vm/internal/vm"
	"os"
)

var flagMap = map[string]func(){}

func registerFlags() {
	helpFn := func() {
		fmt.Println("Usage: focal-vm --run {module}")
	}
	flagMap["-h"] = helpFn
	flagMap["--help"] = helpFn

	versionFn := func() {
		fmt.Println("Focal VM Version 1.0")
	}
	flagMap["-h"] = versionFn
	flagMap["--help"] = versionFn

	runFn := func() {
		fvm := vm.NewVM()
		module := os.Args[2]
		fvm.LoadModule(module)
		fvm.Run(module)
		return
	}
	flagMap["-r"] = runFn
	flagMap["--run"] = runFn
}

func main() {
	registerFlags()
	arg1 := os.Args[1]
	flagMap[arg1]()
}
