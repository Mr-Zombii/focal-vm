package main

import "focal-vm/internal/vm"

func main() {
	vm := vm.NewVM()
	vm.LoadModule("bootstrap")
	vm.Run("bootstrap")
}
