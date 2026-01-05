package vm

import (
	"archive/zip"
	"fmt"
	"focal-vm/internal/bytecode/bcio"
	"focal-vm/internal/bytecode/spec"
	"focal-vm/public/runtimeapi"
	"io"
	"os"
	"path"
	"strings"
)

type ModuleCollection struct {
	focal_archives []string
	directories    []string
}

func NewModuleCollection() runtimeapi.ModuleCollection {
	return &ModuleCollection{focal_archives: []string{}, directories: []string{}}
}

func (mc *ModuleCollection) AddDirectories(dirs ...string) {
	mc.directories = append(mc.directories, dirs...)
}

func (mc *ModuleCollection) AddArchives(focal_archives ...string) {
	mc.focal_archives = append(mc.focal_archives, focal_archives...)
}

func (mc *ModuleCollection) findModuleInArchive(archive string, modulePath string) ([]byte, error) {
	zf, err := zip.OpenReader(archive)
	if err != nil {
		return nil, err
	}

	moduleFile, err := zf.Open(modulePath)
	if err != nil {
		return nil, err
	}
	moduleBytes, err := io.ReadAll(moduleFile)
	if err != nil {
		return nil, err
	}
	moduleFile.Close()
	zf.Close()

	return moduleBytes, nil
}

func (mc *ModuleCollection) findModuleInDirectory(directory string, modulePath string) ([]byte, error) {
	path := path.Join(directory, modulePath)
	moduleFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	moduleBytes, err := io.ReadAll(moduleFile)
	if err != nil {
		return nil, err
	}
	moduleFile.Close()
	return moduleBytes, nil
}

func (mc *ModuleCollection) SearchForModule(moduleName string) (*spec.BCModule, error) {
	modulePath := strings.ReplaceAll(moduleName, ".", "/") + ".fbc"
	for _, archiveFile := range mc.focal_archives {
		bytes, _ := mc.findModuleInArchive(archiveFile, modulePath)
		if bytes != nil {
			moduleReader := bcio.NewReader(bytes)
			module := moduleReader.ReadModule()
			return module, nil
		}
	}
	for _, directory := range mc.directories {
		bytes, _ := mc.findModuleInDirectory(directory, modulePath)
		if bytes != nil {
			moduleReader := bcio.NewReader(bytes)
			module := moduleReader.ReadModule()
			module.SetName(moduleName)
			return module, nil
		}
	}
	return nil, fmt.Errorf("Could not find module \"%s\" in the module-collection!", modulePath)
}

func (mc *ModuleCollection) SearchForModuleBytes(moduleName string) ([]byte, error) {
	modulePath := strings.ReplaceAll(moduleName, ".", "/") + ".fbc"
	for _, archiveFile := range mc.focal_archives {
		bytes, _ := mc.findModuleInArchive(archiveFile, modulePath)
		if bytes != nil {
			return bytes, nil
		}
	}
	for _, directory := range mc.directories {
		bytes, _ := mc.findModuleInDirectory(directory, modulePath)
		if bytes != nil {
			return bytes, nil
		}
	}
	return nil, fmt.Errorf("Could not find module \"%s\" in the module-collection!", modulePath)
}
