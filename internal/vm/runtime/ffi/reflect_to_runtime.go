package ffi

import (
	"fmt"
	"focal-vm/internal/vm/runtime"
	"focal-vm/public/runtimeapi"
	"reflect"
)

func ReflectionValueToRuntimeValue(v reflect.Value) (runtimeapi.Value, error) {
	var err error

	switch v.Kind() {
	case reflect.Int:
		return runtime.NewInt64Value(v.Int()), nil
	case reflect.Uint:
		return runtime.NewInt64Value(int64(v.Uint())), nil
	case reflect.Uintptr:
		return runtime.NewInt64Value(int64(v.Pointer())), nil
	case reflect.Pointer:
		return runtime.NewInt64Value(int64(v.Pointer())), nil
	case reflect.UnsafePointer:
		return runtime.NewInt64Value(int64(v.UnsafeAddr())), nil
	case reflect.Int8:
		return runtime.NewInt8Value(int8(v.Int())), nil
	case reflect.Uint8:
		return runtime.NewInt8Value(int8(v.Int())), nil
	case reflect.Int16:
		return runtime.NewInt16Value(int16(v.Int())), nil
	case reflect.Uint16:
		return runtime.NewInt16Value(int16(v.Int())), nil
	case reflect.Int32:
		return runtime.NewInt32Value(int32(v.Int())), nil
	case reflect.Uint32:
		return runtime.NewInt32Value(int32(v.Int())), nil
	case reflect.Int64:
		return runtime.NewInt64Value(v.Int()), nil
	case reflect.Uint64:
		return runtime.NewInt64Value(v.Int()), nil
	case reflect.String:
		return runtime.NewUTF8StringValue(v.String()), nil
	case reflect.Bool:
		return runtime.NewBooleanValue(v.Bool()), nil
	case reflect.Float32:
		return runtime.NewFloat32Value(float32(v.Float())), nil
	case reflect.Float64:
		return runtime.NewFloat64Value(v.Float()), nil
	case reflect.Array:
		length := v.Len()
		values := make([]runtimeapi.Value, length)
		for i := range length {
			value := v.Index(i)
			values[i], err = ReflectionValueToRuntimeValue(value)
			if err != nil {
				return nil, err
			}
		}
		return runtime.NewArrayValue(values), nil
	case reflect.Slice:
		length := v.Len()
		values := make([]runtimeapi.Value, length)
		for i := range length {
			value := v.Index(i)
			values[i], err = ReflectionValueToRuntimeValue(value)
			if err != nil {
				return nil, err
			}
		}
		return runtime.NewArrayValue(values), nil
	case reflect.Interface:
		return nil, nil

	default:
		return nil, fmt.Errorf("unhandled type conversion for native type \"" + v.Type().Kind().String() + "\"")
	}
}
