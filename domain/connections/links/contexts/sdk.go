package contexts

import "github.com/google/uuid"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewContextBuilder creates a new context builder
func NewContextBuilder() ContextBuilder {
	return createContextBuilder()
}

// Builder represents the contexts builder
type Builder interface {
	Create() Builder
	WithList(list []Context) Builder
	Now() (Contexts, error)
}

// Contexts represents contexts
type Contexts interface {
	List() []Context
}

// ContextBuilder represents the context builder
type ContextBuilder interface {
	Create() ContextBuilder
	WithIdentifier(identifier uuid.UUID) ContextBuilder
	WithName(name string) ContextBuilder
	WithParent(parent uuid.UUID) ContextBuilder
	Now() (Context, error)
}

// Context represents a context
type Context interface {
	Identifier() uuid.UUID
	Name() string
	HasParent() bool
	Parent() *uuid.UUID
}
