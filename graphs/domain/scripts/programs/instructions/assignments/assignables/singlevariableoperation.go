package assignables

type singleVariableOperation struct {
	name     string
	operator uint8
}

func createSingleVariableOperation(
	name string,
	operator uint8,
) SingleVariableOperation {
	return &singleVariableOperation{
		name:     name,
		operator: operator,
	}
}

// Name returns the name of the variable
func (obj *singleVariableOperation) Name() string {
	return obj.name
}

// Operator returns the operator
func (obj *singleVariableOperation) Operator() uint8 {
	return obj.operator
}
