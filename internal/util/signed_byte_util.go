package util

func ReadI8LE(arr []byte, idx int32) int8 {
	return int8(arr[idx])
}

func ReadI8ArrLE(arr []byte, idx int32, size int32) []int8 {
	out := []int8{}
	for i := range size {
		out = append(out, int8(arr[idx+i]))
	}
	return out
}

func ReadI16LE(arr []byte, idx int32) int16 {
	return int16(arr[idx])<<8 | int16(arr[idx+1])
}

func ReadI16ArrLE(arr []byte, idx int32, size int32) []int16 {
	out := []int16{}
	byteCount := size * 2
	for i := int32(0); i < byteCount; i += 2 {
		out = append(out, ReadI16LE(arr, idx+i))
	}
	return out
}

func ReadI32LE(arr []byte, idx int32) int32 {
	return int32(ReadU16LE(arr, idx))<<16 | int32(ReadU16LE(arr, idx+2))
}

func ReadI32ArrLE(arr []byte, idx int32, size int32) []int32 {
	out := []int32{}
	byteCount := size * 4
	for i := int32(0); i < byteCount; i += 4 {
		out = append(out, ReadI32LE(arr, idx+i))
	}
	return out
}

func ReadI64LE(arr []byte, idx int32) int64 {
	return int64(ReadU32LE(arr, idx))<<32 | int64(ReadU32LE(arr, idx+4))
}

func ReadI64ArrLE(arr []byte, idx int32, size int32) []int64 {
	out := []int64{}
	byteCount := size * 8
	for i := int32(0); i < byteCount; i += 8 {
		out = append(out, ReadI64LE(arr, idx+i))
	}
	return out
}

func ReadVariableLEI32(arr []byte, idx int32, size int32) int32 {
	v := int32(0)
	for i := range size {
		v <<= 8
		v |= int32(arr[idx+i])
	}
	return v
}

func ReadVariableLEI64(arr []byte, idx int32, size int32) int64 {
	v := int64(0)
	for i := range size {
		v <<= 8
		v |= int64(arr[idx+i])
	}
	return v
}

func WriteI8LE(arr []byte, v int8) []byte {
	return append(arr, uint8(v))
}

func WriteI8ArrLE(arr []byte, v []int8) []byte {
	for _, b := range v {
		arr = WriteI8LE(arr, b)
	}
	return arr
}

func WriteI16LE(arr []byte, v int16) []byte {
	return WriteU16LE(arr, uint16(v))
}

func WriteI16ArrLE(arr []byte, v []int16) []byte {
	for i := range len(v) {
		arr = WriteI16LE(arr, v[i])
	}
	return arr
}

func WriteI32LE(arr []byte, v int32) []byte {
	return WriteU32LE(arr, uint32(v))
}

func WriteI32ArrLE(arr []byte, v []int32) []byte {
	for i := range len(v) {
		arr = WriteI32LE(arr, v[i])
	}
	return arr
}

func WriteI64LE(arr []byte, v int64) []byte {
	return WriteU64LE(arr, uint64(v))
}

func WriteI64ArrLE(arr []byte, v []int64) []byte {
	for i := range len(v) {
		arr = WriteI64LE(arr, v[i])
	}
	return arr
}

func WriteVariableLEI32(arr []byte, v int32, size int32) []byte {
	bytes := arr
	o := v
	for range size {
		bytes = append(bytes, uint8(o&0xFF))
		o >>= 8
	}
	return bytes
}

func WriteVariableLEI64(arr []byte, v int64, size int32) []byte {
	bytes := arr
	o := v
	for range size {
		bytes = append(bytes, uint8(o&0xFF))
		o >>= 8
	}
	return bytes
}
