package constants

import (
	"fmt"
	"strconv"
)

type ConstantPool struct {
	constants []Constant
}

func NewConstantPool() *ConstantPool {
	return &ConstantPool{
		[]Constant{},
	}
}

func (c *ConstantPool) String() string {
	return fmt.Sprint(c.constants)
}

func (p *ConstantPool) Size() int32 {
	return int32(len(p.constants))
}

func (p *ConstantPool) AddConstant(c Constant) int32 {
	idx := len(p.constants)
	p.constants = append(p.constants, c)
	return int32(idx)
}

func (p *ConstantPool) ExpectConstant(idx int32, tag ConstantTag) Constant {
	constant := p.constants[idx]
	if constant.GetTag() != tag {
		panic("Expected constant type " + strconv.Itoa(int(tag)) + " got " + strconv.Itoa(int(constant.GetTag())))
	}
	return constant
}

func (p *ConstantPool) GetConstant(idx int32) Constant {
	return p.constants[idx]
}

func (p *ConstantPool) GetOrCreateBool(value bool) int32 {
	for i, v := range p.constants {
		v, ok := v.(*ConstantBoolean)
		if ok && v.value == value {
			return int32(i)
		}
	}

	return p.AddConstant(NewBooleanConstant(value))
}

func (p *ConstantPool) GetOrCreateI8(value int8) int32 {
	for i, v := range p.constants {
		v, ok := v.(*ConstantInt8)
		if ok && v.value == value {
			return int32(i)
		}
	}

	return p.AddConstant(NewI8Constant(value))
}

func (p *ConstantPool) GetOrCreateI16(value int16) int32 {
	for i, v := range p.constants {
		v, ok := v.(*ConstantInt16)
		if ok && v.value == value {
			return int32(i)
		}
	}

	return p.AddConstant(NewI16Constant(value))
}

func (p *ConstantPool) GetOrCreateI32(value int32) int32 {
	for i, v := range p.constants {
		v, ok := v.(*ConstantInt32)
		if ok && v.value == value {
			return int32(i)
		}
	}

	return p.AddConstant(NewI32Constant(value))
}

func (p *ConstantPool) GetOrCreateI64(value int64) int32 {
	for i, v := range p.constants {
		v, ok := v.(*ConstantInt64)
		if ok && v.value == value {
			return int32(i)
		}
	}

	return p.AddConstant(NewI64Constant(value))
}

func (p *ConstantPool) GetOrCreateF32(value float32) int32 {
	for i, v := range p.constants {
		v, ok := v.(*ConstantFloat32)
		if ok && v.value == value {
			return int32(i)
		}
	}

	return p.AddConstant(NewF32Constant(value))
}

func (p *ConstantPool) GetOrCreateF64(value float64) int32 {
	for i, v := range p.constants {
		v, ok := v.(*ConstantFloat64)
		if ok && v.value == value {
			return int32(i)
		}
	}

	return p.AddConstant(NewF64Constant(value))
}

func (p *ConstantPool) GetOrCreateUTF8(value string) int32 {
	for i, v := range p.constants {
		v, ok := v.(*ConstantUTF8String)
		if ok && v.value == value {
			return int32(i)
		}
	}

	return p.AddConstant(NewUTF8Constant(value))
}

func (p *ConstantPool) GetUTF8(value string) int32 {
	for i, v := range p.constants {
		v, ok := v.(*ConstantUTF8String)
		if ok && v.value == value {
			return int32(i)
		}
	}

	return -1
}
