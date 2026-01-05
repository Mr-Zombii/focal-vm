package runtimeapi

import "focal-vm/internal/bytecode/spec"

type ModuleCollection interface {
	AddDirectories(...string)
	AddArchives(...string)
	SearchForModule(string) (*spec.BCModule, error)
	SearchForModuleBytes(string) ([]byte, error)
}
