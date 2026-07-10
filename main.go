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

			uint8(opcodes.OP_LE),

			uint8(opcodes.OP_JUMP_IF_FALSE),
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

			uint8(opcodes.OP_SUB),

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

			uint8(opcodes.OP_SUB),

			uint8(opcodes.OP_LOAD_STATIC_FUNCTION),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("fibb")),
			uint8(cpool.GetOrCreateUTF8("fibonacci")),
			uint8(fnTypeIdx),
			uint8(opcodes.OP_CALL),

			uint8(opcodes.OP_ADD),

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

			uint8(opcodes.OP_JUMP_IF_FALSE),
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

			uint8(opcodes.OP_SUB),

			uint8(opcodes.OP_GET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("b")),
			uint8(opcodes.OP_DUP),

			uint8(opcodes.OP_GET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("a")),

			uint8(opcodes.OP_ADD),

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
			uint8(cpool.GetOrCreateI64(92)),

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
	/*
		package testing;
		// boostrap.fbc

		void main(string[] argv) {
		  uint64 programPtr;
		  programPtr = _builtin_load_foreign("D:\project-archive\golang\c-test\mycode");

		  int64[] test;
		  test = new int64[2];

		  test[0] = 1;
		  test[1] = 2;

		  _builtin_print((int32)(_builtin_call_sym_foreign(programPtr, "add", array)));
		  _builtin_print("\n");

		  _builtin_free_foreign(programPtr);
		  return;
		}
	*/
	mainModule := newModule("testing.boostrap", func(
		fnCreator func(modifier uint8, fnName string, fnType *bctypes.FunctionType, paramNames []string, code []uint8),
		module *spec.BCModule, tpool *bctypes.TypePool, cpool *constants.ConstantPool,
	) {
		_, i := tpool.GetOrCreateUTFStringType()
		strArryIdx := tpool.AddType(bctypes.NewArrayType(tpool, i))
		mainFnType := bctypes.NewFunctionType(tpool, []int32{strArryIdx})
		tpool.AddType(mainFnType)

		//structTypeIdx := tpool.AddType(bctypes.NewStructType(tpool, "thing", strArryIdx))
		_, u64Ptr := tpool.GetOrCreateU64Type()

		t64, u64 := tpool.GetOrCreateI64Type()
		_, au64 := tpool.GetOrCreateArrayTypeFromElemType(t64)

		fnCreator(0, "main", mainFnType, []string{"argv"}, []uint8{
			// var programPtr
			uint8(opcodes.OP_DEFINE_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("programPtr")),
			uint8(u64Ptr),

			// _builtin_load_foreign("D:\project-archive\golang\c-test\mycode")
			uint8(opcodes.OP_LOAD_CONST),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("D:\\project-archive\\golang\\c-test\\mycode")),
			uint8(opcodes.OP_GET_GLOBAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("_builtin_load_foreign")),
			uint8(opcodes.OP_CALL),

			// programPtr = _builtin_load_foreign("D:\project-archive\golang\c-test\mycode")
			uint8(opcodes.OP_SET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("programPtr")),

			// var test
			uint8(opcodes.OP_DEFINE_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("test")),
			uint8(au64),

			// new array[2]
			uint8(opcodes.OP_LOAD_CONST),
			uint8(0),
			uint8(cpool.GetOrCreateI32(2)),
			uint8(opcodes.OP_ARRAY_NEW),
			uint8(0),
			uint8(u64),

			// test = array
			uint8(opcodes.OP_SET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("test")),

			// array[0] = 1
			uint8(opcodes.OP_LOAD_CONST),
			uint8(0),
			uint8(cpool.GetOrCreateI64(1)),

			uint8(opcodes.OP_LOAD_CONST),
			uint8(0),
			uint8(cpool.GetOrCreateI32(0)),

			uint8(opcodes.OP_GET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("test")),
			uint8(opcodes.OP_ARRAY_STORE),

			// array[1] = 2
			uint8(opcodes.OP_LOAD_CONST),
			uint8(0),
			uint8(cpool.GetOrCreateI64(2)),

			uint8(opcodes.OP_LOAD_CONST),
			uint8(0),
			uint8(cpool.GetOrCreateI32(1)),

			uint8(opcodes.OP_GET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("test")),
			uint8(opcodes.OP_ARRAY_STORE),

			// load arguments
			uint8(opcodes.OP_GET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("test")),

			// load procName
			uint8(opcodes.OP_LOAD_CONST),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("add")),

			uint8(opcodes.OP_GET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("programPtr")),

			// _builtin_call_sym_foreign(programPtr, "add", array)
			uint8(opcodes.OP_GET_GLOBAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("_builtin_call_sym_foreign")),
			uint8(opcodes.OP_CALL),

			uint8(opcodes.OP_CONV_TO_I32),

			uint8(opcodes.OP_GET_GLOBAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("_builtin_print")),
			uint8(opcodes.OP_CALL),

			uint8(opcodes.OP_LOAD_CONST),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("\n")),

			uint8(opcodes.OP_GET_GLOBAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("_builtin_print")),
			uint8(opcodes.OP_CALL),

			uint8(opcodes.OP_GET_LOCAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("programPtr")),

			// _builtin_free_foreign(programPtr)
			uint8(opcodes.OP_GET_GLOBAL),
			uint8(0),
			uint8(cpool.GetOrCreateUTF8("_builtin_free_foreign")),
			uint8(opcodes.OP_CALL),

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
