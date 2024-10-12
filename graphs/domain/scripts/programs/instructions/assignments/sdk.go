package assignments

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/containers/kinds"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/assignments/assignables"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewAssignmentMultipleBuilder creates a new assignment multiple builder
func NewAssignmentMultipleBuilder() AssignmentMultipleBuilder {
	return createAssignmentMultipleBuilder()
}

// NewAssignmentOperatorBuilder creates a new assignment operator builder
func NewAssignmentOperatorBuilder() AssignmentOperationBuilder {
	return createAssignmentOperationBuilder()
}

// NewAssigneeBuilder creates a new assignee builderg
func NewAssigneeBuilder() AssigneeBuilder {
	return createAssigneeBuilder()
}

// NewAssigneeNameBuilder creates a new assignee name builder
func NewAssigneeNameBuilder() AssigneeNameBuilder {
	return createAssigneeNameBuilder()
}

// Builder represents the assignment builder
type Builder interface {
	Create() Builder
	WithMultiple(multiple AssignmentMultiple) Builder
	WithOperation(operation AssignmentOperation) Builder
	Now() (Assignment, error)
}

// Assignment represents an assignment
type Assignment interface {
	IsMultiple() bool
	Multiple() AssignmentMultiple
	IsOperation() bool
	Operation() AssignmentOperation
}

// AssignmentMultipleBuilder represents the assignment multiple builder
type AssignmentMultipleBuilder interface {
	Create() AssignmentMultipleBuilder
	WithAssignees(assignees Assignees) AssignmentMultipleBuilder
	WithAssignables(assignables assignables.Assignables) AssignmentMultipleBuilder
	Now() (AssignmentMultiple, error)
}

// AssignmentMultiple represents a multiple assignment
type AssignmentMultiple interface {
	Assignees() Assignees
	Assignables() assignables.Assignables
}

// AssignmentOperationBuilder represents the assignment operation builder
type AssignmentOperationBuilder interface {
	Create() AssignmentOperationBuilder
	WithAssignee(assignee Assignee) AssignmentOperationBuilder
	WithOperator(operator uint8) AssignmentOperationBuilder
	WithAssignable(assignable assignables.Assignable) AssignmentOperationBuilder
	Now() (AssignmentOperation, error)
}

// AssignmentOperation represents an assignment operation (ex: myVar++)
type AssignmentOperation interface {
	Assignee() Assignee
	Operator() uint8                    // arithmetic operator
	Assignable() assignables.Assignable // must compute to an arithmetic-valid value
}

// Assignees represents assignees
type Assignees interface {
	List() []Assignee
}

// AssigneeBuilder represents the assignee builder
type AssigneeBuilder interface {
	Create() AssigneeBuilder
	WithName(name AssigneeName) AssigneeBuilder
	WithKind(kind kinds.Kind) AssigneeBuilder
	Now() (Assignee, error)
}

// Assignee represents an assignee
type Assignee interface {
	Name() AssigneeName
	HasKind() bool
	Kind() kinds.Kind
}

// AssigneeNameBuilder represents the assigneename builder
type AssigneeNameBuilder interface {
	Create() AssigneeNameBuilder
	WithReferences(references references.References) AssigneeNameBuilder
	WithIndex(index uint) AssigneeNameBuilder
	Now() (AssigneeName, error)
}

// AssigneeName represents an assignee name
type AssigneeName interface {
	References() references.References
	HasIndex() bool
	Index() *uint
}
