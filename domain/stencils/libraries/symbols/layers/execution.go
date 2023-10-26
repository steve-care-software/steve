package layers

type execution struct {
	isStop     bool
	assignment Assignment
	condition  Condition
}

func createExecutionWithStop() Execution {
	return createExecutionInternally(true, nil, nil)
}

func createExecutionWithAssignment(
	assignment Assignment,
) Execution {
	return createExecutionInternally(false, assignment, nil)
}

func createExecutionWithCondition(
	condition Condition,
) Execution {
	return createExecutionInternally(false, nil, condition)
}

func createExecutionInternally(
	isStop bool,
	assignment Assignment,
	condition Condition,
) Execution {
	out := execution{
		isStop:     isStop,
		assignment: assignment,
		condition:  condition,
	}

	return &out
}

// IsStop returns true if there is a stop, false otherwise
func (obj *execution) IsStop() bool {
	return obj.isStop
}

// IsAssignment returns true if there is an assignment, false otherwise
func (obj *execution) IsAssignment() bool {
	return obj.assignment != nil
}

// Assignment returns the assignment, if any
func (obj *execution) Assignment() Assignment {
	return obj.assignment
}

// IsCondition returns true if there is a condition, false otherwise
func (obj *execution) IsCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *execution) Condition() Condition {
	return obj.condition
}
