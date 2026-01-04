package waspec

//var WASM_MAGIC = []byte{0x00, 0x61, 0x73, 0x6D}
//
//type WasmModule struct {
//	magic   [4]byte
//	version uint32
//}
//
//type WasmSectionType uint8
//
//const (
//	WasmSectionType_CustomSection WasmSectionType = iota
//	WasmSectionType_TypeSection
//	WasmSectionType_ImportSection
//	WasmSectionType_FunctionSection
//	WasmSectionType_TableSection
//	WasmSectionType_MemorySection
//	WasmSectionType_GlobalSection
//	WasmSectionType_ExportSection
//	WasmSectionType_StartSection
//	WasmSectionType_ElementSection
//	WasmSectionType_CodeSection
//	WasmSectionType_DataSection
//	WasmSectionType_DataCountSection
//	WasmSectionType_TagSection
//)
//
//type WasmSection interface {
//	GetType() WasmSectionType
//	GetLength() uint32
//}
//
//type WasmCustomSection struct {
//	id   WasmSectionType
//	len  uint32
//	name string
//}
//
//func (wc *WasmCustomSection) GetType() WasmSectionType {
//	return wc.id
//}
//
//func (wc *WasmCustomSection) GetLength() uint32 {
//	return wc.len
//}
//
//func (wc WasmCustomSection) GetName() string {
//	return wc.name
//}
//
//type WasmTypeSection struct {
//	id   WasmSectionType
//	len  uint32
//	name string
//}
//
//func (wt *WasmTypeSection) GetType() WasmSectionType {
//	return wt.id
//}
//
//func (wt *WasmTypeSection) GetLength() uint32 {
//	return wt.len
//}
//
//func (wt WasmTypeSection) GetName() string {
//	return wt.name
//}
//
//type WasmType interface {
//}
//
//type WasmRecursiveType struct {
//	length uint32
//	types  []WasmType
//}
