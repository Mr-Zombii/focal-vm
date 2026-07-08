package ffi

import (
	"fmt"
	"focal-vm/internal/bytecode/bctypes"
	"focal-vm/internal/vm/rtvalue"
	"reflect"
	"unsafe"
)

func ReflectionValueToRuntimeValue(rtpool *rtvalue.RTValuePool, tpool *bctypes.TypePool, v reflect.Value) (rtvalue.RTValue, error) {
	var err error

	switch v.Kind() {
	case reflect.Int:
		return rtpool.GetOrMakeRTValueI64(v.Int()), nil
	case reflect.Uint:
		return rtpool.GetOrMakeRTValueI64(int64(v.Uint())), nil
	case reflect.Uintptr:
		return rtpool.GetOrMakeRTValueI64(int64(v.Pointer())), nil
	case reflect.Pointer:
		return rtpool.GetOrMakeRTValueI64(int64(v.Pointer())), nil
	case reflect.UnsafePointer:
		return rtpool.GetOrMakeRTValueI64(int64(v.UnsafeAddr())), nil
	case reflect.Int8:
		return rtpool.GetOrMakeRTValueI8(int8(v.Int())), nil
	case reflect.Uint8:
		return rtpool.GetOrMakeRTValueI8(int8(v.Int())), nil
	case reflect.Int16:
		return rtpool.GetOrMakeRTValueI16(int16(v.Int())), nil
	case reflect.Uint16:
		return rtpool.GetOrMakeRTValueI16(int16(v.Int())), nil
	case reflect.Int32:
		return rtpool.GetOrMakeRTValueI32(int32(v.Int())), nil
	case reflect.Uint32:
		return rtpool.GetOrMakeRTValueI32(int32(v.Int())), nil
	case reflect.Int64:
		return rtpool.GetOrMakeRTValueI64(v.Int()), nil
	case reflect.Uint64:
		return rtpool.GetOrMakeRTValueI64(v.Int()), nil
	case reflect.String:
		return rtpool.GetOrMakeRTValueString(v.String()), nil
	case reflect.Bool:
		return rtpool.GetOrMakeRTValueBool(v.Bool()), nil
	case reflect.Float32:
		return rtpool.GetOrMakeRTValueF32(float32(v.Float())), nil
	case reflect.Float64:
		return rtpool.GetOrMakeRTValueF64(v.Float()), nil
	case reflect.Array:
		length := v.Len()
		//elemType := v.Type().Elem()

		values := make([]rtvalue.RTValue, length)
		for i := range length {
			value := v.Index(i)
			values[i], err = ReflectionValueToRuntimeValue(rtpool, tpool, value)
			if err != nil {
				return nil, err
			}
		}
		t, err := ReflectionTypeToRuntimeType(tpool, v.Type())
		if err != nil {
			return nil, err
		}
		return rtpool.CreateArray(t, int32(length)), nil
	case reflect.Slice:
		length := v.Len()
		values := make([]rtvalue.RTValue, length)
		for i := range length {
			value := v.Index(i)
			values[i], err = ReflectionValueToRuntimeValue(rtpool, tpool, value)
			if err != nil {
				return nil, err
			}
		}
		t, err := ReflectionTypeToRuntimeType(tpool, v.Type())
		if err != nil {
			return nil, err
		}
		return rtpool.CreateArray(t, int32(length)), nil
	case reflect.Interface:
		return nil, nil

	default:
		return nil, fmt.Errorf("unhandled value conversion for native type \"" + v.Type().Kind().String() + "\"")
	}
}

func ReflectionTypeToRuntimeType(tpool *bctypes.TypePool, v reflect.Type) (bctypes.BCType, error) {
	switch v.Kind() {
	case reflect.Int:
		t, _ := tpool.GetOrCreateI32Type()
		return t, nil
	case reflect.Uint:
		t, _ := tpool.GetOrCreateU32Type()
		return t, nil
	case reflect.Uintptr:
		if unsafe.Sizeof(uintptr(0)) == 4 {
			t, _ := tpool.GetOrCreateU32Type()
			return t, nil
		}
		t, _ := tpool.GetOrCreateU64Type()
		return t, nil
	case reflect.Pointer:
		if unsafe.Sizeof(uintptr(0)) == 4 {
			t, _ := tpool.GetOrCreateU32Type()
			return t, nil
		}
		t, _ := tpool.GetOrCreateU64Type()
		return t, nil
	case reflect.UnsafePointer:
		if unsafe.Sizeof(uintptr(0)) == 4 {
			t, _ := tpool.GetOrCreateU32Type()
			return t, nil
		}
		t, _ := tpool.GetOrCreateU64Type()
		return t, nil
	case reflect.Int8:
		t, _ := tpool.GetOrCreateI8Type()
		return t, nil
	case reflect.Uint8:
		t, _ := tpool.GetOrCreateI8Type()
		return t, nil
	case reflect.Int16:
		t, _ := tpool.GetOrCreateI16Type()
		return t, nil
	case reflect.Uint16:
		t, _ := tpool.GetOrCreateI16Type()
		return t, nil
	case reflect.Int32:
		t, _ := tpool.GetOrCreateI32Type()
		return t, nil
	case reflect.Uint32:
		t, _ := tpool.GetOrCreateI32Type()
		return t, nil
	case reflect.Int64:
		t, _ := tpool.GetOrCreateI64Type()
		return t, nil
	case reflect.Uint64:
		t, _ := tpool.GetOrCreateI64Type()
		return t, nil
	case reflect.Float32:
		t, _ := tpool.GetOrCreateF32Type()
		return t, nil
	case reflect.Float64:
		t, _ := tpool.GetOrCreateF64Type()
		return t, nil
	case reflect.String:
		t, _ := tpool.GetOrCreateUTFStringType()
		return t, nil
	case reflect.Array:
		elemType := v.Elem()

		t, err := ReflectionTypeToRuntimeType(tpool, elemType)
		if err != nil {
			return nil, err
		}
		tt, _ := tpool.GetOrCreateArrayTypeFromElemType(t)

		return tt, nil
	case reflect.Slice:
		elemType := v.Elem()

		t, err := ReflectionTypeToRuntimeType(tpool, elemType)
		if err != nil {
			return nil, err
		}
		tt, _ := tpool.GetOrCreateArrayTypeFromElemType(t)

		return tt, nil
	case reflect.Interface:
		return nil, nil
	case reflect.Func:
		inIndexes := make([]int32, 0)
		outIndexes := make([]int32, 0)

		for i := range v.NumIn() {
			arg := v.In(i)
			if !(arg.String() == "runtimeapi.VM" || arg.String() == "rtvalue.RTValue") {
				t, err := ReflectionTypeToRuntimeType(tpool, arg)
				if err != nil {
					return nil, err
				}
				inIndexes = append(inIndexes, tpool.GetTypeIdx(t))
			}
		}
		for i := range v.NumOut() {
			arg := v.Out(i)
			if !(arg.String() == "runtimeapi.VM" || arg.String() == "rtvalue.RTValue") {
				t, err := ReflectionTypeToRuntimeType(tpool, arg)
				if err != nil {
					return nil, err
				}
				outIndexes = append(outIndexes, tpool.GetTypeIdx(t))
			}
		}
		t, _ := tpool.GetOrCreateFunctionType(bctypes.NewFunctionType(tpool, inIndexes, outIndexes...))
		return t, nil

	default:
		return nil, fmt.Errorf("unhandled type conversion for native type \"" + v.Kind().String() + "\"")
	}
}
