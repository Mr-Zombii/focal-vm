package erroring

import (
	"fmt"
	"os"
	"strings"
)

type ErrorHandler interface {
	Panic(title string, args ...interface{})
}

type _ErrorHandler struct {
}

var GlobalErrorHandler ErrorHandler = &_ErrorHandler{}

func (e *_ErrorHandler) Panic(title string, args ...interface{}) {
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

	fmt.Println(title + ": ")
	fmt.Println(outStr)
	os.Exit(-1)
}
