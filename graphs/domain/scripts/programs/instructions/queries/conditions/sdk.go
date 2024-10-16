package conditions

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/assignments"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewClausesBuilder creates a new clauses builder
func NewClausesBuilder() ClausesBuilder {
	return createClausesBuilder()
}

// NewClauseBuilder creates a new clause builder
func NewClauseBuilder() ClauseBuilder {
	return createClauseBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// Builder represents a condition builder
type Builder interface {
	Create() Builder
	WithElement(element Element) Builder
	WithClauses(clauses Clauses) Builder
	Now() (Condition, error)
}

// Condition represents a query condition
type Condition interface {
	Element() Element
	HasClauses() bool
	Clauses() Clauses
}

// ClausesBuilder represents the clauses builder
type ClausesBuilder interface {
	Create() ClausesBuilder
	WithList(list []Clause) ClausesBuilder
	Now() (Clauses, error)
}

// Clauses represents query condition clauses
type Clauses interface {
	List() []Clause
}

// ClauseBuilder represents the clause builder
type ClauseBuilder interface {
	Create() ClauseBuilder
	WithOperator(operator uint8) ClauseBuilder
	WithElement(element Element) ClauseBuilder
	Now() (Clause, error)
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
