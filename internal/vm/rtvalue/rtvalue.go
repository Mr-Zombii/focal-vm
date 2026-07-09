package rtvalue

import (
	"fmt"
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/bytecode/spec"
	"focal-vm/internal/erroring"
	"focal-vm/internal/vm/runtime/allocator"
	"unsafe"
)

type RTValueTag uint8

const (
	RTValueTag_UNKNOWN RTValueTag = iota
	RTValueTag_I8
	RTValueTag_I16
	RTValueTag_I32
	RTValueTag_I64
	//RTValueTag_U8
	//RTValueTag_U16
	//RTValueTag_U32
	//RTValueTag_U64
	RTValueTag_F32
	RTValueTag_F64

	RTValueTag_BOOL
	RTValueTag_STRING
	RTValueTag_ARRAY
	RTValueTag_VMFUNCTION
	RTValueTag_GOFUNCTION
	RTValueTag_STRUCT
	//RTValueTag_POINTER
)

var RTValueTag_I8_SIZE = int32(unsafe.Sizeof(RTValueI8{}))
var RTValueTag_I16_SIZE = int32(unsafe.Sizeof(RTValueI16{}))
var RTValueTag_I32_SIZE = int32(unsafe.Sizeof(RTValueI32{}))
var RTValueTag_I64_SIZE = int32(unsafe.Sizeof(RTValueI64{}))
var RTValueTag_F32_SIZE = int32(unsafe.Sizeof(RTValueF32{}))
var RTValueTag_F64_SIZE = int32(unsafe.Sizeof(RTValueF64{}))
var RTValueTag_BOOL_SIZE = int32(unsafe.Sizeof(RTValueBool{}))
var RTValueTag_STRING_SIZE = int32(unsafe.Sizeof(RTValueString{}))
var RTValueTag_ARRAY_SIZE = int32(unsafe.Sizeof(RTValueArray{}))
var RTValueTag_VMFUNCTION_SIZE = int32(unsafe.Sizeof(RTValueVMFunction{}))
var RTValueTag_GOFUNCTION_SIZE = int32(unsafe.Sizeof(RTValueGOFunction{}))
var RTValueTag_STRUCT_SIZE = int32(unsafe.Sizeof(RTValueStruct{}))

type RTValue interface {
	GetTag() RTValueTag
	GetType() bctypes.BCType
	String() string

	Mark()
	IsMarked() bool
	ResetMark()

	IncRefCount()
	DecRefCount()
	GetRefCount() int32

	GetPoolIdx() int32

	SetPool(pool *RTValuePool)
	SetIdx(idx int32)
	Walk(f func(v RTValue))
	GetHeader() *rtValueHeader

	OnFree()
}

type rtValueHeader struct {
	pool     *RTValuePool
	self     RTValue
	refCount int32
	poolIdx  int32
	mark     bool
	tag      RTValueTag
	typ      bctypes.BCType
}

func (v *rtValueHeader) postFree() {
	if v.mark {
		return
	}
	v.mark = true
	v.self.OnFree()
}

func (v *rtValueHeader) GetHeader() *rtValueHeader {
	return v
}

func (v *rtValueHeader) SetPool(pool *RTValuePool) {
	v.pool = pool
}

func (v *rtValueHeader) SetIdx(idx int32) {
	v.poolIdx = idx
}

func (v *rtValueHeader) GetTag() RTValueTag {
	return v.tag
}

func (v *rtValueHeader) GetType() bctypes.BCType {
	return v.typ
}

func (v *rtValueHeader) ResetMark() {
	v.mark = false
}

func (v *rtValueHeader) Mark() {
	v.mark = true
}

func (v *rtValueHeader) IsMarked() bool {
	return v.mark
}

func (v *rtValueHeader) IncRefCount() {
	v.refCount++
}

func (v *rtValueHeader) DecRefCount() {
	v.refCount--
	if v.refCount <= 0 {
		v.pool.RemoveAndFree(v.poolIdx)
	}
}

