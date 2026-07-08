package opcore

import (
	"fmt"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/public/runtimeapi"
)

func CheckFunction(vm runtimeapi.VM, value rtvalue.RTValue) {
	if value.GetTag() != rtvalue.RTValueTag_VMFUNCTION {
		vm.Panic(fmt.Sprintf("Stack value should be of function type, not type %s", value.GetType()))
	}
}
