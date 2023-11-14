package programs

import (
	bytes_programs "github.com/steve-care-software/steve/vms/bytes/programs"
	"github.com/steve-care-software/steve/vms/bytes/results/hash"
	assignment_programs "github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/programs"
)

// Builder represents a programs builder
type Builder interface {
	Create() Builder
	WithList(list []Program) Builder
	Now() (Programs, error)
}

// Programs represents programs
type Programs interface {
	Hash() hash.Hash
	List() []Program
}

// ProgramBuilder represents the program builder
type ProgramBuilder interface {
	Create() ProgramBuilder
	WithAssignment(assignment assignment_programs.Program) ProgramBuilder
	WithCondition(condition Condition) ProgramBuilder
	IsStop() ProgramBuilder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Hash() hash.Hash
	IsStop() bool
	IsAssignment() bool
	Assignment() assignment_programs.Program
	IsCondition() bool
	Condition()
}

// ConditionBuilder represents a condition builder
type ConditionBuilder interface {
	Create() ConditionBuilder
	WithConstraint(constraint bytes_programs.Program) ConditionBuilder
	WithPrograms(programs Programs) ConditionBuilder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Hash() hash.Hash
	Constraint() bytes_programs.Program
	Programs() Programs
}
