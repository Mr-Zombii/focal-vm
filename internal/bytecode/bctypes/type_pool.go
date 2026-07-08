package bctypes

import (
	"errors"
	"fmt"
	"strconv"
)

type TypePool struct {
	bctypes []BCType
}

func NewTypePool() *TypePool {
	return &TypePool{
		[]BCType{},
	}
}

func (p *TypePool) String() string {
	return fmt.Sprint(p.bctypes)
}

func (p *TypePool) Size() int32 {
	return int32(len(p.bctypes))
}

func (p *TypePool) AddType(c BCType) int32 {
	existingIdx := p.GetTypeIdx(c)
	if existingIdx != -1 {
		return existingIdx
	}

	idx := len(p.bctypes)
	p.bctypes = append(p.bctypes, c)
	return int32(idx)
}

func (p *TypePool) ExpectType(idx int32, tag BCTypeTag) BCType {
	bctype := p.bctypes[idx]
	if bctype.GetTag() != tag {
		panic("Expected bctype " + strconv.Itoa(int(tag)) + " got " + strconv.Itoa(int(bctype.GetTag())))
	}
	return bctype
}

func (p *TypePool) GetTypeIdxFromSignature(signature string) int32 {
	for i, v := range p.bctypes {
		if v.String() == signature {
			return int32(i)
		}
	}
	return -1
}

func (p *TypePool) GetTypeFromSignature(signature string) (BCType, error) {
	for _, v := range p.bctypes {
		if v.String() == signature {
			return v, nil
		}
	}
	return nil, errors.New("Failed to find type " + signature)
}

func (p *TypePool) GetType(idx int32) BCType {
	return p.bctypes[idx]
}

func (p *TypePool) GetTypes() []BCType {
	return p.bctypes
}

func (p *TypePool) getTypeByTag(tag BCTypeTag) int32 {
	for i, v := range p.bctypes {
		if v.GetTag() == tag {
			return int32(i)
		}
	}
	return -1
}

func (p *TypePool) getOrCreateTypeByTag(tag BCTypeTag) (BCType, int32) {
	switch tag {
	case BCTYPE_I8:
		return p.GetOrCreateI8Type()
	case BCTYPE_I16:
		return p.GetOrCreateI16Type()
	case BCTYPE_I32:
		return p.GetOrCreateI32Type()
	case BCTYPE_I64:
		return p.GetOrCreateI64Type()
	case BCTYPE_F32:
		return p.GetOrCreateF32Type()
	case BCTYPE_F64:
		return p.GetOrCreateF64Type()
	case BCTYPE_U8:
		return p.GetOrCreateU8Type()
	case BCTYPE_U16:
		return p.GetOrCreateU16Type()
	case BCTYPE_U32:
		return p.GetOrCreateU32Type()
	case BCTYPE_U64:
		return p.GetOrCreateU64Type()
	case BCTYPE_UTFSTRING:
		return p.GetOrCreateUTFStringType()
	default:
		return nil, -1
	}
}

func (p *TypePool) GetTypesByTag(tag BCTypeTag) []int32 {
	var types []int32
	for i, v := range p.bctypes {
		if v.GetTag() == tag {
			types = append(types, int32(i))
		}
	}
	return types
}

func (p *TypePool) GetTypeIdx(refType BCType) int32 {
	if refType == nil {
		return -1
	}
	tag := refType.GetTag()

	switch tag {
	case BCTYPE_FUNCTION:
		_, v := p.GetFunctionType(refType.(*FunctionType))
		return v
	case BCTYPE_STRUCT:
		_, v := p.GetStructType(refType.(*StructType))
		return v
	case BCTYPE_ARRAY:
		_, v := p.GetArrayType(refType.(*ArrayType))
		return v
	case BCTYPE_MODULE:
		_, v := p.GetModuleType(refType.(*ModuleType).name)
		return v
	default:
		return p.getTypeByTag(tag)

	}
}

func (p *TypePool) CopyToPool(refType BCType) (BCType, int32) {
	if refType == nil {
		return nil, -1
	}
	var idx int32
	idx = p.GetTypeIdx(refType)
	if idx != -1 {
		return refType, idx
	}

	tag := refType.GetTag()
	switch tag {
	case BCTYPE_FUNCTION:
		return p.GetOrCreateFunctionType(refType.(*FunctionType))
	case BCTYPE_STRUCT:
		return p.GetOrCreateStructType(refType.(*StructType))
	case BCTYPE_ARRAY:
		return p.GetOrCreateArrayType(refType.(*ArrayType))
	case BCTYPE_MODULE:
		return p.GetOrCreateModuleType(refType.(*ModuleType).name)
	default:
		return p.getOrCreateTypeByTag(tag)
	}
}

