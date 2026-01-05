package spec

import (
	"focal-vm/internal/bytecode/constants"
	"strings"
)

var BC_MAGIC = []byte("FOCALS")

type BCModule struct {
	magic []uint8
	major uint8
	minor uint8

	name  string
	cpool *constants.ConstantPool
	funcs []*BCFunction
}

func NewBCModule(major uint8, minor uint8, name string, pool *constants.ConstantPool) *BCModule {
	return &BCModule{
		major: major,
		minor: minor,
		name:  name,
		cpool: pool,
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

func (bcm *BCModule) SetName(name string) {
	bcm.name = name
}

func (bcm *BCModule) GetConstantPool() *constants.ConstantPool {
	return bcm.cpool
}

func (bcm *BCModule) GetFunctions() []*BCFunction {
	return bcm.funcs
}

func (bcm *BCModule) SetFunctions(funcs []*BCFunction) {
	bcm.funcs = funcs
}

func (bcm *BCModule) GetFunction(s string) *BCFunction {
	nameIdx := bcm.cpool.GetUTF8(s)
	if nameIdx == -1 {
		return nil
	}

	for _, v := range bcm.funcs {
		if v.nameIdx == nameIdx {
			return v
		}
	}
	return nil
}

const (
	BCFunctionModPrivate uint8 = 0b01
	BCFunctionModSubFunc uint8 = 0b10
)

type BCFunction struct {
	module   *BCModule
	nameIdx  int32
	modifier uint8
	code     []uint8
}

func NewBCFunction(module *BCModule, nameIdx int32, modifier uint8, code []uint8) *BCFunction {
	return &BCFunction{
		module:   module,
		nameIdx:  nameIdx,
		modifier: modifier,
		code:     code,
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
