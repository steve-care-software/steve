package preparations

import "github.com/steve-care-software/steve/domain/hash"

type condition struct {
	hash         hash.Hash
	variable     string
	preparations Preparations
}

func createCondition(
	hash hash.Hash,
	variable string,
	preparations Preparations,
) Condition {
	out := condition{
		hash:         hash,
		variable:     variable,
		preparations: preparations,
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

// Preparations returns the preparations
func (obj *condition) Preparations() Preparations {
	return obj.preparations
}
