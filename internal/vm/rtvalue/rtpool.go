package rtvalue

import (
	"fmt"
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/bytecode/constants"
	"focal-vm/internal/bytecode/spec"
	"focal-vm/internal/erroring"
	"focal-vm/internal/vm/runtime/allocator"
	"reflect"
	"strconv"
)

type RTValuePool struct {
	allocator allocator.Allocator
	values    []RTValue
}

func NewRTValuePool(allocator allocator.Allocator) *RTValuePool {
	return &RTValuePool{values: []RTValue{}, allocator: allocator}
}

func (vp *RTValuePool) ResetMarks() {
	for _, v := range vp.values {
		v.ResetMark()
	}
}

func (vp *RTValuePool) CleanAndTransfer(survivalThreshold int32, nextGen *RTValuePool) {
	//var newValues []RTValue
	//for i, v := range vp.values {
	//	if v.IsMarked() {
	//		if v.GetSurvivedCycles() >= survivalThreshold {
	//			nextGen.add(v)
	//			v.ResetSurvivedCycles()
	//		} else {
	//			newValues = append(newValues, vp.values[i])
	//			v.IncSurvivedCycles()
	//		}
	//	}
	//}
	//vp.values = newValues
}

func (vp *RTValuePool) RemoveAndFree(idx int32) {
	v := vp.values[idx]

	if vp.allocator.IsInvalidOrFree(v) {
		erroring.GlobalErrorHandler.Panic("RTValue Lifetime Error", "RTValue that was already free has been kept and freed twice!!")
		return
	}

	vp.values = append(vp.values[:idx], vp.values[idx+1:]...)
	for i, va := range vp.values {
		va.SetIdx(int32(i))
	}

	v.GetHeader().postFree()
	vp.allocator.Free(v)
}

func (vp *RTValuePool) Clean() {
	var newValues []RTValue
	vpValueCount := len(vp.values)
	for i := 0; i < vpValueCount; i++ {
		v := vp.values[i]

		if v.GetRefCount() <= 0 {
			if vp.allocator.IsInvalidOrFree(v) {
				erroring.GlobalErrorHandler.Panic("RTValue Lifetime Error", "RTValue that was already free has been kept and freed twice!!")
				return
			}
			v.GetHeader().postFree()
			vp.allocator.Free(v)
			continue
		} else {
			newValues = append(newValues, v)
			continue
		}
	}
	vp.values = newValues
}

func (vp *RTValuePool) String() string {
	return fmt.Sprint(vp.values)
}

func (vp *RTValuePool) Length() int {
	return len(vp.values)
}

func (vp *RTValuePool) GetValues() []RTValue {
	return vp.values
}

func (vp *RTValuePool) add(value RTValue) RTValue {
	vp.values = append(vp.values, value)
	value.SetPool(vp)
	value.SetIdx(int32(len(vp.values) - 1))
	value.IncRefCount()
	return value
}

func (vp *RTValuePool) CreateOrGetFromConstant(constant constants.Constant) RTValue {
	switch constant.GetTag() {
	case constants.ConstantTagInt8:
		return vp.GetOrMakeRTValueI8(constant.(*constants.ConstantInt8).GetValue())
	case constants.ConstantTagInt16:
		return vp.GetOrMakeRTValueI16(constant.(*constants.ConstantInt16).GetValue())
	case constants.ConstantTagInt32:
		return vp.GetOrMakeRTValueI32(constant.(*constants.ConstantInt32).GetValue())
	case constants.ConstantTagInt64:
		return vp.GetOrMakeRTValueI64(constant.(*constants.ConstantInt64).GetValue())
	case constants.ConstantTagFloat32:
		return vp.GetOrMakeRTValueF32(constant.(*constants.ConstantFloat32).GetValue())
	case constants.ConstantTagFloat64:
		return vp.GetOrMakeRTValueF64(constant.(*constants.ConstantFloat64).GetValue())
	case constants.ConstantTagBoolean:
		return vp.GetOrMakeRTValueBool(constant.(*constants.ConstantBoolean).GetValue())
	case constants.ConstantTagUTF8String:
		return vp.GetOrMakeRTValueString(constant.(*constants.ConstantUTF8String).GetValue())
	default:
		panic("Unsupported constant tag " + strconv.Itoa(int(constant.GetTag())))
	}
}

func (vp *RTValuePool) GetRTValueBool(value bool) *RTValueBool {
	for _, v := range vp.values {
		if rt, ok := v.(*RTValueBool); ok && rt.value == value {
			rt.IncRefCount()
			return rt
		}
	}
	return nil
}

func (vp *RTValuePool) GetRTValueI8(value int8) *RTValueI8 {
	for _, v := range vp.values {
		if rt, ok := v.(*RTValueI8); ok && rt.value == value {
			rt.IncRefCount()
			return rt
		}
	}
	return nil
}

func (vp *RTValuePool) GetRTValueI16(value int16) *RTValueI16 {
	for _, v := range vp.values {
		if rt, ok := v.(*RTValueI16); ok && rt.value == value {
			rt.IncRefCount()
			return rt
		}
	}
	return nil
}

func (vp *RTValuePool) GetRTValueI32(value int32) *RTValueI32 {
	for _, v := range vp.values {
		if rt, ok := v.(*RTValueI32); ok && rt.value == value {
			rt.IncRefCount()
			return rt
		}
	}
	return nil
}

