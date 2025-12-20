package main

import "focal-lang/internal/vm"

func main() {
	vm := vm.NewVM()
	vm.LoadModule("bootstrap")
	vm.Run("bootstrap")
}
