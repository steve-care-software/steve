package initializations

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/containers"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/instructions/assignments"
)

type initialization struct {
	hash       hash.Hash
	container  containers.Container
	assignment assignments.Assignment
}

func createInitialization(
	hash hash.Hash,
	container containers.Container,
	assignment assignments.Assignment,
) Initialization {
	out := initialization{
		hash:       hash,
		container:  container,
		assignment: assignment,
	}

	return &out
}

// Hash returns the hash
func (obj *initialization) Hash() hash.Hash {
	return obj.hash
}

// Container returns the container
func (obj *initialization) Container() containers.Container {
	return obj.container
}

// Assignment returns the assignment
func (obj *initialization) Assignment() assignments.Assignment {
	return obj.assignment
}
