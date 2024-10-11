package assignables

type operatorAssignable struct {
	operator   Operator
	assignable Assignable
}

func createOperatorAssignable(
	operator Operator,
	assignable Assignable,
) OperatorAssignable {
	return &operatorAssignable{
		operator:   operator,
		assignable: assignable,
	}
}

// Operator returns the operator
func (obj *operatorAssignable) Operator() Operator {
	return obj.operator
}

// Assignable returns the assignable
func (obj *operatorAssignable) Assignable() Assignable {
	return obj.assignable
}
