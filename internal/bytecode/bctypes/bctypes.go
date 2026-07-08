package bctypes

import (
	"focal-vm/internal/util"
	"strconv"
)

type BCTypeTag uint8

const (
	BCTYPE_I8 BCTypeTag = iota
	BCTYPE_I16
	BCTYPE_I32
	BCTYPE_I64

	BCTYPE_U8
	BCTYPE_U16
	BCTYPE_U32
	BCTYPE_U64

	BCTYPE_F32
	BCTYPE_F64

	BCTYPE_BOOLEAN
	BCTYPE_UTFSTRING

	BCTYPE_ARRAY

	BCTYPE_STRUCT   // Struct[a,b,c,d]
	BCTYPE_MODULE   // Module[ModuleName]
	BCTYPE_FUNCTION // Function[a,b,c]d
	BCTYPE_POINTER
)

type BCType interface {
	GetTag() BCTypeTag
	IsPooledType() bool
	GetTypePool() *TypePool
	Equals(BCType) bool
	GetSize() uint32
	String() string
}

type bcTypeHeader struct {
	tpool     *TypePool
	tag       BCTypeTag
	inst      BCType
	size      uint32
	signature string
}

func (th *bcTypeHeader) GetTypePool() *TypePool {
	return th.tpool
}

func (th *bcTypeHeader) IsPooledType() bool {
	return th.tpool != nil
}

func (th *bcTypeHeader) GetTag() BCTypeTag {
	return th.tag
}

func (th *bcTypeHeader) GetSize() uint32 {
	return th.size
}

func (th *bcTypeHeader) Equals(other BCType) bool {
	if th.GetTag() != other.GetTag() {
		return false
	}
	if th.IsPooledType() && !other.IsPooledType() {
		return false
	}
	signatureA := th.String()
	signatureB := other.String()
	return signatureA == signatureB
}

func (th *bcTypeHeader) String() string {
	if th.signature != "" {
		return th.signature
	}
	typeSignature := ""

	stack := util.NewExpandingStack[any](256)
	stack.Push(th.inst)

	for stack.GetPointer() != -1 {
		v := stack.Pop()
		switch vv := v.(type) {
		case BCType:
			switch vv.GetTag() {
			case BCTYPE_BOOLEAN:
				typeSignature += "bool"
			case BCTYPE_I8:
				typeSignature += "int8"
			case BCTYPE_I16:
				typeSignature += "int16"
			case BCTYPE_I32:
				typeSignature += "int32"
			case BCTYPE_I64:
				typeSignature += "int64"
			case BCTYPE_U8:
				typeSignature += "uint8"
			case BCTYPE_U16:
				typeSignature += "uint16"
			case BCTYPE_U32:
				typeSignature += "uint32"
			case BCTYPE_U64:
				typeSignature += "uint64"
			case BCTYPE_F32:
				typeSignature += "float32"
			case BCTYPE_F64:
				typeSignature += "float64"
			case BCTYPE_UTFSTRING:
				typeSignature += "string"
			case BCTYPE_ARRAY:
				inst := vv.(*ArrayType)

				typeSignature += "array["
				stack.Push(2)
				stack.Push(inst.tpool.GetType(inst.elementTypeIndex))
			case BCTYPE_POINTER:
				inst := vv.(*ArrayType)

				typeSignature += "pointer["
				stack.Push(2)
				stack.Push(inst.tpool.GetType(inst.elementTypeIndex))
			case BCTYPE_STRUCT:
				inst := vv.(*StructType)
				typeSignature += "struct(\"" + inst.GetName() + "\")["
				stack.Push(2)
				types := inst.GetFieldTypeIndexes()
				typeCount := len(types)
				for i, _ := range types {
					stack.Push(inst.tpool.GetType(types[typeCount-i-1]))
					if i != typeCount-1 {
						stack.Push(0)
					}
				}
			case BCTYPE_MODULE:
				typeSignature += "module(\"" + vv.(*ModuleType).GetName() + "\")"
			case BCTYPE_FUNCTION:
				inst := vv.(*FunctionType)
				typeSignature += "function["
				stack.Push(2)
				types := inst.GetReturnTypeIndexes()
				typeCount := len(types)
				for i, _ := range types {
					stack.Push(inst.tpool.GetType(types[typeCount-i-1]))
					if i != typeCount-1 {
						stack.Push(0)
					}
				}
				stack.Push(1)
				stack.Push(2)
				types = inst.GetParamTypeIndexes()
				typeCount = len(types)
				for i, _ := range types {
					stack.Push(inst.tpool.GetType(types[typeCount-i-1]))
					if i != typeCount-1 {
						stack.Push(0)
					}
				}
			default:
				typeSignature += "unknown_type(" + strconv.Itoa(int(vv.GetTag())) + ")"
			}
		case int:
			switch v {
			case 0:
				typeSignature += ","
			case 1:
				typeSignature += "["
			case 2:
				typeSignature += "]"
			default:
				panic("invalid end number when making type signature!")
			}
		}

	}

	th.signature = typeSignature
	return typeSignature
}

