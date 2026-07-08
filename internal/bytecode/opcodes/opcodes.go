package opcodes

type OpCode uint8

const (
	OP_NOP        OpCode = iota
	OP_LOAD_CONST        // ✅ Load constant

	// Local Scope Instructions
	OP_DEFINE_LOCAL // ✅ Define local in current scope
	OP_LOCAL_EXISTS // ✅ Check if local is defined in current or parent scope
	OP_OWNS_LOCAL   // ✅ Check if local is owned by current scope
	OP_SET_LOCAL    // ✅ Set local in current or parent scope
	OP_GET_LOCAL    // ✅ Get local in current or parent scope

	// Global Scope Instructions
	OP_DEFINE_GLOBAL // ✅ Define local in global scope
	OP_GLOBAL_EXISTS // ✅ Check if local is defined in global or global's parent scope
	OP_OWNS_GLOBAL   // ✅ Check if local is owned by global scope
	OP_SET_GLOBAL    // ✅ Set global in current or global's parent scope
	OP_GET_GLOBAL    // ✅ Get global in current or global's parent scope

	OP_LOAD_STATIC_FUNCTION // ✅ Load static function (a static function is one that is not held in any scope)

	// Array Instructions
	OP_ARRAY_NEW   // ✅ Create new array
	OP_ARRAY_STORE // ✅ Store value in array at index
	OP_ARRAY_LOAD  // ✅ Get value from array at index

	// Object Instructions
	OP_OBJECT_NEW          // ✅ Create New ScopeValue/ObjectValue
	OP_OBJECT_SET_FIELD    // ✅ Set Object Field
	OP_OBJECT_GET_FIELD    // ✅ Get Object Field
	OP_OBJECT_DEFINE_FIELD // ✅ Define Object Field
	OP_OBJECT_HAS_FIELD    // ✅ Define Object Field

	// Tuple Instructions
	OP_TUPLE_NEW       // ✅ Create New Tuple
	OP_TUPLE_SET_VALUE // ✅ Get Tuple Set Value
	OP_TUPLE_GET_VALUE // ✅ Get Tuple Get Value

	// Object Instructions
	OP_STRUCT_NEW       // ✅ Create New StructValue
	OP_STRUCT_SET_FIELD // ✅ Set Struct Field
	OP_STRUCT_GET_FIELD // ✅ Get Struct Field

	// Stack
	OP_DUP // ✅ Duplicate stack value
	OP_SWP // ✅ Swap 2 stack values
	OP_POP // ✅ Discard top stack Value

	// Basic Logic
	OP_EQ  // ✅ Equals
	OP_NEQ // ✅ Not Equals

	// Boolean Logic
	OP_LNOT // ✅ Logical Not
	OP_LOR  // ✅ Logical OR
	OP_LAND // ✅ Logical AND
	OP_LXOR // ✅ Logical XOR

	// Integer Math
	OP_IADD // ✅ Add Integer
	OP_ISUB // ✅ Subtract Integer
	OP_IMUL // ✅ Mutiply Integer
	OP_IDIV // ✅ Divide Integer
	OP_IMOD // ✅ Modulus

	// Integer Logic & Comparison
	OP_ILT // ✅ Less Than Integer
	OP_IGT // ✅ Greater Than Integer
	OP_ILE // ✅ Less Than Equals Integer
	OP_IGE // ✅ Greater Than Equals Integer

	// Integer Bitwise
	OP_RSH  // ✅ Bit Shift Right
	OP_LSH  // ✅ Bit Shift Left
	OP_RRT  // ✅ Bit Rotate Right
	OP_LRT  // ✅ Bit Rotate Left
	OP_BNOT // ✅ Bitwise Not
	OP_BOR  // ✅ Bitwise OR
	OP_BAND // ✅ Bitwise AND
	OP_BXOR // ✅ Bitwise Xor

	// Numer Conversion
	OP_CONV_TO_I8  // ✅ Convert to I8
	OP_CONV_TO_I16 // ✅ Convert to I16
	OP_CONV_TO_I32 // ✅ Convert to I32
	OP_CONV_TO_I64 // ✅ Convert to I64
	OP_CONV_TO_F32 // ✅ Convert to F32
	OP_CONV_TO_F64 // ✅ Convert to F64

	// Float Math
	OP_FADD // ✅ Add Float
	OP_FSUB // ✅ Subtract Float
	OP_FMUL // ✅ Mutiply Float
	OP_FDIV // ✅ Divide Float

	// Float Logic & Comparison
	OP_FLT // ✅ Less Than Float
	OP_FGT // ✅ Greater Than Float
	OP_FLE // ✅ Less Than Equals Float
	OP_FGE // ✅ Greater Than Equals Float

	// Control Flow
	OP_RET    // ✅ Return from Frame
	OP_CALL   // ✅ Call on callable value
	OP_TCALL  // ✅ Call on callable value and reuse frame
	OP_JUMP   // ✅ Relative Jump
	OP_BRANCH // ✅ Branch If False

	// VM & Native Control
	OP_HALT // ✅ VM Exit

	OPCODE_COUNT
)
