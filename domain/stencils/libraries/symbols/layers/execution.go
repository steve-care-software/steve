package layers

import "github.com/steve-care-software/steve/domain/hash"

type execution struct {
	hash       hash.Hash
	isStop     bool
	assignment Assignment
	condition  Condition
}

func createExecutionWithStop(
	hash hash.Hash,
) Execution {
	return createExecutionInternally(hash, true, nil, nil)
}

func createExecutionWithAssignment(
	hash hash.Hash,
	assignment Assignment,
) Execution {
	return createExecutionInternally(hash, false, assignment, nil)
}

func createExecutionWithCondition(
	hash hash.Hash,
	condition Condition,
) Execution {
	return createExecutionInternally(hash, false, nil, condition)
}

func createExecutionInternally(
	hash hash.Hash,
	isStop bool,
	assignment Assignment,
	condition Condition,
) Execution {
	out := execution{
		hash:       hash,
		isStop:     isStop,
		assignment: assignment,
		condition:  condition,
	}

	return &out
}

// Hash returns the hash
func (obj *execution) Hash() hash.Hash {
	return obj.hash
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