var nilPool = NewTypePool()

type IntegerType struct {
	bcTypeHeader
	unsigned bool
}

func NewIntegerType(tpool *TypePool, tag BCTypeTag, unsigned bool, size uint32) *IntegerType {
	t := &IntegerType{unsigned: unsigned}
	t.tag = tag
	t.inst = t
	t.size = size
	t.tpool = tpool
	return t
}

func (t *IntegerType) IsUnsigned() bool {
	return t.unsigned
}

type FloatType struct {
	bcTypeHeader
}

func NewFloatType(tpool *TypePool, tag BCTypeTag, size uint32) *FloatType {
	ft := &FloatType{}
	ft.tag = tag
	ft.inst = ft
	ft.size = size
	ft.tpool = tpool
	return ft
}

type BooleanType struct {
	bcTypeHeader
}

func NewBooleanType(tpool *TypePool) BCType {
	t := &BooleanType{}
	t.tag = BCTYPE_BOOLEAN
	t.inst = t
	t.size = 1
	t.tpool = tpool
	return t
}

type PointerType struct {
	bcTypeHeader
	typeIdx int32
}

func NewPointerType(tpool *TypePool, typeIdx int32) *PointerType {
	ft := &PointerType{}
	ft.tag = BCTYPE_POINTER
	ft.typeIdx = typeIdx
	ft.tpool = tpool
	ft.size = 4 // size of uint32
	return ft
}

func (t *PointerType) GetTypeIndex() int32 {
	return t.typeIdx
}

type UTFStringType struct {
	bcTypeHeader
}

func NewUTFStringType(tpool *TypePool) *UTFStringType {
	t := &UTFStringType{}
	t.tag = BCTYPE_UTFSTRING
	t.inst = t
	t.tpool = tpool
	return t
}

type ArrayType struct {
	bcTypeHeader
	elementTypeIndex int32
}

func NewArrayType(tpool *TypePool, elementTypeIndex int32) *ArrayType {
	a := &ArrayType{elementTypeIndex: elementTypeIndex}
	a.tag = BCTYPE_ARRAY
	a.tpool = tpool
	a.inst = a
	return a
}

func (t *ArrayType) GetElementTypeIndex() int32 {
	return t.elementTypeIndex
}

type StructType struct {
	bcTypeHeader
	name             string
	fieldTypeIndexes []int32
}

func NewStructType(tpool *TypePool, name string, fieldTypeIndexes ...int32) *StructType {
	t := &StructType{name: name, fieldTypeIndexes: fieldTypeIndexes}
	t.tag = BCTYPE_STRUCT
	t.tpool = tpool
	t.inst = t
	return t
}

func (t *StructType) GetName() string {
	return t.name
}

func (t *StructType) GetFieldTypeIndexes() []int32 {
	return t.fieldTypeIndexes
}

type ModuleType struct {
	bcTypeHeader
	name string
}

func NewModuleType(tpool *TypePool, name string) *ModuleType {
	mt := &ModuleType{name: name}
	mt.tag = BCTYPE_MODULE
	mt.inst = mt
	mt.tpool = tpool
	return mt
}

func (t *ModuleType) GetName() string {
	return t.name
}

type FunctionType struct {
	bcTypeHeader
	paramTypeIndexes  []int32
	returnTypeIndexes []int32
}

func NewFunctionType(tpool *TypePool, paramTypeIndexes []int32, returnTypeIndexes ...int32) *FunctionType {
	ft := &FunctionType{paramTypeIndexes: paramTypeIndexes, returnTypeIndexes: returnTypeIndexes}
	ft.tag = BCTYPE_FUNCTION
	ft.tpool = tpool
	ft.inst = ft
	return ft
}

func (t *FunctionType) GetParamTypeIndexes() []int32 {
	return t.paramTypeIndexes
}

func (t *FunctionType) GetReturnTypeIndexes() []int32 {
	return t.returnTypeIndexes
}
