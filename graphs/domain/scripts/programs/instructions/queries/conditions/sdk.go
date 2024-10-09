package conditions

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/assignments"
)

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

// Element represents a query condition element
type Element interface {
	IsAssignment() bool
	Assignment() assignments.Assignment
	IsCondition() bool
	Condition() Condition
}