func (v *rtValueHeader) GetRefCount() int32 {
	return v.refCount
}

func (v *rtValueHeader) GetPoolIdx() int32 {
	return v.poolIdx
}

type RTValueI8 struct {
	rtValueHeader
	value int8
}

func (v *RTValueI8) GetValue() int8 {
	return v.value
}

func (v *RTValueI8) String() string {
	return fmt.Sprint(v.value)
}

func (v *RTValueI8) Walk(f func(v RTValue)) {}

func (v *RTValueI8) OnFree() {}

func newRTValueI8(allocator allocator.Allocator, value int8) *RTValueI8 {
	v := (*RTValueI8)(allocator.Alloc(RTValueTag_I8_SIZE))
	v.value = value
	v.tag = RTValueTag_I8
	v.typ = bctypes.I8
	v.self = v
	return v
}

type RTValueI16 struct {
	rtValueHeader
	value int16
}

func (v *RTValueI16) GetValue() int16 {
	return v.value
}

func (v *RTValueI16) String() string {
	return fmt.Sprint(v.value)
}

func (v *RTValueI16) Walk(f func(v RTValue)) {}

func (v *RTValueI16) OnFree() {}

func newRTValueI16(allocator allocator.Allocator, value int16) *RTValueI16 {
	v := (*RTValueI16)(allocator.Alloc(RTValueTag_I16_SIZE))
	v.value = value
	v.tag = RTValueTag_I16
	v.typ = bctypes.I16
	v.self = v
	return v
}

type RTValueI32 struct {
	rtValueHeader
	value int32
}

func (v *RTValueI32) GetValue() int32 {
	return v.value
}

func (v *RTValueI32) String() string {
	return fmt.Sprint(v.value)
}

func (v *RTValueI32) Walk(f func(v RTValue)) {}

func (v *RTValueI32) OnFree() {}

func newRTValueI32(allocator allocator.Allocator, value int32) *RTValueI32 {
	v := (*RTValueI32)(allocator.Alloc(RTValueTag_I32_SIZE))
	v.value = value
	v.tag = RTValueTag_I32
	v.typ = bctypes.I32
	v.self = v
	return v
}

type RTValueI64 struct {
	value int64
	rtValueHeader
}

func (v *RTValueI64) GetValue() int64 {
	return v.value
}

func (v *RTValueI64) String() string {
	return fmt.Sprint(v.value)
}

func (v *RTValueI64) Walk(f func(v RTValue)) {}

func (v *RTValueI64) OnFree() {}

func newRTValueI64(allocator allocator.Allocator, value int64) *RTValueI64 {
	v := (*RTValueI64)(allocator.Alloc(RTValueTag_I64_SIZE))
	v.value = value
	v.tag = RTValueTag_I64
	v.typ = bctypes.I64
	v.self = v
	return v
}

type RTValueF32 struct {
	rtValueHeader
	value float32
}

func (v *RTValueF32) GetValue() float32 {
	return v.value
}

func (v *RTValueF32) String() string {
	return fmt.Sprint(v.value)
}

func (v *RTValueF32) Walk(f func(v RTValue)) {}

func (v *RTValueF32) OnFree() {}

func newRTValueF32(allocator allocator.Allocator, value float32) *RTValueF32 {
	v := (*RTValueF32)(allocator.Alloc(RTValueTag_F32_SIZE))
	v.value = value
	v.tag = RTValueTag_F32
	v.typ = bctypes.F32
	v.self = v
	return v
}

type RTValueF64 struct {
	rtValueHeader
	value float64
}

func (v *RTValueF64) GetValue() float64 {
	return v.value
}

func (v *RTValueF64) String() string {
	return fmt.Sprint(v.value)
}

func (v *RTValueF64) Walk(f func(v RTValue)) {}

func (v *RTValueF64) OnFree() {}

func newRTValueF64(allocator allocator.Allocator, value float64) *RTValueF64 {
	v := (*RTValueF64)(allocator.Alloc(RTValueTag_F64_SIZE))
	v.value = value
	v.tag = RTValueTag_F64
	v.typ = bctypes.F64
	v.self = v
	return v
}