func (p *TypePool) GetArrayType(reference *ArrayType) (BCType, int32) {
	if reference == nil {
		return nil, -1
	}
	elementType := reference.tpool.GetType(reference.elementTypeIndex)
	elementTypeIndex := p.GetTypeIdx(elementType)

	if elementTypeIndex == -1 {
		return nil, -1
	}

	for i, v := range p.bctypes {
		if at, ok := v.(*ArrayType); ok && at.elementTypeIndex == elementTypeIndex {
			return v, int32(i)
		}
	}
	return nil, -1
}

func (p *TypePool) GetPointerType(reference *PointerType) (BCType, int32) {
	if reference == nil {
		return nil, -1
	}
	typ := reference.tpool.GetType(reference.typeIdx)
	typeIdx := p.GetTypeIdx(typ)

	if typeIdx == -1 {
		return nil, -1
	}

	for i, v := range p.bctypes {
		if at, ok := v.(*PointerType); ok && at.typeIdx == typeIdx {
			return v, int32(i)
		}
	}
	return nil, -1
}

func (p *TypePool) GetStructType(reference *StructType) (BCType, int32) {
	if reference == nil {
		return nil, -1
	}
	for i, v := range p.bctypes {
		if st, ok := v.(*StructType); ok && st.name == reference.name && len(st.fieldTypeIndexes) == len(reference.fieldTypeIndexes) {
			matched := true
			for ftI, ft := range st.fieldTypeIndexes {
				if ft != reference.fieldTypeIndexes[ftI] {
					matched = false
					break
				}
			}
			if matched || reference.String() == v.String() {
				return v, int32(i)
			}
		}
	}
	return nil, -1
}

func (p *TypePool) GetFunctionType(reference *FunctionType) (BCType, int32) {
	if reference == nil {
		return nil, -1
	}
	for i, v := range p.bctypes {
		if ft, ok := v.(*FunctionType); ok && len(ft.paramTypeIndexes) == len(reference.paramTypeIndexes) && len(ft.returnTypeIndexes) == len(reference.returnTypeIndexes) {
			matched := true
			for tI, t := range ft.paramTypeIndexes {
				if t != reference.paramTypeIndexes[tI] {
					matched = false
					break
				}
			}
			for tI, t := range ft.returnTypeIndexes {
				if t != reference.returnTypeIndexes[tI] {
					matched = false
					break
				}
			}
			if matched || ft.String() == reference.String() {
				return v, int32(i)
			}
		}
	}
	return nil, -1
}

func (p *TypePool) GetModuleType(name string) (BCType, int32) {
	for i, v := range p.bctypes {
		if mt, ok := v.(*ModuleType); ok && mt.name == name {
			return v, int32(i)
		}
	}
	return nil, -1
}

func (p *TypePool) GetOrCreateI8Type() (BCType, int32) {
	idx := p.getTypeByTag(BCTYPE_I8)
	if idx == -1 {
		t := NewIntegerType(p, BCTYPE_I8, false, 1)
		return t, p.AddType(t)
	}
	return p.bctypes[idx], idx
}

func (p *TypePool) GetOrCreateI16Type() (BCType, int32) {
	idx := p.getTypeByTag(BCTYPE_I16)
	if idx == -1 {
		t := NewIntegerType(p, BCTYPE_I16, false, 2)
		return t, p.AddType(t)
	}
	return p.bctypes[idx], idx
}

func (p *TypePool) GetOrCreateI32Type() (BCType, int32) {
	idx := p.getTypeByTag(BCTYPE_I32)
	if idx == -1 {
		t := NewIntegerType(p, BCTYPE_I32, false, 4)
		return t, p.AddType(t)
	}
	return p.bctypes[idx], idx
}

func (p *TypePool) GetOrCreateI64Type() (BCType, int32) {
	idx := p.getTypeByTag(BCTYPE_I64)
	if idx == -1 {
		t := NewIntegerType(p, BCTYPE_I64, false, 8)
		return t, p.AddType(t)
	}
	return p.bctypes[idx], idx
}

func (p *TypePool) GetOrCreateU8Type() (BCType, int32) {
	idx := p.getTypeByTag(BCTYPE_U8)
	if idx == -1 {
		t := NewIntegerType(p, BCTYPE_U8, true, 1)
		return t, p.AddType(t)
	}
	return p.bctypes[idx], idx
}

func (p *TypePool) GetOrCreateU16Type() (BCType, int32) {
	idx := p.getTypeByTag(BCTYPE_U16)
	if idx == -1 {
		t := NewIntegerType(p, BCTYPE_U16, true, 2)
		return t, p.AddType(t)
	}
	return p.bctypes[idx], idx
}

