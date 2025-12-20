package main

import (
	"fmt"
	"focal-lang/internal/bytecode/bcio"
	"io"
	"os"
)

func main() {
	f, _ := os.OpenFile("out.fbc", os.O_RDONLY, 0)
	in, _ := io.ReadAll(f)
	f.Close()

	reader := bcio.NewReader(in)
	module := reader.ReadModule()
	fmt.Println(module.GetName())
}
