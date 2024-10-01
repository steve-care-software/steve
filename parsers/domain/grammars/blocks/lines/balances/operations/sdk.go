package operations

import (
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/operations/selectors"
)

const (
	// OperatorAnd represents the && operator
	OperatorAnd (uint8) = iota

	// OperatorOr represents the || operator
	OperatorOr

	// OperatorXor represents the <> operator
	OperatorXor
)

// NewActorBuilder creates a new actor builder
func NewActorBuilder() ActorBuilder {
	return createActorBuilder()
}

// NewTailBuilder creates a new tail builder
func NewTailBuilder() TailBuilder {
	return createTailBuilder()
}

// Builder represents the operation builder
type Builder interface {
	Create() Builder
	WithList(list []Operation) Builder
	Now() (Operations, error)
}

// Operations represents operations
type Operations interface {
	List() []Operation
}

// OperationBuilder represents an operation builder
type OperationBuilder interface {
	Create() OperationBuilder
	WithActor(actor Actor) OperationBuilder
	WithTail(tail Tail) OperationBuilder
	IsNot() OperationBuilder
	Now() (Operation, error)
}

// Operation represents the operation
type Operation interface {
	Actor() Actor
	Tail() Tail
	IsNot() bool
}

// ActorBuilder represents an actor builder
type ActorBuilder interface {
	Create() ActorBuilder
	WithSelector(selector selectors.Selector) ActorBuilder
	WithOperation(operation Operation) ActorBuilder
	Now() (Actor, error)
}

// Actor represents an operation actor
type Actor interface {
	IsSelector() bool
	Selector() selectors.Selector
	IsOperation() bool
	Operation() Operation
}

// TailBuilder represents a tail builder
type TailBuilder interface {
	Create() TailBuilder
	WithOperator(operator uint8) TailBuilder
	WithActor(actor Actor) TailBuilder
	Now() (Tail, error)
}

// Tail represents the operation tail
type Tail interface {
	Operator() uint8
	Actor() Actor
}
