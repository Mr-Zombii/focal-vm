package main

import (
	"archive/zip"
	"focal-vm/internal/bytecode/bcio"
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/bytecode/constants"
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/bytecode/spec"
	"os"
)

var modules []*spec.BCModule

func newModule(moduleName string, action func(fnCreator func(modifier uint8, fnName string, fnType *bctypes.FunctionType, paramNames []string, code []uint8), module *spec.BCModule, tpool *bctypes.TypePool, cpool *constants.ConstantPool)) *spec.BCModule {
	tpool := bctypes.NewTypePool()
	cpool := constants.NewConstantPool()
	module := spec.NewBCModule(1, 0, moduleName, tpool, cpool)
	var functions []*spec.BCFunction
	fnMaker := func(modifier uint8, fnName string, fnType *bctypes.FunctionType, paramNames []string, code []uint8) {
		paramNameIndexes := make([]int32, len(paramNames))
		for i, v := range paramNames {
			paramNameIndexes[i] = cpool.GetOrCreateUTF8(v)
		}
		functions = append(functions, spec.NewBCFunction(module, modifier, cpool.GetOrCreateUTF8(fnName), tpool.AddType(fnType), paramNameIndexes, code))
	}
	action(fnMaker, module, tpool, cpool)
	module.SetFunctions(functions)
	modules = append(modules, module)
	return module
}

