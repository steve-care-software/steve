package chains

type interpreter struct {
	variable string
	next     Chain
}

func createInterpreter(
	variable string,
) Interpreter {
	return createInterpreterInternally(variable, nil)
}

func createInterpreterWithNext(
	variable string,
	next Chain,
) Interpreter {
	return createInterpreterInternally(variable, next)
}

func createInterpreterInternally(
	variable string,
	next Chain,
) Interpreter {
	out := interpreter{
		variable: variable,
		next:     next,
	}

	return &out
}

// Variable returns the variable
func (obj *interpreter) Variable() string {
	return obj.variable
}

// HasNext returns true if next, false otherwise
func (obj *interpreter) HasNext() bool {
	return obj.next != nil
}

// Next returns the next, if any
func (obj *interpreter) Next() Chain {
	return obj.next
}
