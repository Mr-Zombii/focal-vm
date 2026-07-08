package constants

import (
	"fmt"
	"focal-vm/internal/bytecode/bctypes"
)

type ConstantTag uint8

const (
	ConstantTagUnknown ConstantTag = iota

	ConstantTagBoolean

	ConstantTagInt8
	ConstantTagInt16
	ConstantTagInt32
	ConstantTagInt64

	ConstantTagFloat32
	ConstantTagFloat64

	ConstantTagUTF8String
)

type constantHeader struct {
	tag ConstantTag
	typ bctypes.BCType
}

func (v constantHeader) GetTag() ConstantTag {
	return v.tag
}

func (v constantHeader) GetType() bctypes.BCType {
	return v.typ
}

type Constant interface {
	GetTag() ConstantTag
	GetType() bctypes.BCType
}

type ConstantBoolean struct {
	constantHeader
	value bool
}

func NewBooleanConstant(v bool) Constant {
	c := &ConstantBoolean{value: v}
	c.tag = ConstantTagBoolean
	c.typ = bctypes.BOOL
	return c
}

func (c *ConstantBoolean) String() string {
	return fmt.Sprint(c.value)
}

func (c *ConstantBoolean) GetValue() bool {
	return c.value
}

type ConstantInt8 struct {
	constantHeader
	value int8
}

func NewI8Constant(v int8) Constant {
	c := &ConstantInt8{value: v}
	c.tag = ConstantTagInt8
	c.typ = bctypes.I8
	return c
}

func (c *ConstantInt8) String() string {
	return fmt.Sprint(c.value)
}

func (c *ConstantInt8) GetValue() int8 {
	return c.value
}

type ConstantInt16 struct {
	constantHeader
	value int16
}

func NewI16Constant(v int16) Constant {
	c := &ConstantInt16{value: v}
	c.tag = ConstantTagInt16
	c.typ = bctypes.I16
	return c
}

func (c *ConstantInt16) String() string {
	return fmt.Sprint(c.value)
}

func (c *ConstantInt16) GetValue() int16 {
	return c.value
}

type ConstantInt32 struct {
	constantHeader
	value int32
}

func NewI32Constant(v int32) Constant {
	c := &ConstantInt32{value: v}
	c.tag = ConstantTagInt32
	c.typ = bctypes.I32
	return c
}

func (c *ConstantInt32) String() string {
	return fmt.Sprint(c.value)
}

func (c *ConstantInt32) GetValue() int32 {
	return c.value
}

type ConstantInt64 struct {
	constantHeader
	value int64
}

func NewI64Constant(v int64) Constant {
	c := &ConstantInt64{value: v}
	c.tag = ConstantTagInt64
	c.typ = bctypes.I64
	return c
}

func (c *ConstantInt64) String() string {
	return fmt.Sprint(c.value)
}

func (c *ConstantInt64) GetValue() int64 {
	return c.value
}

type ConstantFloat32 struct {
	constantHeader
	value float32
}

func NewF32Constant(v float32) Constant {
	c := &ConstantFloat32{value: v}
	c.tag = ConstantTagFloat32
	c.typ = bctypes.F32
	return c
}

func (c *ConstantFloat32) String() string {
	return fmt.Sprint(c.value)
}

func (c *ConstantFloat32) GetValue() float32 {
	return c.value
}

type ConstantFloat64 struct {
	constantHeader
	value float64
}

func NewF64Constant(v float64) Constant {
	c := &ConstantFloat64{value: v}
	c.tag = ConstantTagFloat64
	c.typ = bctypes.F64
	return c
}

func (c *ConstantFloat64) String() string {
	return fmt.Sprint(c.value)
}

func (c *ConstantFloat64) GetValue() float64 {
	return c.value
}

type ConstantUTF8String struct {
	constantHeader
	length int32
	value  string
}

func NewUTF8Constant(v string) Constant {
	c := &ConstantUTF8String{length: int32(len(v)), value: v}
	c.tag = ConstantTagUTF8String
	c.typ = bctypes.UTF_STRING
	return c
}

func (c *ConstantUTF8String) String() string {
	return c.value
}

func (c *ConstantUTF8String) GetValue() string {
	return c.value
}

func (c *ConstantUTF8String) GetLength() int32 {
	return c.length
}
