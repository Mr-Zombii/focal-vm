package ffi

//func ReflectionValueToRuntimeValue(rtpool *rtvalue.RTValuePool, v reflect.Value) (rtvalue.RTValue, error) {
//	//var err error
//
//	switch v.Kind() {
//	case reflect.Int:
//		return rtpool.GetOrMakeRTValueI64(v.Int()), nil
//	case reflect.Uint:
//		return rtpool.GetOrMakeRTValueI64(int64(v.Uint())), nil
//	case reflect.Uintptr:
//		return rtpool.GetOrMakeRTValueI64(int64(v.Pointer())), nil
//	case reflect.Pointer:
//		return rtpool.GetOrMakeRTValueI64(int64(v.Pointer())), nil
//	case reflect.UnsafePointer:
//		return rtpool.GetOrMakeRTValueI64(int64(v.UnsafeAddr())), nil
//	case reflect.Int8:
//		return rtpool.GetOrMakeRTValueI8(int8(v.Int())), nil
//	case reflect.Uint8:
//		return rtpool.GetOrMakeRTValueI8(int8(v.Int())), nil
//	case reflect.Int16:
//		return rtpool.GetOrMakeRTValueI16(int16(v.Int())), nil
//	case reflect.Uint16:
//		return rtpool.GetOrMakeRTValueI16(int16(v.Int())), nil
//	case reflect.Int32:
//		return rtpool.GetOrMakeRTValueI32(int32(v.Int())), nil
//	case reflect.Uint32:
//		return rtpool.GetOrMakeRTValueI32(int32(v.Int())), nil
//	case reflect.Int64:
//		return rtpool.GetOrMakeRTValueI64(v.Int()), nil
//	case reflect.Uint64:
//		return rtpool.GetOrMakeRTValueI64(v.Int()), nil
//	case reflect.String:
//		return rtpool.GetOrMakeRTValueString(v.String()), nil
//	case reflect.Bool:
//		return rtpool.GetOrMakeRTValueBool(v.Bool()), nil
//	case reflect.Float32:
//		return rtpool.GetOrMakeRTValueF32(float32(v.Float())), nil
//	case reflect.Float64:
//		return rtpool.GetOrMakeRTValueF64(v.Float()), nil
//	//case reflect.Array:
//	//	length := v.Len()
//	//	elemType := v.Type().Elem()
//	//
//	//	values := make([]runtimeapi.Value, length)
//	//	for i := range length {
//	//		value := v.Index(i)
//	//		values[i], err = ReflectionValueToRuntimeValue(value)
//	//		if err != nil {
//	//			return nil, err
//	//		}
//	//	}
//	//	return runtime.NewArrayValue(values), nil
//	//case reflect.Slice:
//	//	length := v.Len()
//	//	values := make([]runtimeapi.Value, length)
//	//	for i := range length {
//	//		value := v.Index(i)
//	//		values[i], err = ReflectionValueToRuntimeValue(value)
//	//		if err != nil {
//	//			return nil, err
//	//		}
//	//	}
//	//	return runtime.NewArrayValue(values), nil
//	//case reflect.Interface:
//	//	return nil, nil
//
//	default:
//		return nil, fmt.Errorf("unhandled type conversion for native type \"" + v.Type().Kind().String() + "\"")
//	}
//}
