package instructions

type operation struct {
	first       Assignable
	assignables OperatorAssignables
}

func createOperation(
	first Assignable,
) Operation {
	return createOperationInternally(first, nil)
}

func createOperationWithAssignables(
	first Assignable,
	assignables OperatorAssignables,
) Operation {
	return createOperationInternally(first, assignables)
}

func createOperationInternally(
	first Assignable,
	assignables OperatorAssignables,
) Operation {
	return &operation{
		first:       first,
		assignables: assignables,
	}
}

// First returns the first assignable
func (obj *operation) First() Assignable {
	return obj.first
}

// HasAssignables returns true if there are operator assignables
func (obj *operation) HasAssignables() bool {
	return obj.assignables != nil
}

// Assignables returns the operator assignables if they exist
func (obj *operation) Assignables() OperatorAssignables {
	return obj.assignables
}
