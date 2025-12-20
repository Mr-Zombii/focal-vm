package spec

import "focal-lang/internal/bytecode/constants"

var BC_MAGIC []uint8 = []uint8([]byte("FOCALS"))

type BCModule struct {
	magic []uint8
	major uint8
	minor uint8

	name              string
	cpool             *constants.ConstantPool
	mainFunctionIndex int32
	funcs             []*BCFunction
}

func NewBCModule(major uint8, minor uint8, name string, mainFunctionIndex int32, pool *constants.ConstantPool) *BCModule {
	return &BCModule{
		major:             major,
		minor:             minor,
		name:              name,
		cpool:             pool,
		mainFunctionIndex: mainFunctionIndex,
		funcs:             []*BCFunction{},
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

func (bcm *BCModule) GetConstantPool() *constants.ConstantPool {
	return bcm.cpool
}

func (bcm *BCModule) GetMainFunctionIndex() int32 {
	return bcm.mainFunctionIndex
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
	BCFunctionModPub uint8 = 0b01
	BCFunctionModFFI uint8 = 0b10
)

type BCFunction struct {
	nameIdx  int32
	modifier uint8
	code     []uint8
}

func NewBCFunction(nameIdx int32, modifier uint8, code []uint8) *BCFunction {
	return &BCFunction{
		nameIdx:  nameIdx,
		modifier: modifier,
		code:     code,
	}
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
