package runtime

import (
	"fmt"
	"focal-vm/internal/bytecode/constants"
	"focal-vm/internal/bytecode/spec"
	"focal-vm/public/runtimeapi"
	"reflect"
	"runtime"
	"strconv"
)

type ScopeValue struct {
	scope runtimeapi.Scope
}

func (v *ScopeValue) DefineField(name string) {
	v.scope.DefineLocal(name)
}

func (v *ScopeValue) HasField(name string) bool {
	return v.scope.OwnsLocal(name)
}

func (v *ScopeValue) SetField(name string, value runtimeapi.Value) error {
	if !v.scope.OwnsLocal(name) {
		return fmt.Errorf("Tried to set field on object when field was not defined!")
	}
	return v.scope.SetLocal(name, value)
}

func (v *ScopeValue) GetField(name string) (runtimeapi.Value, error) {
	if !v.scope.OwnsLocal(name) {
		return nil, fmt.Errorf("Tried to get field on object when field was not defined!")
	}
	return v.scope.GetLocal(name)
}

func NewScopeValue(scope runtimeapi.Scope) *ScopeValue {
	return &ScopeValue{scope: scope}
}

func (v *ScopeValue) GetTag() runtimeapi.ValueTag {
	return runtimeapi.ValueTagScope
}

func (v *ScopeValue) GetScope() runtimeapi.Scope {
	return v.scope
}

func (v *ScopeValue) String() string {
	return "ScopeValue"
}

func (v *ScopeValue) GetRawValue() interface{} {
	return v.scope
}

type NativeFunctionValue struct {
	name      string
	function  func(runtimeapi.VM)
	reflected reflect.Value
}

func NewNativeFunction(fn func(runtimeapi.VM)) *NativeFunctionValue {
	reflection := reflect.ValueOf(fn)
	return &NativeFunctionValue{function: fn, reflected: reflection, name: runtime.FuncForPC(reflection.Pointer()).Name()}
}

func (v *NativeFunctionValue) Call(vm runtimeapi.VM) {
	vm.GetCallStack().Push(NewPseudoFrame(vm.GetCallStack().GetTop(), vm.GetScope(), "{ Native-Function }", runtime.FuncForPC(reflect.ValueOf(v.function).Pointer()).Name()))
	v.function(vm)
}

func (v *NativeFunctionValue) GetTag() runtimeapi.ValueTag {
	return runtimeapi.ValueTagNativeFunction
}

func (v *NativeFunctionValue) GetFunction() func(runtimeapi.VM) {
	return v.function
}

func (v *NativeFunctionValue) String() string {
	return "NativeFunction { " + v.name + " }"
}

func (v *NativeFunctionValue) GetName() string {
	return v.name
}

func (v *NativeFunctionValue) GetReflection() reflect.Value {
	return v.reflected
}

func (v *NativeFunctionValue) GetRawValue() interface{} {
	return v.function
}

type FunctionValue struct {
	parent   runtimeapi.Scope
	function *spec.BCFunction
}

func NewFunction(parent runtimeapi.Scope, function *spec.BCFunction) *FunctionValue {
	return &FunctionValue{parent: parent, function: function}
}

func (v *FunctionValue) Call(vm runtimeapi.VM) {
	parentScope := v.parent
	if parentScope == nil {
		parentScope = vm.GetScope()
	}
	frame := NewFrame(vm.GetCallStack().GetTop(), parentScope, v.function.GetModule(), v.function)
	vm.GetCallStack().Push(frame)
}

func (v *FunctionValue) GetTag() runtimeapi.ValueTag {
	return runtimeapi.ValueTagFunction
}

func (v *FunctionValue) GetFunction() *spec.BCFunction {
	return v.function
}

func (v *FunctionValue) String() string {
	return v.function.GetModule().GetName() + " -> " + v.function.GetModule().GetConstantPool().GetConstant(v.function.GetNameIndex()).(*constants.ConstantUTF8String).GetValue()
}

func (v *FunctionValue) GetRawValue() interface{} {
	return v.function
}

type ArrayValue struct {
	value []runtimeapi.Value
}

func NewArrayValue(value []runtimeapi.Value) runtimeapi.Value {
	return &ArrayValue{value: value}
}

func (v *ArrayValue) GetTag() runtimeapi.ValueTag {
	return runtimeapi.ValueTagArray
}

func (v *ArrayValue) GetValue() []runtimeapi.Value {
	return v.value
}

func (v *ArrayValue) GetLength() int32 {
	return int32(len(v.value))
}

func (v *ArrayValue) String() string {
	return fmt.Sprintf("%v", v.value)
}

func (v *ArrayValue) GetRawValue() interface{} {
	return v.value
}

var BOOLEAN_VALUE_TRUE = NewBooleanValue(true)
var BOOLEAN_VALUE_FALSE = NewBooleanValue(false)

type BooleanValue struct {
	value bool
}

func NewBooleanValue(value bool) runtimeapi.Value {
	return &BooleanValue{value: value}
}

func (v *BooleanValue) GetTag() runtimeapi.ValueTag {
	return runtimeapi.ValueTagBoolean
}

func (v *BooleanValue) GetValue() bool {
	return v.value
}

func (v *BooleanValue) String() string {
	return strconv.FormatBool(v.value)
}

func (v *BooleanValue) GetRawValue() interface{} {
	return v.value
}

type Int8Value struct {
	value int8
}

