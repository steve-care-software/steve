package compilers

// Compiler represents the compiler
type Compiler interface {
	// Compiles a script to bytecode
	Compile(input []byte) ([]byte, error)

	// Decompile decompiles bytecode to a script
	Decompile(input []byte) ([]byte, error)
}
