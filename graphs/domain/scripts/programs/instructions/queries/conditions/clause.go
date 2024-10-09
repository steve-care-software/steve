package conditions

// clause represents the implementation of the Clause interface
type clause struct {
	operator uint8
	element  Element
}

func createClause(operator uint8, element Element) Clause {
	return &clause{
		operator: operator,
		element:  element,
	}
}

// Operator returns the logical operator of the clause
func (obj *clause) Operator() uint8 {
	return obj.operator
}

// Element returns the element of the clause
func (obj *clause) Element() Element {
	return obj.element
}