func (vp *RTValuePool) GetRTValueI64(value int64) *RTValueI64 {
	for _, v := range vp.values {
		if rt, ok := v.(*RTValueI64); ok && rt.value == value {
			rt.IncRefCount()
			return rt
		}
	}
	return nil
}

func (vp *RTValuePool) GetRTValueF32(value float32) *RTValueF32 {
	for _, v := range vp.values {
		if rt, ok := v.(*RTValueF32); ok && rt.value == value {
			rt.IncRefCount()
			return rt
		}
	}
	return nil
}

func (vp *RTValuePool) GetRTValueF64(value float64) *RTValueF64 {
	for _, v := range vp.values {
		if rt, ok := v.(*RTValueF64); ok && rt.value == value {
			rt.IncRefCount()
			return rt
		}
	}
	return nil
}

func (vp *RTValuePool) GetRTValueString(value string) *RTValueString {
	for _, v := range vp.values {
		if rt, ok := v.(*RTValueString); ok && rt.GetValue() == value {
			rt.IncRefCount()
			return rt
		}
	}
	return nil
}

func (vp *RTValuePool) GetOrMakeRTValueBool(value bool) *RTValueBool {
	v := vp.GetRTValueBool(value)
	if v == nil {
		va := newRTValueBool(vp.allocator, value)
		vp.add(va)
		return va
	}
	return v
}

func (vp *RTValuePool) GetOrMakeRTValueI8(value int8) *RTValueI8 {
	v := vp.GetRTValueI8(value)
	if v == nil {
		va := newRTValueI8(vp.allocator, value)
		vp.add(va)
		return va
	}
	return v
}

func (vp *RTValuePool) GetOrMakeRTValueI16(value int16) *RTValueI16 {
	v := vp.GetRTValueI16(value)
	if v == nil {
		va := newRTValueI16(vp.allocator, value)
		vp.add(va)
		return va
	}
	return v
}

func (vp *RTValuePool) GetOrMakeRTValueI32(value int32) *RTValueI32 {
	v := vp.GetRTValueI32(value)
	if v == nil {
		va := newRTValueI32(vp.allocator, value)
		vp.add(va)
		return va
	}
	return v
}

func (vp *RTValuePool) GetOrMakeRTValueI64(value int64) *RTValueI64 {
	v := vp.GetRTValueI64(value)
	if v == nil {
		va := newRTValueI64(vp.allocator, value)
		vp.add(va)
		return va
	}
	return v
}

func (vp *RTValuePool) GetOrMakeRTValueF32(value float32) *RTValueF32 {
	v := vp.GetRTValueF32(value)
	if v == nil {
		va := newRTValueF32(vp.allocator, value)
		vp.add(va)
		return va
	}
	return v
}

func (vp *RTValuePool) GetOrMakeRTValueF64(value float64) *RTValueF64 {
	v := vp.GetRTValueF64(value)
	if v == nil {
		va := newRTValueF64(vp.allocator, value)
		vp.add(va)
		return va
	}
	return v
}

func (vp *RTValuePool) GetOrMakeRTValueString(value string) *RTValueString {
	v := vp.GetRTValueString(value)
	if v == nil {
		va := newRTValueString(vp.allocator, value)
		vp.add(va)
		return va
	}
	return v
}

func (vp *RTValuePool) CreateArray(elemType bctypes.BCType, length int32) *RTValueArray {
	va := newRTValueArray(vp.allocator, elemType, length)
	vp.add(va)
	return va
}

func (vp *RTValuePool) CreateVMFunction(function *spec.BCFunction) *RTValueVMFunction {
	va := newRTValueVMFunction(vp.allocator, function)
	vp.add(va)
	return va
}

func (vp *RTValuePool) CreateGOFunction(fnType bctypes.BCType, function interface{}) *RTValueGOFunction {
	if fnType.GetTag() != bctypes.BCTYPE_FUNCTION {
		erroring.GlobalErrorHandler.Panic("RtPool", fmt.Sprintf("Expected function type, but was provided type '%s", fnType))
		return nil
	}

	fnReflectType := reflect.TypeOf(function)
	if fnReflectType.Kind() != reflect.Func {
		erroring.GlobalErrorHandler.Panic("RtPool", fmt.Sprintf("Expected go function, but was provided '%s", function))
		return nil
	}

	if fnReflectType.IsVariadic() {
		erroring.GlobalErrorHandler.Panic("RtPool", fmt.Sprintf("Go Function '%s' cannot be a variadic function, please rewrite your function or add a different function.", function))
		return nil
	}

	va := newRTValueGOFunction(vp.allocator, fnType.(*bctypes.FunctionType), function)
	vp.add(va)
	return va
}

func (vp *RTValuePool) CreateStruct(structType bctypes.BCType) *RTValueStruct {
	if structType.GetTag() != bctypes.BCTYPE_STRUCT {
		erroring.GlobalErrorHandler.Panic("RtPool", fmt.Sprintf("Expected struct type, but was provided type '%s", structType))
		return nil
	}

	va := newRTValueStruct(vp.allocator, structType.(*bctypes.StructType))
	vp.add(va)
	return va
}
