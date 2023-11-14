package programs

import (
	"github.com/steve-care-software/steve/vms/bytes/results/hash"
	assignable_programs "github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/scopes/assignables/programs"
)

// Builder represents a program builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithAssignable(assignable assignable_programs.Program) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Hash() hash.Hash
	Name() string
	Assignable() assignable_programs.Program
}
