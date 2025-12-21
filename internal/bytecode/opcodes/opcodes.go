package opcodes

type OpCode uint8

const (
	OP_NOP OpCode = iota

	// Load & Store
	OP_CLOAD   // Load constant
	OP_VSTORE  // Store to local/global
	OP_VLOAD   // Load local/global
	OP_DGLOBAL // Define global
	OP_DLOCAL  // Define local
	OP_FLOAD   // Load function
	OP_ALOAD   // Load array element
	OP_ASTORE  // Store element to array

	OP_NEWARRAY

	// Stack
	OP_DUP // Duplicate stack value
	OP_SWP // Swap 2 stack values
	OP_POP // Discard top stack Value

	// Integer Math
	OP_IADD // Add Integer
	OP_ISUB // Subtract Integer
	OP_IMUL // Mutiply Integer
	OP_IDIV // Divide Integer
	OP_IMOD // Modulus

	// Basic Logic
	OP_EQ  // Equals
	OP_NEQ // Not Equals

	// Boolean Logic
	OP_LNOT // Logical Not
	OP_LOR  // Logical OR
	OP_LAND // Logical AND
	OP_LXOR // Logical XOR

	// Integer Logic & Comparison
	OP_ILT // Less Than Integer
	OP_IGT // Greater Than Integer
	OP_ILE // Less Than Equals Integer
	OP_IGE // Greater Than Equals Integer

	// Integer Bitwise
	OP_RSH  // Bit Shift Right
	OP_LSH  // Bit Shift Left
	OP_RRT  // Bit Rotate Right
	OP_LRT  // Bit Rotate Left
	OP_BNOT // Bitwise Not
	OP_BOR  // Bitwise OR
	OP_BAND // Bitwise AND
	OP_BXOR // Bitwise Xor

	// Float Math
	OP_FADD // Add Float
	OP_FSUB // Subtract Float
	OP_FMUL // Mutiply Float
	OP_FDIV // Divide Float

	// Float Logic & Comparison
	OP_FLT // Less Than Float
	OP_FGT // Greater Than Float
	OP_FLE // Less Than Equals Float
	OP_FGE // Greater Than Equals Float

	// Control Flow
	OP_RET // Return from Frame
	OP_CALL
	OP_TCALL // Call Frame by Name
	OP_SCALL // Call Child Frame by Name
	OP_JMP   // Relative Jump
	OP_BR    // Branch If False

	// VM & Native Control
	OP_HALT  // VM Exit
	OP_PRINT // Print Stack Value
	OP_NCALL // Native Call

	OPCODE_COUNT
)
