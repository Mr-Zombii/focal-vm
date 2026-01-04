package main

import (
	"focal-vm/internal/bytecode/bcio"
	"focal-vm/internal/bytecode/constants"
	"focal-vm/internal/bytecode/opcodes"
	"focal-vm/internal/bytecode/spec"
	"os"
)

func main() {
	/*
		fn fib(n) {
			fn aux(m, a, b) {
				if (m == 0) {
					return a
				}
				return aux(m - 1, b, a + b)
			}
			return aux(n, 0, 1)
		}

		fn main() {
			print(fib(10))
		}
	*/

	/*
		fn fib$aux(m, a, b) {
			if (m == 0) {
				return a
			}
			return fib$aux(m - 1, b, a + b)
		}

		fn fib(n) {
			return fib$aux(n, 0, 1)
		}

		fn main() {
			print(fib(10))
		}
	*/

	/*
		fn fib$aux() {
			OP_LSTORE1 0 // m
			OP_LSTORE1 1 // a
			OP_LSTORE1 2 // b

			OP_LLOAD1 1 // a
			OP_CLOAD1 4 // 0
			OP_EQ

			OP_BR
				OP_CLOAD1 4 // 0
				OP_RET

			OP_LLOAD1 0 // m
			OP_CLOAD1 5 // 1
			OP_ISUB

			OP_LLOAD1 2 // b

			OP_LLOAD1 1 // a
			OP_LLOAD1 2 // b
			OP_IADD

			OP_CLOAD1 7 // fib$aux
			OP_TCALL
			OP_RET
		}

		fn fib() {
			OP_LSTORE1 6 // n

			OP_LLOAD1 6 // n
			OP_CLOAD1 4 // 0
			OP_CLOAD1 5 // 1

			OP_CLOAD1 7 // fib$aux
			OP_SCALL
			OP_RET
		}

		fn main() void {
			OP_CLOAD1 9 // 10
			OP_CLOAD1 8 // fib
			OP_CALL
		}
	*/

	bcm2 := spec.NewBCModule(1, 0, "test", 0, constants.NewConstantPool())

	bcm2.SetFunctions([]*spec.BCFunction{spec.NewBCFunction(bcm2, bcm2.GetConstantPool().GetOrCreateUTF8("loadMe"), spec.BCFunctionModPub, []uint8{
		uint8(opcodes.OP_CLOAD),
		uint8(0),
		uint8(bcm2.GetConstantPool().GetOrCreateUTF8("Hi from \"test\" module!\n")),

		uint8(opcodes.OP_VLOAD),
		uint8(0),
		uint8(bcm2.GetConstantPool().GetOrCreateUTF8("_builtin_print")),
		uint8(opcodes.OP_CALL),

		uint8(opcodes.OP_VLOAD),
		uint8(0),
		uint8(bcm2.GetConstantPool().GetOrCreateUTF8("_builtin_read_line")),
		uint8(opcodes.OP_CALL),

		uint8(opcodes.OP_VLOAD),
		uint8(0),
		uint8(bcm2.GetConstantPool().GetOrCreateUTF8("_builtin_print")),
		uint8(opcodes.OP_CALL),

		uint8(opcodes.OP_CLOAD),
		uint8(0),
		uint8(bcm2.GetConstantPool().GetOrCreateUTF8("\n")),

		uint8(opcodes.OP_VLOAD),
		uint8(0),
		uint8(bcm2.GetConstantPool().GetOrCreateUTF8("_builtin_print")),
		uint8(opcodes.OP_CALL),

		uint8(opcodes.OP_RET),
	})})

	bcm := spec.NewBCModule(1, 0, "bootstrap", 0, constants.NewConstantPool())
	testIdx := uint32(bcm.GetConstantPool().GetOrCreateUTF8("test"))
	loadMeIdx := uint32(bcm.GetConstantPool().GetOrCreateUTF8("loadMe"))
	bcm.SetFunctions([]*spec.BCFunction{spec.NewBCFunction(bcm, bcm.GetConstantPool().GetOrCreateUTF8("main"), spec.BCFunctionModPub, []uint8{
		//uint8(opcodes.OP_CLOAD),
		//uint8(0),
		//uint8(bcm.GetConstantPool().GetOrCreateUTF8("RegisterLunnoNativeFunctions")),
		//uint8(opcodes.OP_CLOAD),
		//uint8(0),
		//uint8(bcm.GetConstantPool().GetOrCreateUTF8("./testplugin/testplugin.so")),
		//uint8(opcodes.OP_VLOAD),
		//uint8(0),
		//uint8(bcm.GetConstantPool().GetOrCreateUTF8("loadPlugin")),
		//uint8(opcodes.OP_CALL),

		uint8(opcodes.OP_CLOAD),
		uint8(0),
		uint8(bcm.GetConstantPool().GetOrCreateUTF8("Hi from \"boostrap\" module!\n")),

		uint8(opcodes.OP_VLOAD),
		uint8(0),
		uint8(bcm.GetConstantPool().GetOrCreateUTF8("_builtin_print")),
		uint8(opcodes.OP_CALL),

		uint8(opcodes.OP_FLOAD),
		uint8(0),
		uint8(testIdx),
		uint8(loadMeIdx),
		uint8(opcodes.OP_CALL),

		uint8(opcodes.OP_RET),
	})})

	writer := bcio.NewWriter(bcm)
	writer.WriteModule()
	out := writer.GetBytes()

	os.Remove(bcm.GetName() + ".fbc")
	f, _ := os.Create(bcm.GetName() + ".fbc")
	f.Write(out)
	f.Close()

	writer2 := bcio.NewWriter(bcm2)
	writer2.WriteModule()
	out2 := writer2.GetBytes()

	os.Remove(bcm2.GetName() + ".fbc")
	f, _ = os.Create(bcm2.GetName() + ".fbc")
	f.Write(out2)
	f.Close()
}
