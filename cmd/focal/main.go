package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"focal-vm/internal/erroring"
	"focal-vm/internal/vm"
	"io"
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
	}
	flagMap["-r"] = runFn
	flagMap["--run"] = runFn

	runArchiveFn := func() {
		archive := os.Args[2]
		zf, err := zip.OpenReader(archive)
		if err != nil {
			erroring.GlobalErrorHandler.Panic("CLI Module Launcher", err)
		}

		manifest, err := zf.Open("focal-manifest.json")
		if err != nil {
			erroring.GlobalErrorHandler.Panic("CLI Module Launcher", err)
		}
		manifestBytes, err := io.ReadAll(manifest)
		if err != nil {
			erroring.GlobalErrorHandler.Panic("CLI Module Launcher", err)
		}
		manifest.Close()
		zf.Close()

		var data map[string]any
		err = json.Unmarshal(manifestBytes, &data)
		if err != nil {
			panic(err)
		}
		if mainModule, ok := data["main-module"].(string); ok {
			fvm := vm.NewVM()
			//fmt.Printf("Loading module \"%s\"\n", mainModule)
			moduleCollection := fvm.GetModuleCollection()
			moduleCollection.AddArchives(archive)

			fvm.LoadModule(mainModule)
			fvm.Run(mainModule)
			fvm.GetModuleCollection()
			return
		}
		fmt.Printf("Focal Archive \"%s\" does not have the \"main-module\" attribute in the manifest!\n", archive)

	}
	flagMap["-a"] = runArchiveFn
	flagMap["--run-archive"] = runArchiveFn
}

func main() {
	registerFlags()
	arg1 := os.Args[1]
	flagMap[arg1]()
}
