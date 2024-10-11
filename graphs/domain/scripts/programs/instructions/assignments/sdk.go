package assignments

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/containers/kinds"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/assignments/assignables"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references"
)

// Assignment represents an assignment
type Assignment interface {
	IsMultiple() bool
	Multiple() AssignmentMultiple
	IsOperation() bool
	Operation() AssignmentOperation
}

// AssignmentMultiple represents a multiple assignment
type AssignmentMultiple interface {
	Assignees() Assignees
	Assignables() assignables.Assignables
}

// AssignmentOperation represents an assignment operation (ex: myVAr++)
type AssignmentOperation interface {
	Assignee() Assignee
	Operator() uint8                    // arithmetic operator
	Assignable() assignables.Assignable // must compute to an arithmetic-valid value
}

// Assignees represents assignees
type Assignees interface {
	List() []Assignee
}

// Assignee represents an assignee
type Assignee interface {
	Name() AssigneeName
	HasKind() bool
	Kind() kinds.Kind
}

// AssigneeName represents an assignee name
type AssigneeName interface {
	Name() string
	HasReferences() bool
	References() references.References
	HasIndex() bool
	Index() *uint
}
