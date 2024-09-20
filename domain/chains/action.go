package chains

type action struct {
	interpreter Interpreter
	transpile   Transpile
}

func createActionWithInterpreter(
	interpreter Interpreter,
) Action {
	return createActionInternally(interpreter, nil)
}

func createActionWithTranspile(
	transpile Transpile,
) Action {
	return createActionInternally(nil, transpile)
}

func createActionInternally(
	interpreter Interpreter,
	transpile Transpile,
) Action {
	out := action{
		interpreter: interpreter,
		transpile:   transpile,
	}

	return &out
}

// IsInterpret returns true if there is an interpreter, false otherwise
func (obj *action) IsInterpret() bool {
	return obj.interpreter != nil
}

// Interpret returns the interpreter, if any
func (obj *action) Interpret() Interpreter {
	return obj.interpreter
}

// IsTranspile returns true if there is a transpile, false otherwise
func (obj *action) IsTranspile() bool {
	return obj.transpile != nil
}

// Transpile returns the transpile, if any
func (obj *action) Transpile() Transpile {
	return obj.transpile
}
