package ffi

import (
	"errors"
	"fmt"
	"focal-vm/internal/vm/runtime"
	"focal-vm/public/runtimeapi"
	"reflect"
	"unsafe"
)

func RuntimeValueToReflectionValue(v runtimeapi.Value, r reflect.Type) (reflect.Value, error) {
	var err error
	conversionFunction := func() any {
		switch r.Kind() {
		case reflect.Interface:
			return reflect.ValueOf(v.GetRawValue())
		case reflect.Int:
			if runtime.ValueIsInteger(v) {
				return runtime.GetValueAsInt(v)
			}
			return fmt.Errorf("Could not convert value to int, %v", v)
		case reflect.Uint:
			if runtime.ValueIsInteger(v) {
				return uint(runtime.GetValueAsInt(v))
			}
			return fmt.Errorf("Could not convert value to uint, %v", v)
		case reflect.Int8:
			if v.GetTag() == runtimeapi.ValueTagInt8 {
				return v.(*runtime.Int8Value).GetValue()
			}
			return fmt.Errorf("Could not convert value to int8, %v", v)
		case reflect.Uint8:
			if v.GetTag() == runtimeapi.ValueTagInt8 {
				return uint8(v.(*runtime.Int8Value).GetValue())
			}
			return fmt.Errorf("Could not convert value to uint8, %v", v)
		case reflect.Int16:
			if v.GetTag() == runtimeapi.ValueTagInt16 {
				return v.(*runtime.Int16Value).GetValue()
			}
			return fmt.Errorf("Could not convert value to int16, %v", v)
		case reflect.Uint16:
			if v.GetTag() == runtimeapi.ValueTagInt16 {
				return uint16(v.(*runtime.Int16Value).GetValue())
			}
			return fmt.Errorf("Could not convert value to uint16, %v", v)
		case reflect.Int32:
			if v.GetTag() == runtimeapi.ValueTagInt32 {
				return v.(*runtime.Int32Value).GetValue()
			}
			return fmt.Errorf("Could not convert value to int32, %v", v)
		case reflect.Uint32:
			if v.GetTag() == runtimeapi.ValueTagInt32 {
				return uint32(v.(*runtime.Int32Value).GetValue())
			}
			return fmt.Errorf("Could not convert value to uint32, %v", v)
		case reflect.Int64:
			if v.GetTag() == runtimeapi.ValueTagInt64 {
				return v.(*runtime.Int64Value).GetValue()
			}
			return fmt.Errorf("Could not convert value to int64, %v", v)
		case reflect.Uint64:
			if v.GetTag() == runtimeapi.ValueTagInt64 {
				return uint64(v.(*runtime.Int64Value).GetValue())
			}
			return fmt.Errorf("Could not convert value to uint64, %v", v)
		case reflect.Uintptr:
			if v.GetTag() == runtimeapi.ValueTagInt64 {
				return uintptr(v.(*runtime.Int64Value).GetValue())
			}
			return fmt.Errorf("Could not convert value to uintptr, %v", v)
		case reflect.Pointer:
			if v.GetTag() == runtimeapi.ValueTagInt64 {
				return uintptr(v.(*runtime.Int64Value).GetValue())
			}
			return fmt.Errorf("Could not convert value to pointer, %v", v)
		case reflect.UnsafePointer:
			if v.GetTag() == runtimeapi.ValueTagInt64 {
				return unsafe.Pointer(uintptr(v.(*runtime.Int64Value).GetValue()))
			}
			return fmt.Errorf("Could not convert value to unsafe pointer, %v", v)
		case reflect.Slice:
			if v.GetTag() == runtimeapi.ValueTagArray {
				elemType := r.Elem()
				runtimeElems := v.(*runtime.ArrayValue).GetValue()
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
			if v.GetTag() == runtimeapi.ValueTagUTF8String {
				return v.(*runtime.UTF8StringValue).GetValue()
			}
			return fmt.Errorf("Could not convert value to string, %v", v)
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