func (p *TypePool) GetOrCreateU32Type() (BCType, int32) {
	idx := p.getTypeByTag(BCTYPE_U32)
	if idx == -1 {
		t := NewIntegerType(p, BCTYPE_U32, true, 4)
		return t, p.AddType(t)
	}
	return p.bctypes[idx], idx
}

func (p *TypePool) GetOrCreateU64Type() (BCType, int32) {
	idx := p.getTypeByTag(BCTYPE_U64)
	if idx == -1 {
		t := NewIntegerType(p, BCTYPE_U64, true, 8)
		return t, p.AddType(t)
	}
	return p.bctypes[idx], idx
}

func (p *TypePool) GetOrCreateF32Type() (BCType, int32) {
	idx := p.getTypeByTag(BCTYPE_F32)
	if idx == -1 {
		t := NewFloatType(p, BCTYPE_F32, 4)
		return t, p.AddType(t)
	}
	return p.bctypes[idx], idx
}

func (p *TypePool) GetOrCreateF64Type() (BCType, int32) {
	idx := p.getTypeByTag(BCTYPE_F64)
	if idx == -1 {
		t := NewFloatType(p, BCTYPE_F64, 8)
		return t, p.AddType(t)
	}
	return p.bctypes[idx], idx
}

func (p *TypePool) GetOrCreateBoolType() (BCType, int32) {
	idx := p.getTypeByTag(BCTYPE_BOOLEAN)
	if idx == -1 {
		t := NewBooleanType(p)
		return t, p.AddType(t)
	}
	return p.bctypes[idx], idx
}

func (p *TypePool) GetOrCreateUTFStringType() (BCType, int32) {
	idx := p.getTypeByTag(BCTYPE_UTFSTRING)
	if idx == -1 {
		t := NewUTFStringType(p)
		return t, p.AddType(t)
	}
	return p.bctypes[idx], idx
}

func (p *TypePool) GetOrCreateArrayType(reference *ArrayType) (BCType, int32) {
	if reference == nil {
		return nil, -1
	}
	v, idx := p.GetArrayType(reference)
	if idx == -1 {
		_, i := p.CopyToPool(reference.tpool.GetType(reference.elementTypeIndex))
		typ := NewArrayType(p, i)
		return typ, p.AddType(typ)
	}
	return v, idx
}

func (p *TypePool) GetOrCreatePointerType(reference *PointerType) (BCType, int32) {
	if reference == nil {
		return nil, -1
	}
	v, idx := p.GetPointerType(reference)
	if idx == -1 {
		_, i := p.CopyToPool(reference.tpool.GetType(reference.typeIdx))
		typ := NewPointerType(p, i)
		return typ, p.AddType(typ)
	}
	return v, idx
}

func (p *TypePool) GetOrCreateStructType(reference *StructType) (BCType, int32) {
	if reference == nil {
		return nil, -1
	}
	v, idx := p.GetStructType(reference)
	if idx == -1 {
		fieldTypeIndexes := make([]int32, len(reference.GetFieldTypeIndexes()))
		for i, oldI := range reference.GetFieldTypeIndexes() {
			_, fieldTypeIndexes[i] = p.CopyToPool(reference.tpool.GetType(oldI))
		}
		typ := NewStructType(p, reference.name, fieldTypeIndexes...)
		return typ, p.AddType(typ)
	}
	return v, idx
}

func (p *TypePool) GetOrCreateModuleType(name string) (BCType, int32) {
	v, idx := p.GetModuleType(name)
	if idx == -1 {
		typ := NewModuleType(p, name)
		return typ, p.AddType(typ)
	}
	return v, idx
}

func (p *TypePool) GetOrCreateFunctionType(reference *FunctionType) (BCType, int32) {
	if reference == nil {
		return nil, -1
	}
	v, idx := p.GetFunctionType(reference)
	if idx == -1 {
		paramTypeIndexes := make([]int32, len(reference.GetParamTypeIndexes()))
		for i, oldI := range reference.GetParamTypeIndexes() {
			_, paramTypeIndexes[i] = p.CopyToPool(reference.tpool.GetType(oldI))
		}
		returnTypeIndexes := make([]int32, len(reference.GetReturnTypeIndexes()))
		for i, oldI := range reference.GetReturnTypeIndexes() {
			_, returnTypeIndexes[i] = p.CopyToPool(reference.tpool.GetType(oldI))
		}
		typ := NewFunctionType(p, paramTypeIndexes, returnTypeIndexes...)
		return typ, p.AddType(typ)
	}
	return v, idx
}
