package opcodes

type OpCode uint8

const (
	OP_NOP OpCode = iota
	// Constants
	OP_CLOAD1 // Load Constant 1 byte
	OP_CLOAD2 // Load Constant 2 bytes
	OP_CLOAD3 // Load Constant 3 bytes
	OP_CLOAD4 // Load Constant 4 bytes

	OP_LSTORE1 // Store Local
	OP_LSTORE2 // Store Local
	OP_LSTORE3 // Store Local
	OP_LSTORE4 // Store Local
	OP_LLOAD1  // Load Local
	OP_LLOAD2  // Load Local
	OP_LLOAD3  // Load Local
	OP_LLOAD4  // Load Local

	OP_GLOAD  // Load from Global
	OP_GSTORE // Store to Global

	OP_DUP // Duplicate Stack Value
	OP_SWP // Swap 2 Stack Values

	// Integer Math
	OP_IADD // Add Integer
	OP_ISUB // Subtract Integer
	OP_IMUL // Mutiply Integer
	OP_IDIV // Divide Integer
	OP_MOD  // Modulus

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
	OP_CALL11
	OP_CALL12
	OP_CALL13
	OP_CALL14
	OP_CALL21
	OP_CALL22
	OP_CALL23
	OP_CALL24
	OP_CALL31
	OP_CALL32
	OP_CALL33
	OP_CALL34
	OP_CALL41
	OP_CALL42
	OP_CALL43
	OP_CALL44
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
