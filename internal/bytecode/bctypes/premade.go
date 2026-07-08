package bctypes

var I8 = NewIntegerType(nilPool, BCTYPE_I8, false, 1)
var I16 = NewIntegerType(nilPool, BCTYPE_I16, false, 2)
var I32 = NewIntegerType(nilPool, BCTYPE_I32, false, 4)
var I64 = NewIntegerType(nilPool, BCTYPE_I64, false, 8)

var U8 = NewIntegerType(nilPool, BCTYPE_U8, true, 1)
var U16 = NewIntegerType(nilPool, BCTYPE_U16, true, 2)
var U32 = NewIntegerType(nilPool, BCTYPE_U32, true, 4)
var U64 = NewIntegerType(nilPool, BCTYPE_U64, true, 8)

var F32 = NewFloatType(nilPool, BCTYPE_F32, 4)
var F64 = NewFloatType(nilPool, BCTYPE_F64, 8)

var BOOL = NewBooleanType(nilPool)

var UTF_STRING = NewUTFStringType(nilPool)