func NewInt8Value(value int8) runtimeapi.Value {
	return &Int8Value{value: value}
}

func (v *Int8Value) GetTag() runtimeapi.ValueTag {
	return runtimeapi.ValueTagInt8
}

func (v *Int8Value) GetValue() int8 {
	return v.value
}

func (v *Int8Value) String() string {
	return strconv.Itoa(int(v.value))
}

func (v *Int8Value) GetRawValue() interface{} {
	return v.value
}

type Int16Value struct {
	value int16
}

func NewInt16Value(value int16) runtimeapi.Value {
	return &Int16Value{value: value}
}

func (v *Int16Value) GetTag() runtimeapi.ValueTag {
	return runtimeapi.ValueTagInt16
}

func (v *Int16Value) GetValue() int16 {
	return v.value
}

func (v *Int16Value) String() string {
	return strconv.Itoa(int(v.value))
}

func (v *Int16Value) GetRawValue() interface{} {
	return v.value
}

type Int32Value struct {
	value int32
}

func NewInt32Value(value int32) runtimeapi.Value {
	return &Int32Value{value: value}
}

func (v *Int32Value) GetTag() runtimeapi.ValueTag {
	return runtimeapi.ValueTagInt32
}

func (v *Int32Value) GetValue() int32 {
	return v.value
}

func (v *Int32Value) String() string {
	return strconv.Itoa(int(v.value))
}

func (v *Int32Value) GetRawValue() interface{} {
	return v.value
}

type Int64Value struct {
	value int64
}

func NewInt64Value(value int64) runtimeapi.Value {
	return &Int64Value{value: value}
}

func (v *Int64Value) GetTag() runtimeapi.ValueTag {
	return runtimeapi.ValueTagInt64
}

func (v *Int64Value) GetValue() int64 {
	return v.value
}

func (v *Int64Value) String() string {
	return strconv.Itoa(int(v.value))
}

func (v *Int64Value) GetRawValue() interface{} {
	return v.value
}

type Float32Value struct {
	value float32
}

func NewFloat32Value(value float32) runtimeapi.Value {
	return &Float32Value{value: value}
}

func (v *Float32Value) GetTag() runtimeapi.ValueTag {
	return runtimeapi.ValueTagFloat32
}

func (v *Float32Value) GetValue() float32 {
	return v.value
}

func (v *Float32Value) String() string {
	return fmt.Sprintf("%f", v.value)
}

func (v *Float32Value) GetRawValue() interface{} {
	return v.value
}

type Float64Value struct {
	value float64
}

func NewFloat64Value(value float64) runtimeapi.Value {
	return &Float64Value{value: value}
}

func (v *Float64Value) GetTag() runtimeapi.ValueTag {
	return runtimeapi.ValueTagFloat64
}

func (v *Float64Value) GetValue() float64 {
	return v.value
}

func (v *Float64Value) String() string {
	return fmt.Sprintf("%f", v.value)
}

func (v *Float64Value) GetRawValue() interface{} {
	return v.value
}

type UTF8StringValue struct {
	value string
}

func NewUTF8StringValue(value string) runtimeapi.Value {
	return &UTF8StringValue{value: value}
}

func (v *UTF8StringValue) GetTag() runtimeapi.ValueTag {
	return runtimeapi.ValueTagUTF8String
}

func (v *UTF8StringValue) GetValue() string {
	return v.value
}

func (v *UTF8StringValue) String() string {
	return v.value
}

func (v *UTF8StringValue) GetRawValue() interface{} {
	return v.value
}

func ConstantToValue(c constants.Constant) runtimeapi.Value {
	switch cv := c.(type) {
	case *constants.ConstantBoolean:
		if cv.GetValue() {
			return BOOLEAN_VALUE_TRUE
		}
		return BOOLEAN_VALUE_FALSE
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

func ValueIsInteger(v runtimeapi.Value) bool {
	switch v.GetTag() {
	case runtimeapi.ValueTagInt8:
		return true
	case runtimeapi.ValueTagInt16:
		return true
	case runtimeapi.ValueTagInt32:
		return true
	case runtimeapi.ValueTagInt64:
		return true
	default:
		return false
	}
}

func ValueIsFloat(v runtimeapi.Value) bool {
	switch v.GetTag() {
	case runtimeapi.ValueTagFloat32:
		return true
	case runtimeapi.ValueTagFloat64:
		return true
	default:
		return false
	}
}

func GetValueAsInt(v runtimeapi.Value) int {
	switch v.GetTag() {
	case runtimeapi.ValueTagInt8:
		return int(v.(*Int8Value).value)
	case runtimeapi.ValueTagInt16:
		return int(v.(*Int16Value).value)
	case runtimeapi.ValueTagInt32:
		return int(v.(*Int32Value).value)
	case runtimeapi.ValueTagInt64:
		return int(v.(*Int64Value).value)
	default:
		return -1
	}
}

func GetValueAsFloat64(v runtimeapi.Value) float64 {
	switch v.GetTag() {
	case runtimeapi.ValueTagFloat32:
		return float64(v.(*Float32Value).value)
	case runtimeapi.ValueTagFloat64:
		return float64(v.(*Float64Value).value)
	default:
		return -1
	}
}

func GetValueAsFloat32(v runtimeapi.Value) float32 {
	switch v.GetTag() {
	case runtimeapi.ValueTagFloat32:
		return float32(v.(*Float32Value).value)
	default:
		return -1
	}
}
