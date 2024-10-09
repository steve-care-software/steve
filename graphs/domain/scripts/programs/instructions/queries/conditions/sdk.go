package conditions

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/assignments"
)

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// Condition represents a query condition
type Condition interface {
	Element() Element
	HasClauses() bool
	Clauses() Clauses
}

// Clauses represents query condition clauses
type Clauses interface {
	List() []Clause
}

// Clause represents a query condition clause
type Clause interface {
	Operator() uint8 // logical operator
	Element() Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithAssignment(assignment assignments.Assignment) ElementBuilder
	WithCondition(condition Condition) ElementBuilder
	Now() (Element, error)
}

// Element represents a query condition element
type Element interface {
	IsAssignment() bool
	Assignment() assignments.Assignment
	IsCondition() bool
	Condition() Condition
}
