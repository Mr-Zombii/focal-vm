package opstruct

import (
	"fmt"
	"focal-vm/internal/vm/rtvalue"
	"focal-vm/public/runtimeapi"
)

func CheckStruct(vm runtimeapi.VM, value rtvalue.RTValue) {
	if value.GetTag() != rtvalue.RTValueTag_STRUCT {
		vm.Panic(fmt.Sprintf("Stack value should be a struct type, not type %s", value.GetType()))
	}
}
