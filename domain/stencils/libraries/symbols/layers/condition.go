package layers

type condition struct {
	variable   string
	executions Executions
}

func createCondition(
	variable string,
	executions Executions,
) Condition {
	out := condition{
		variable:   variable,
		executions: executions,
	}

	return &out
}

// Variable returns the variable
func (obj *condition) Variable() string {
	return obj.variable
}

// Executions returns the executions
func (obj *condition) Executions() Executions {
	return obj.executions
}
