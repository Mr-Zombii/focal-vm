package util

func ReadU8LE(arr []byte, idx int32) uint8 {
	return arr[idx]
}

func ReadU8ArrLE(arr []byte, idx int32, size int32) []uint8 {
	out := []uint8{}
	for i := range size {
		out = append(out, arr[idx+i])
	}
	return out
}

func ReadU16LE(arr []byte, idx int32) uint16 {
	return uint16(arr[idx])<<8 | uint16(arr[idx+1])
}

func ReadU16ArrLE(arr []byte, idx int32, size int32) []uint16 {
	out := []uint16{}
	byteCount := size * 2
	for i := int32(0); i < byteCount; i += 2 {
		out = append(out, ReadU16LE(arr, idx+i))
	}
	return out
}

func ReadU32LE(arr []byte, idx int32) uint32 {
	return uint32(ReadU16LE(arr, idx))<<16 | uint32(ReadU16LE(arr, idx+2))
}

func ReadU32ArrLE(arr []byte, idx int32, size int32) []uint32 {
	out := []uint32{}
	byteCount := size * 4
	for i := int32(0); i < byteCount; i += 4 {
		out = append(out, ReadU32LE(arr, idx+i))
	}
	return out
}

func ReadU64LE(arr []byte, idx int32) uint64 {
	return uint64(ReadU32LE(arr, idx))<<32 | uint64(ReadU32LE(arr, idx+4))
}

func ReadU64ArrLE(arr []byte, idx int32, size int32) []uint64 {
	out := []uint64{}
	byteCount := size * 8
	for i := int32(0); i < byteCount; i += 8 {
		out = append(out, ReadU64LE(arr, idx+i))
	}
	return out
}

func ReadVariableLEU32(arr []byte, idx int32, size int32) uint32 {
	v := uint32(0)
	for i := range size {
		v <<= 8
		v |= uint32(arr[idx+i])
	}
	return v
}

func ReadVariableLEU64(arr []byte, idx int32, size int32) uint64 {
	v := uint64(0)
	for i := range size {
		v <<= 8
		v |= uint64(arr[idx+i])
	}
	return v
}

func WriteU8LE(arr []byte, v uint8) []byte {
	return append(arr, v)
}

func WriteU8ArrLE(arr []byte, v []uint8) []byte {
	return append(arr, v...)
}

func WriteU16LE(arr []byte, v uint16) []byte {
	return append(append(arr, uint8(v&0xFF)), uint8((v&0xFF00)>>8))
}

func WriteU16ArrLE(arr []byte, v []uint16) []byte {
	for i := range len(v) {
		arr = WriteU16LE(arr, v[i])
	}
	return arr
}

func WriteU32LE(arr []byte, v uint32) []byte {
	return WriteU16LE(WriteU16LE(arr, uint16(v&0xFFFF)), uint16((v&0xFFFF0000)>>16))
}

func WriteU32ArrLE(arr []byte, v []uint32) []byte {
	for i := range len(v) {
		arr = WriteU32LE(arr, v[i])
	}
	return arr
}

func WriteU64LE(arr []byte, v uint64) []byte {
	return WriteU32LE(WriteU32LE(arr, uint32(v&0xFFFFFFFF)), uint32((v&0xFFFFFFFF00000000)>>32))
}

func WriteU64ArrLE(arr []byte, v []uint64) []byte {
	for i := range len(v) {
		arr = WriteU64LE(arr, v[i])
	}
	return arr
}

func WriteVariableLEU32(arr []byte, v uint32, size int32) []byte {
	bytes := arr
	o := v
	for range size {
		bytes = append(bytes, uint8(o&0xFF))
		o >>= 8
	}
	return bytes
}

func WriteVariableLEU64(arr []byte, v uint64, size int32) []byte {
	bytes := arr
	o := v
	for range size {
		bytes = append(bytes, uint8(o&0xFF))
		o >>= 8
	}
	return bytes
}
