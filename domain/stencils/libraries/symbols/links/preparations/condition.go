package preparations

type condition struct {
	variable     string
	preparations Preparations
}

func createCondition(
	variable string,
	preparations Preparations,
) Condition {
	out := condition{
		variable:     variable,
		preparations: preparations,
	}

	return &out
}

// Variable returns the variable
func (obj *condition) Variable() string {
	return obj.variable
}

// Preparations returns the preparations
func (obj *condition) Preparations() Preparations {
	return obj.preparations
}
