package spec

import (
	"fmt"
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/bytecode/constants"
	"strings"
)

var BC_MAGIC = []byte("FOCALS")

type BCModule struct {
	magic []uint8
	major uint8
	minor uint8

	name  string
	tpool *bctypes.TypePool
	cpool *constants.ConstantPool
	funcs []*BCFunction
}

func NewBCModule(major uint8, minor uint8, name string, tpool *bctypes.TypePool, cpool *constants.ConstantPool) *BCModule {
	return &BCModule{
		major: major,
		minor: minor,
		name:  name,
		cpool: cpool,
		tpool: tpool,
		funcs: []*BCFunction{},
	}
}

func (bcm *BCModule) GetBytecodeMajorVersion() uint8 {
	return bcm.major
}

func (bcm *BCModule) GetBytecodeMinorVersion() uint8 {
	return bcm.minor
}

func (bcm *BCModule) GetName() string {
	return bcm.name
}

func (bcm *BCModule) GetFileName() string {
	return strings.ReplaceAll(bcm.name, ".", "/") + ".fbc"
}

func (bcm *BCModule) GetConstantPool() *constants.ConstantPool {
	return bcm.cpool
}

func (bcm *BCModule) GetTypePool() *bctypes.TypePool {
	return bcm.tpool
}

func (bcm *BCModule) GetFunctions() []*BCFunction {
	return bcm.funcs
}

func (bcm *BCModule) SetFunctions(funcs []*BCFunction) {
	bcm.funcs = funcs
}

func (bcm *BCModule) GetFunction(name string, fnType *bctypes.FunctionType) (*BCFunction, error) {
	nameIdx := bcm.cpool.GetUTF8(name)
	_, typeIdx := bcm.tpool.GetFunctionType(fnType)
	if nameIdx == -1 {
		return nil, fmt.Errorf("function with name '%s' not found in module '%s'", name, bcm.name)
	}

	if typeIdx == -1 {
		return nil, fmt.Errorf("function with name '%s' type signature '%s' not found in module '%s'", name, fnType.String(), bcm.name)
	}

	for _, v := range bcm.funcs {
		if v.nameIdx == nameIdx && v.typeIdx == typeIdx {
			return v, nil
		}
	}
	return nil, fmt.Errorf("function with name '%s' type signature '%s' not found in module '%s'", name, fnType.String(), bcm.name)
}

const (
	BCFunctionModPrivate uint8 = 0b01
	BCFunctionModSubFunc uint8 = 0b10
)

type BCFunction struct {
	module           *BCModule
	nameIdx          int32
	typeIdx          int32
	paramNameIndexes []int32
	modifier         uint8
	code             []uint8
}

func NewBCFunction(module *BCModule, modifier uint8, nameIdx int32, typeIdx int32, paramNameIndexes []int32, code []uint8) *BCFunction {
	module.cpool.ExpectConstant(nameIdx, constants.ConstantTagUTF8String)
	fnType := module.tpool.ExpectType(typeIdx, bctypes.BCTYPE_FUNCTION)

	if len(fnType.(*bctypes.FunctionType).GetParamTypeIndexes()) != len(paramNameIndexes) {
		panic(fmt.Sprintf("Mismatch between param type count and param name count"))
	}

	existingIndexes := map[int32]bool{}
	for i := range paramNameIndexes {
		if _, ok := existingIndexes[paramNameIndexes[i]]; ok {
			panic(fmt.Sprintf("Duplicate param name %s", module.cpool.ExpectConstant(paramNameIndexes[i], constants.ConstantTagUTF8String).(*constants.ConstantUTF8String).GetValue()))
		}
		existingIndexes[paramNameIndexes[i]] = true
	}

	return &BCFunction{
		module:           module,
		nameIdx:          nameIdx,
		typeIdx:          typeIdx,
		paramNameIndexes: paramNameIndexes,
		modifier:         modifier,
		code:             code,
	}
}

func (bcf *BCFunction) GetName() string {
	return bcf.module.GetConstantPool().ExpectConstant(bcf.nameIdx, constants.ConstantTagUTF8String).(*constants.ConstantUTF8String).GetValue()
}

func (bcf *BCFunction) GetModule() *BCModule {
	return bcf.module
}

func (bcf *BCFunction) GetNameIndex() int32 {
	return bcf.nameIdx
}

func (bcf *BCFunction) GetModifier() uint8 {
	return bcf.modifier
}

func (bcf *BCFunction) GetCode() []uint8 {
	return bcf.code
}

func (bcf *BCFunction) GetTypeIdx() int32 {
	return bcf.typeIdx
}

func (bcf *BCFunction) GetParamNameIndexes() []int32 {
	return bcf.paramNameIndexes
}

func (bcf *BCFunction) GetType() *bctypes.FunctionType {
	return bcf.module.GetTypePool().ExpectType(bcf.typeIdx, bctypes.BCTYPE_FUNCTION).(*bctypes.FunctionType)
}
