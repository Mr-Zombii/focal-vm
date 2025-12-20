package constants

type ConstantTag uint8

const (
	ConstantTagUnknown ConstantTag = iota
	ConstantTagInt8
	ConstantTagInt16
	ConstantTagInt32
	ConstantTagInt64
	ConstantTagInt128 //unused

	ConstantTagFloat8  //unused
	ConstantTagFloat16 //unused
	ConstantTagFloat32
	ConstantTagFloat64
	ConstantTagFloat128 // unused

	ConstantTagUTF8String
)

type Constant interface {
	GetTag() ConstantTag
}

type ConstantInt8 struct {
	value int8
}

func NewI8Constant(v int8) Constant {
	return &ConstantInt8{value: v}
}

func (c *ConstantInt8) GetTag() ConstantTag {
	return ConstantTagInt8
}

func (c *ConstantInt8) GetValue() int8 {
	return c.value
}

type ConstantInt16 struct {
	value int16
}

func NewI16Constant(v int16) Constant {
	return &ConstantInt16{value: v}
}

func (c *ConstantInt16) GetTag() ConstantTag {
	return ConstantTagInt16
}

func (c *ConstantInt16) GetValue() int16 {
	return c.value
}

type ConstantInt32 struct {
	value int32
}

func NewI32Constant(v int32) Constant {
	return &ConstantInt32{value: v}
}

func (c *ConstantInt32) GetTag() ConstantTag {
	return ConstantTagInt16
}

func (c *ConstantInt32) GetValue() int32 {
	return c.value
}

type ConstantInt64 struct {
	value int64
}

func NewI64Constant(v int64) Constant {
	return &ConstantInt64{value: v}
}

func (c *ConstantInt64) GetTag() ConstantTag {
	return ConstantTagInt16
}

func (c *ConstantInt64) GetValue() int64 {
	return c.value
}

type ConstantFloat32 struct {
	value float32
}

func NewF32Constant(v float32) Constant {
	return &ConstantFloat32{value: v}
}

func (c *ConstantFloat32) GetTag() ConstantTag {
	return ConstantTagFloat32
}

func (c *ConstantFloat32) GetValue() float32 {
	return c.value
}

type ConstantFloat64 struct {
	value float64
}

func NewF64Constant(v float64) Constant {
	return &ConstantFloat64{value: v}
}

func (c *ConstantFloat64) GetTag() ConstantTag {
	return ConstantTagFloat64
}

func (c *ConstantFloat64) GetValue() float64 {
	return c.value
}

type ConstantUTF8String struct {
	length int32
	value  string
}

func NewUTF8Constant(v string) Constant {
	return &ConstantUTF8String{length: int32(len(v)), value: v}
}

func (c *ConstantUTF8String) GetTag() ConstantTag {
	return ConstantTagUTF8String
}

func (c *ConstantUTF8String) GetValue() string {
	return c.value
}

func (c *ConstantUTF8String) GetLength() int32 {
	return c.length
}
