package instructions

import "errors"

type operatorAssignableBuilder struct {
	operator   Operator
	assignable Assignable
}

func createOperatorAssignableBuilder() OperatorAssignableBuilder {
	return &operatorAssignableBuilder{}
}

// Create initializes the operator assignable builder
func (obj *operatorAssignableBuilder) Create() OperatorAssignableBuilder {
	return createOperatorAssignableBuilder()
}

// WithOperator adds an operator to the builder
func (obj *operatorAssignableBuilder) WithOperator(operator Operator) OperatorAssignableBuilder {
	obj.operator = operator
	return obj
}

// WithAssignable adds an assignable to the builder
func (obj *operatorAssignableBuilder) WithAssignable(assignable Assignable) OperatorAssignableBuilder {
	obj.assignable = assignable
	return obj
}

// Now builds a new OperatorAssignable instance
func (obj *operatorAssignableBuilder) Now() (OperatorAssignable, error) {
	if obj.operator == nil {
		return nil, errors.New("the operator is mandatory to build an OperatorAssignable instance")
	}
	if obj.assignable == nil {
		return nil, errors.New("the assignable is mandatory to build an OperatorAssignable instance")
	}

	return createOperatorAssignable(obj.operator, obj.assignable), nil
}