func main() {
	newModule("fibb", func(
		fnCreator func(modifier uint8, fnName string, fnType *bctypes.FunctionType, paramNames []string, code []uint8),
		module *spec.BCModule, tpool *bctypes.TypePool, cpool *constants.ConstantPool,
	) {
		_, i64TypeIdx := tpool.GetOrCreateI64Type()
		fnType := bctypes.NewFunctionType(tpool, []int32{i64TypeIdx}, i64TypeIdx)
		_, fnTypeIdx := tpool.GetOrCreateFunctionType(fnType)
		fnCreator(0, "fibonacci", fnType, []string{"n"}, []uint8{
			uint8(opcodes.OP_LOAD_CONST),
			uint8(0),
			uint8(cpool.GetOrCreateI64(1)),

			uint8(opcodes.OP_GET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("n")),

			uint8(opcodes.OP_ILE),

			uint8(opcodes.OP_BRANCH),
			uint8(0),
			uint8(4),

			uint8(opcodes.OP_GET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("n")),

			uint8(opcodes.OP_RET),

			uint8(opcodes.OP_LOAD_CONST),
			uint8(0),
			uint8(cpool.GetOrCreateI64(1)),

			uint8(opcodes.OP_GET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("n")),

			uint8(opcodes.OP_ISUB),

			uint8(opcodes.OP_LOAD_STATIC_FUNCTION),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("fibb")),
			uint8(cpool.GetOrCreateUTF8("fibonacci")),
			uint8(fnTypeIdx),
			uint8(opcodes.OP_CALL),

			uint8(opcodes.OP_LOAD_CONST),
			uint8(0),
			uint8(cpool.GetOrCreateI64(2)),

			uint8(opcodes.OP_GET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("n")),

			uint8(opcodes.OP_ISUB),

			uint8(opcodes.OP_LOAD_STATIC_FUNCTION),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("fibb")),
			uint8(cpool.GetOrCreateUTF8("fibonacci")),
			uint8(fnTypeIdx),
			uint8(opcodes.OP_CALL),

			uint8(opcodes.OP_IADD),

			uint8(opcodes.OP_RET),
		})

		fnType = bctypes.NewFunctionType(tpool, []int32{i64TypeIdx, i64TypeIdx, i64TypeIdx}, i64TypeIdx)
		_, fnTypeIdx = tpool.GetOrCreateFunctionType(fnType)
		fnCreator(0, "fibonacci", fnType, []string{"n", "a", "b"}, []uint8{
			uint8(opcodes.OP_LOAD_CONST),
			uint8(0),
			uint8(cpool.GetOrCreateI64(0)),

			uint8(opcodes.OP_GET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("n")),

			uint8(opcodes.OP_EQ),

			uint8(opcodes.OP_BRANCH),
			uint8(0),
			uint8(4),

			uint8(opcodes.OP_GET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("a")),

			uint8(opcodes.OP_RET),

			uint8(opcodes.OP_LOAD_CONST),
			uint8(0),
			uint8(cpool.GetOrCreateI64(1)),

			uint8(opcodes.OP_GET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("n")),

			uint8(opcodes.OP_ISUB),

			uint8(opcodes.OP_GET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("b")),
			uint8(opcodes.OP_DUP),

			uint8(opcodes.OP_GET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("a")),

			uint8(opcodes.OP_IADD),

			uint8(opcodes.OP_LOAD_STATIC_FUNCTION),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("fibb")),
			uint8(cpool.GetOrCreateUTF8("fibonacci")),
			uint8(fnTypeIdx),
			uint8(opcodes.OP_TCALL),

			uint8(opcodes.OP_RET),
		})

	})

	newModule("test", func(
		fnCreator func(modifier uint8, fnName string, fnType *bctypes.FunctionType, paramNames []string, code []uint8),
		module *spec.BCModule, tpool *bctypes.TypePool, cpool *constants.ConstantPool,
	) {
		_, i64TypeIdx := tpool.GetOrCreateI64Type()
		fnCreator(0, "loadMe", bctypes.NewFunctionType(tpool, []int32{}), []string{}, []uint8{
			uint8(opcodes.OP_LOAD_CONST),
			uint8(0),
			uint8(cpool.GetOrCreateI64(15)),

			uint8(opcodes.OP_LOAD_CONST),
			uint8(0),
			uint8(cpool.GetOrCreateI64(0)),

			uint8(opcodes.OP_LOAD_CONST),
			uint8(0),
			uint8(cpool.GetOrCreateI64(1)),

			uint8(opcodes.OP_LOAD_STATIC_FUNCTION),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("fibb")),
			uint8(cpool.GetOrCreateUTF8("fibonacci")),
			uint8(tpool.AddType(bctypes.NewFunctionType(tpool, []int32{i64TypeIdx, i64TypeIdx, i64TypeIdx}, i64TypeIdx))),
			uint8(opcodes.OP_CALL),

			uint8(opcodes.OP_RET),
		})
	})

	mainModule := newModule("testing.boostrap", func(
		fnCreator func(modifier uint8, fnName string, fnType *bctypes.FunctionType, paramNames []string, code []uint8),
		module *spec.BCModule, tpool *bctypes.TypePool, cpool *constants.ConstantPool,
	) {
		_, i := tpool.GetOrCreateUTFStringType()
		strArryIdx := tpool.AddType(bctypes.NewArrayType(tpool, i))
		mainFnType := bctypes.NewFunctionType(tpool, []int32{strArryIdx})
		tpool.AddType(mainFnType)

		structTypeIdx := tpool.AddType(bctypes.NewStructType(tpool, "thing", strArryIdx))

		fnCreator(0, "main", mainFnType, []string{"argv"}, []uint8{
			uint8(opcodes.OP_STRUCT_NEW),
			uint8(0),
			uint8(structTypeIdx),

			uint8(opcodes.OP_DUP),

			uint8(opcodes.OP_DEFINE_GLOBAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("MyStruct")),
			uint8(structTypeIdx),

			uint8(opcodes.OP_SET_GLOBAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("MyStruct")),

			uint8(opcodes.OP_GET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("argv")),

			uint8(opcodes.OP_LOAD_CONST),
			uint8(0),
			uint8(cpool.GetOrCreateI32(0)),

			uint8(opcodes.OP_STRUCT_SET_FIELD),

			uint8(opcodes.OP_GET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("MyStruct")),

			uint8(opcodes.OP_LOAD_STATIC_FUNCTION),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("test")),
			uint8(cpool.GetOrCreateUTF8("loadMe")),
			uint8(tpool.AddType(bctypes.NewFunctionType(tpool, []int32{}))),
			uint8(opcodes.OP_CALL),
			uint8(opcodes.OP_SWP),
			uint8(opcodes.OP_POP),

			uint8(opcodes.OP_RET),
		})
	})

	os.Remove("focal-archive.zip")
	archive, _ := os.Create("focal-archive.zip")
	focalArchive := zip.NewWriter(archive)

	for _, m := range modules {
		writer := bcio.NewWriter(m)
		writer.WriteModule()
		out := writer.GetBytes()

		f, _ := focalArchive.Create(m.GetFileName())
		f.Write(out)
	}

	f, _ := focalArchive.Create("focal-manifest.json")
	f.Write([]byte("{ \"main-module\": \"" + mainModule.GetName() + "\" }"))

	focalArchive.Close()
	archive.Close()
}