type RTValueBool struct {
	rtValueHeader
	value bool
}

func (v *RTValueBool) GetValue() bool {
	return v.value
}

func (v *RTValueBool) String() string {
	return fmt.Sprint(v.value)
}

func (v *RTValueBool) Walk(f func(v RTValue)) {}

func (v *RTValueBool) OnFree() {}

func newRTValueBool(allocator allocator.Allocator, value bool) *RTValueBool {
	v := (*RTValueBool)(allocator.Alloc(RTValueTag_BOOL_SIZE))
	v.value = value
	v.tag = RTValueTag_BOOL
	v.typ = bctypes.BOOL
	v.self = v
	return v
}

type RTValueString struct {
	rtValueHeader
	value *string
}

func (v *RTValueString) GetValue() string {
	return *v.value
}

func (v *RTValueString) GetLength() int32 {
	return int32(len(*v.value))
}

func (v *RTValueString) String() string {
	//return "\"" + *v.value + "\""
	return *v.value
}

func (v *RTValueString) Walk(f func(v RTValue)) {}

func (v *RTValueString) OnFree() {
	v.pool.allocator.Free(v.value)
}

func newRTValueString(rtAllocator allocator.Allocator, value string) *RTValueString {
	v := (*RTValueString)(rtAllocator.Alloc(RTValueTag_STRING_SIZE))
	bStrPtr := rtAllocator.Alloc(allocator.GetTotalStringSize(value))
	strPtr, _ := rtAllocator.CopyString(bStrPtr, value)

	v.value = strPtr
	v.tag = RTValueTag_STRING
	v.typ = bctypes.UTF_STRING
	v.self = v
	return v
}

type RTValueArray struct {
	rtValueHeader
	value *[]RTValue
}

func (v *RTValueArray) GetValue() []RTValue {
	return *v.value
}

func (v *RTValueArray) GetLength() int32 {
	return int32(len(*v.value))
}

func (v *RTValueArray) String() string {
	return fmt.Sprint(*v.value)
}

func (v *RTValueArray) Walk(f func(v RTValue)) {
	for _, val := range *v.value {
		if val != nil {
			f(val)
		}
	}
}

func (v *RTValueArray) OnFree() {
	values := *v.value
	length := int32(len(values))
	for i := range length {
		val := values[i]
		if val != nil {
			val.DecRefCount()
		}
	}
}

func newRTValueArray(allocator allocator.Allocator, elemType bctypes.BCType, length int32) *RTValueArray {
	backing := make([]RTValue, length)

	v := (*RTValueArray)(allocator.Alloc(RTValueTag_ARRAY_SIZE))

	tpool := elemType.GetTypePool()
	v.value = &backing
	v.tag = RTValueTag_ARRAY
	v.typ, _ = tpool.GetOrCreateArrayType(bctypes.NewArrayType(tpool, tpool.GetTypeIdx(elemType)))
	v.self = v
	return v
}

type RTValueStruct struct {
	rtValueHeader
	name   string
	fields []RTValue
}

func (v *RTValueStruct) GetFields() []RTValue {
	return v.fields
}

func (v *RTValueStruct) GetField(idx int32) RTValue {
	return v.fields[idx]
}

func (v *RTValueStruct) SetField(idx int32, value RTValue) {
	tpool := v.typ.GetTypePool()
	expectedFieldType := tpool.GetType(v.typ.(*bctypes.StructType).GetFieldTypeIndexes()[idx])

	if value.GetType().Equals(expectedFieldType) {
		v.fields[idx] = value
		return
	}

	erroring.GlobalErrorHandler.Panic(fmt.Sprintf(
		"Cannot set field at index '%d' to value '%s' when the expected type '%s' and the actual type '%s' don't match",
		idx, value, expectedFieldType, value.GetType(),
	))
}

func (v *RTValueStruct) Walk(f func(v RTValue)) {
	for _, val := range v.fields {
		if val != nil {
			f(val)
		}
	}
}

