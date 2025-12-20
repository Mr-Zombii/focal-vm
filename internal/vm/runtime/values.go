package runtime

import (
	"fmt"
	"focal-lang/internal/bytecode/constants"
	"focal-lang/internal/vm/api"
	"strconv"
)

const (
	ValueTagUnknown api.ValueTag = iota

	ValueTagInt8
	ValueTagInt16
	ValueTagInt32
	ValueTagInt64
	ValueTagInt128 // Unused

	ValueTagFloat8  // Unused
	ValueTagFloat16 // Unused
	ValueTagFloat32
	ValueTagFloat64
	ValueTagFloat128 // Unused

	ValueTagUTF8String

	ValueTagArray
)

type ArrayValue struct {
	value []api.Value
}

func NewArrayValue(value []api.Value) api.Value {
	return &ArrayValue{value: value}
}

func (v *ArrayValue) GetTag() api.ValueTag {
	return ValueTagArray
}

func (v *ArrayValue) GetValue() []api.Value {
	return v.value
}

func (v *ArrayValue) String() string {
	return fmt.Sprintf("%s", v.value)
}

type Int8Value struct {
	value int8
}

func NewInt8Value(value int8) api.Value {
	return &Int8Value{value: value}
}

func (v *Int8Value) GetTag() api.ValueTag {
	return ValueTagInt8
}

func (v *Int8Value) GetValue() int8 {
	return v.value
}

func (v *Int8Value) String() string {
	return strconv.Itoa(int(v.value))
}

type Int16Value struct {
	value int16
}

func NewInt16Value(value int16) api.Value {
	return &Int16Value{value: value}
}

func (v *Int16Value) GetTag() api.ValueTag {
	return ValueTagInt16
}

func (v *Int16Value) GetValue() int16 {
	return v.value
}

func (v *Int16Value) String() string {
	return strconv.Itoa(int(v.value))
}

type Int32Value struct {
	value int32
}

func NewInt32Value(value int32) api.Value {
	return &Int32Value{value: value}
}

func (v *Int32Value) GetTag() api.ValueTag {
	return ValueTagInt32
}

func (v *Int32Value) GetValue() int32 {
	return v.value
}

func (v *Int32Value) String() string {
	return strconv.Itoa(int(v.value))
}

type Int64Value struct {
	value int64
}

func NewInt64Value(value int64) api.Value {
	return &Int64Value{value: value}
}

func (v *Int64Value) GetTag() api.ValueTag {
	return ValueTagInt64
}

func (v *Int64Value) GetValue() int64 {
	return v.value
}

func (v *Int64Value) String() string {
	return strconv.Itoa(int(v.value))
}

type Float32Value struct {
	value float32
}

func NewFloat32Value(value float32) api.Value {
	return &Float32Value{value: value}
}

func (v *Float32Value) GetTag() api.ValueTag {
	return ValueTagFloat32
}

func (v *Float32Value) GetValue() float32 {
	return v.value
}

func (v *Float32Value) String() string {
	return fmt.Sprintf("%f", v.value)
}

type Float64Value struct {
	value float64
}

func NewFloat64Value(value float64) api.Value {
	return &Float64Value{value: value}
}

func (v *Float64Value) GetTag() api.ValueTag {
	return ValueTagFloat64
}

func (v *Float64Value) GetValue() float64 {
	return v.value
}

func (v *Float64Value) String() string {
	return fmt.Sprintf("%f", v.value)
}

type UTF8StringValue struct {
	value string
}

func NewUTF8StringValue(value string) api.Value {
	return &UTF8StringValue{value: value}
}

func (v *UTF8StringValue) GetTag() api.ValueTag {
	return ValueTagUTF8String
}

func (v *UTF8StringValue) GetValue() string {
	return v.value
}

func (v *UTF8StringValue) String() string {
	return v.value
}

func ConstantToValue(c constants.Constant) api.Value {
	switch cv := c.(type) {
	case *constants.ConstantInt8:
		return NewInt8Value(cv.GetValue())
	case *constants.ConstantInt16:
		return NewInt16Value(cv.GetValue())
	case *constants.ConstantInt32:
		return NewInt32Value(cv.GetValue())
	case *constants.ConstantInt64:
		return NewInt64Value(cv.GetValue())
	case *constants.ConstantFloat32:
		return NewFloat32Value(cv.GetValue())
	case *constants.ConstantFloat64:
		return NewFloat64Value(cv.GetValue())
	case *constants.ConstantUTF8String:
		return NewUTF8StringValue(cv.GetValue())
	}
	return nil
}
