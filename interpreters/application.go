package interpreters

type interpreter struct {
	instructions []byte
	params       map[string][]byte
}

func createInterpreter(
	instructions []byte,
	params map[string][]byte,
) Interpreter {
	out := interpreter{
		instructions: instructions,
		params:       params,
	}

	return &out
}

// Execute executes the interpreter
func (app *interpreter) Execute() ([]byte, error) {
	return nil, nil
}
