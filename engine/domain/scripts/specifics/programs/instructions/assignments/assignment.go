package assignments

import (
	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/containers"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/assignments/assignables"
)

type assignment struct {
	hash       hash.Hash
	variables  []string
	assignable assignables.Assignable
	initial    containers.Container
}

func createAssignment(
	hash hash.Hash,
	variables []string,
	assignable assignables.Assignable,
) Assignment {
	return createAssignmentInternally(
		hash,
		variables,
		assignable,
		nil,
	)
}

func createAssignmentWithInitial(
	hash hash.Hash,
	variables []string,
	assignable assignables.Assignable,
	initial containers.Container,
) Assignment {
	return createAssignmentInternally(
		hash,
		variables,
		assignable,
		initial,
	)
}

func createAssignmentInternally(
	hash hash.Hash,
	variables []string,
	assignable assignables.Assignable,
	initial containers.Container,
) Assignment {
	out := assignment{
		hash:       hash,
		variables:  variables,
		assignable: assignable,
		initial:    initial,
	}

	return &out
}

// Hash returns the hash
func (obj *assignment) Hash() hash.Hash {
	return obj.hash
}

// Variables returns the variables
func (obj *assignment) Variables() []string {
	return obj.variables
}

// Assignable returns the assignable
func (obj *assignment) Assignable() assignables.Assignable {
	return obj.assignable
}

// HasInitial returns true if there is an initial, false otherwise
func (obj *assignment) HasInitial() bool {
	return obj.initial != nil
}

// Initial returns the initial, if any
func (obj *assignment) Initial() containers.Container {
	return obj.initial
}
