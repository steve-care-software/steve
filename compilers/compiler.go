package compilers

type compiler struct {
}

func createCompiler() Compiler {
	out := compiler{}
	return &out
}

// Compiles a script to bytecode
func (app *compiler) Compile(input []byte) ([]byte, error) {
	return nil, nil
}

// Decompile decompiles bytecode to a script
func (app *compiler) Decompile(input []byte) ([]byte, error) {
	return nil, nil
}
