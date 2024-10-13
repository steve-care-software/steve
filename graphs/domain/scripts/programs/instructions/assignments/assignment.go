package assignments

type assignment struct {
	multiple  AssignmentMultiple
	operation AssignmentOperation
}

func createAssignmentWithMultiple(multiple AssignmentMultiple) Assignment {
	return createAssignmentInternally(multiple, nil)
}

func createAssignmentWithOperation(operation AssignmentOperation) Assignment {
	return createAssignmentInternally(nil, operation)
}

func createAssignmentInternally(
	multiple AssignmentMultiple,
	operation AssignmentOperation,
) Assignment {
	out := assignment{
		multiple:  multiple,
		operation: operation,
	}

	return &out
}

// IsMultiple checks if the assignment is multiple
func (obj *assignment) IsMultiple() bool {
	return obj.multiple != nil
}

// Multiple returns the multiple assignment if present
func (obj *assignment) Multiple() AssignmentMultiple {
	return obj.multiple
}

// IsOperation checks if the assignment is an operation
func (obj *assignment) IsOperation() bool {
	return obj.operation != nil
}

// Operation returns the assignment operation if present
func (obj *assignment) Operation() AssignmentOperation {
	return obj.operation
}