func (v *RTValueStruct) OnFree() {
	for _, field := range v.fields {
		field.DecRefCount()
	}
}

func (v *RTValueStruct) String() string {
	return fmt.Sprintf("{ name: %s, type: %s, fields: %s }", v.name, v.typ.String(), v.fields)
}

func newRTValueStruct(allocator allocator.Allocator, structType *bctypes.StructType) *RTValueStruct {
	length := int32(len(structType.GetFieldTypeIndexes()))
	backing := make([]RTValue, length)

	v := (*RTValueStruct)(allocator.Alloc(RTValueTag_STRUCT_SIZE))
	v.name = structType.GetName()
	v.fields = backing
	v.tag = RTValueTag_STRUCT
	v.typ = structType
	v.self = v
	return v
}

type RTValueVMFunction struct {
	fnPointer *spec.BCFunction
	rtValueHeader
}

func (v *RTValueVMFunction) GetFunction() *spec.BCFunction {
	return v.fnPointer
}

func (v *RTValueVMFunction) String() string {
	return "{ Module: \"" + v.fnPointer.GetModule().GetName() + "\", Name: \"" + v.fnPointer.GetName() + "\" }"
}

func (v *RTValueVMFunction) Walk(f func(v RTValue)) {}

func (v *RTValueVMFunction) OnFree() {}

func newRTValueVMFunction(allocator allocator.Allocator, fnPointer *spec.BCFunction) *RTValueVMFunction {
	v := (*RTValueVMFunction)(allocator.Alloc(RTValueTag_VMFUNCTION_SIZE))
	v.tag = RTValueTag_VMFUNCTION
	v.typ = fnPointer.GetType()
	v.fnPointer = fnPointer
	v.self = v
	return v
}

type RTValueGOFunction struct {
	fnPointer interface{}
	rtValueHeader
}

func (v *RTValueGOFunction) GetFunction() interface{} {
	return v.fnPointer
}

func (v *RTValueGOFunction) String() string {
	return "{ GOFunction }"
}

func (v *RTValueGOFunction) Walk(f func(v RTValue)) {}

func (v *RTValueGOFunction) OnFree() {}

func newRTValueGOFunction(allocator allocator.Allocator, fnType *bctypes.FunctionType, fnPointer interface{}) *RTValueGOFunction {
	v := (*RTValueGOFunction)(allocator.Alloc(RTValueTag_GOFUNCTION_SIZE))
	v.tag = RTValueTag_GOFUNCTION
	v.typ = fnType
	v.fnPointer = fnPointer
	v.self = v
	return v
}

func ValueIsInteger(v RTValue) bool {
	switch v.GetTag() {
	case RTValueTag_I8:
		return true
	case RTValueTag_I16:
		return true
	case RTValueTag_I32:
		return true
	case RTValueTag_I64:
		return true
	default:
		return false
	}
}

func ValueIsFloat(v RTValue) bool {
	switch v.GetTag() {
	case RTValueTag_F32:
		return true
	case RTValueTag_F64:
		return true
	default:
		return false
	}
}

func GetValueAsInt(v RTValue) int {
	switch v.GetTag() {
	case RTValueTag_I8:
		return int(v.(*RTValueI8).value)
	case RTValueTag_I16:
		return int(v.(*RTValueI16).value)
	case RTValueTag_I32:
		return int(v.(*RTValueI32).value)
	case RTValueTag_I64:
		return int(v.(*RTValueI64).value)
	default:
		return -1
	}
}

func GetValueAsFloat64(v RTValue) float64 {
	switch v.GetTag() {
	case RTValueTag_F32:
		return float64(v.(*RTValueF32).value)
	case RTValueTag_F64:
		return v.(*RTValueF64).value
	default:
		return -1
	}
}

func GetValueAsFloat32(v RTValue) float32 {
	switch v.GetTag() {
	case RTValueTag_F32:
		return v.(*RTValueF32).value
	default:
		return -1
	}
}
