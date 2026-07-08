package ffi

import (
	"errors"
	"fmt"
	"focal-vm/internal/vm/rtvalue"
	"reflect"
	"unsafe"
)

func RuntimeValueToReflectionValue(v rtvalue.RTValue, r reflect.Type) (reflect.Value, error) {
	var err error
	conversionFunction := func() any {
		switch r.Kind() {
		case reflect.Interface:
			switch v.GetTag() {
			case rtvalue.RTValueTag_I8:
				return reflect.ValueOf(v.(*rtvalue.RTValueI8).GetValue())
			case rtvalue.RTValueTag_I16:
				return reflect.ValueOf(v.(*rtvalue.RTValueI16).GetValue())
			case rtvalue.RTValueTag_I32:
				return reflect.ValueOf(v.(*rtvalue.RTValueI32).GetValue())
			case rtvalue.RTValueTag_I64:
				return reflect.ValueOf(v.(*rtvalue.RTValueI64).GetValue())
			case rtvalue.RTValueTag_F32:
				return reflect.ValueOf(v.(*rtvalue.RTValueF32).GetValue())
			case rtvalue.RTValueTag_F64:
				return reflect.ValueOf(v.(*rtvalue.RTValueF64).GetValue())
			case rtvalue.RTValueTag_BOOL:
				return reflect.ValueOf(v.(*rtvalue.RTValueBool).GetValue())
			default:
				panic("unhandled default case")
			}
		case reflect.Int:
			if rtvalue.ValueIsInteger(v) {
				return rtvalue.GetValueAsInt(v)
			}
			return fmt.Errorf("Could not convert value to int, %v", v)
		case reflect.Uint:
			if rtvalue.ValueIsInteger(v) {
				return uint(rtvalue.GetValueAsInt(v))
			}
			return fmt.Errorf("Could not convert value to uint, %v", v)
		case reflect.Int8:
			if v.GetTag() == rtvalue.RTValueTag_I8 {
				return v.(*rtvalue.RTValueI8).GetValue()
			}
			return fmt.Errorf("Could not convert value to int8, %v", v)
		case reflect.Uint8:
			if v.GetTag() == rtvalue.RTValueTag_I8 {
				return uint8(v.(*rtvalue.RTValueI8).GetValue())
			}
			return fmt.Errorf("Could not convert value to uint8, %v", v)
		case reflect.Int16:
			if v.GetTag() == rtvalue.RTValueTag_I16 {
				return v.(*rtvalue.RTValueI16).GetValue()
			}
			return fmt.Errorf("Could not convert value to int16, %v", v)
		case reflect.Uint16:
			if v.GetTag() == rtvalue.RTValueTag_I16 {
				return uint16(v.(*rtvalue.RTValueI16).GetValue())
			}
			return fmt.Errorf("Could not convert value to uint16, %v", v)
		case reflect.Int32:
			if v.GetTag() == rtvalue.RTValueTag_I32 {
				return v.(*rtvalue.RTValueI32).GetValue()
			}
			return fmt.Errorf("Could not convert value to int32, %v", v)
		case reflect.Uint32:
			if v.GetTag() == rtvalue.RTValueTag_I32 {
				return uint32(v.(*rtvalue.RTValueI32).GetValue())
			}
			return fmt.Errorf("Could not convert value to uint32, %v", v)
		case reflect.Int64:
			if v.GetTag() == rtvalue.RTValueTag_I64 {
				return v.(*rtvalue.RTValueI64).GetValue()
			}
			return fmt.Errorf("Could not convert value to int64, %v", v)
		case reflect.Uint64:
			if v.GetTag() == rtvalue.RTValueTag_I64 {
				return uint64(v.(*rtvalue.RTValueI64).GetValue())
			}
			return fmt.Errorf("Could not convert value to uint64, %v", v)
		case reflect.Uintptr:
			if v.GetTag() == rtvalue.RTValueTag_I64 {
				return uintptr(v.(*rtvalue.RTValueI64).GetValue())
			}
			return fmt.Errorf("Could not convert value to uintptr, %v", v)
		case reflect.Pointer:
			if v.GetTag() == rtvalue.RTValueTag_I64 {
				return uintptr(v.(*rtvalue.RTValueI64).GetValue())
			}
			return fmt.Errorf("Could not convert value to pointer, %v", v)
		case reflect.UnsafePointer:
			if v.GetTag() == rtvalue.RTValueTag_I64 {
				return unsafe.Pointer(uintptr(v.(*rtvalue.RTValueI64).GetValue()))
			}
			return fmt.Errorf("Could not convert value to unsafe pointer, %v", v)
		case reflect.Slice:
			if v.GetTag() == rtvalue.RTValueTag_ARRAY {
				elemType := r.Elem()
				runtimeElems := v.(*rtvalue.RTValueArray).GetValue()
				elems := make([]reflect.Value, len(runtimeElems))
				for i := range runtimeElems {
					elems[i], err = RuntimeValueToReflectionValue(runtimeElems[i], elemType)
					if err != nil {
						return err
					}
				}
				nSlice := reflect.MakeSlice(reflect.SliceOf(elemType), 0, 0)
				nSlice = reflect.Append(nSlice, elems...)
				return nSlice
			} else if v != nil {
				elemType := r.Elem()
				elems := make([]reflect.Value, 1)
				elems[0], err = RuntimeValueToReflectionValue(v, elemType)
				if err != nil {
					return err
				}
				nSlice := reflect.MakeSlice(reflect.SliceOf(elemType), 0, 0)
				nSlice = reflect.Append(nSlice, elems...)
				return nSlice
			}
			return fmt.Errorf("Could not convert value to slice, %v", v)
		case reflect.String:
			if v.GetTag() == rtvalue.RTValueTag_STRING {
				return v.(*rtvalue.RTValueString).GetValue()
			}
			return fmt.Errorf("could not convert value to string, %v", v)
		default:
			return fmt.Errorf("unhandled type conversion for runtime value %v to %v", reflect.TypeOf(v), r)
		}
	}
	value := conversionFunction()
	if reflect.TypeOf(value) == reflect.TypeOf(errors.New("error")) {
		return reflect.ValueOf(-1), value.(error)
	}
	if _, ok := value.(reflect.Value); !ok {
		return reflect.ValueOf(value), nil
	}
	return value.(reflect.Value), nil
}
