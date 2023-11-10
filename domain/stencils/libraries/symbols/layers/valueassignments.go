package layers

import "github.com/steve-care-software/steve/domain/blockchains/hash"

type valueAssignments struct {
	hash hash.Hash
	list []ValueAssignment
}

func createValueAssignments(
	hash hash.Hash,
	list []ValueAssignment,
) ValueAssignments {
	out := valueAssignments{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *valueAssignments) Hash() hash.Hash {
	return obj.hash
}

// List returns the valueAssignments
func (obj *valueAssignments) List() []ValueAssignment {
	return obj.list
}
