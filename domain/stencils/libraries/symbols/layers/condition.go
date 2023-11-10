package layers

import "github.com/steve-care-software/steve/domain/blockchains/hash"

type condition struct {
	hash       hash.Hash
	variable   string
	executions Executions
}

func createCondition(
	hash hash.Hash,
	variable string,
	executions Executions,
) Condition {
	out := condition{
		hash:       hash,
		variable:   variable,
		executions: executions,
	}

	return &out
}

// Hash returns the hash
func (obj *condition) Hash() hash.Hash {
	return obj.hash
}

// Variable returns the variable
func (obj *condition) Variable() string {
	return obj.variable
}

// Executions returns the executions
func (obj *condition) Executions() Executions {
	return obj.executions
}
