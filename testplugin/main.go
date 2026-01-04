package main

import (
	"focal-vm/public/runtimeapi"
)

func main() {}

func RegisterLunnoNativeFunctions(vm runtimeapi.VM) {

}

type Hello struct {
	value1 int
	value2 float32
	value3 string
	value4 bool
}

//func Builtin_print(a string) Hello {
//	fmt.Println(a)
//	return Hello{value1: 10, value2: 20, value3: a, value4: false}
//}
