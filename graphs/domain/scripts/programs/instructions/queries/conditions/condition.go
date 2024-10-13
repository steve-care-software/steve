package conditions

type condition struct {
	element Element
	clauses Clauses
}

func createCondition(
	element Element,
) Condition {
	return createConditionInternally(element, nil)
}

func createConditionWithClauses(
	element Element,
	clauses Clauses,
) Condition {
	return createConditionInternally(element, clauses)
}

func createConditionInternally(
	element Element,
	clauses Clauses,
) Condition {
	out := condition{
		element: element,
		clauses: clauses,
	}

	return &out
}

// Element returns the element of the condition
func (obj *condition) Element() Element {
	return obj.element
}

// HasClauses returns true if the condition has clauses
func (obj *condition) HasClauses() bool {
	return obj.clauses != nil
}

// Clauses returns the clauses of the condition
func (obj *condition) Clauses() Clauses {
	return obj.clauses
}
