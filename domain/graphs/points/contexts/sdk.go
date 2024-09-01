package contexts

import (
	"github.com/steve-care-software/steve/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewContextBuilder creates a new context builder
func NewContextBuilder() ContextBuilder {
	hashAdapter := hash.NewAdapter()
	return createContextBuilder(
		hashAdapter,
	)
}

// Builder represents the contexts builder
type Builder interface {
	Create() Builder
	WithList(list []Context) Builder
	Now() (Contexts, error)
}

// Contexts represents contexts
type Contexts interface {
	Hash() hash.Hash
	List() []Context
}

// ContextBuilder represents the context builder
type ContextBuilder interface {
	Create() ContextBuilder
	WithName(name string) ContextBuilder
	WithParent(parent hash.Hash) ContextBuilder
	Now() (Context, error)
}

// Context represents a context
type Context interface {
	Hash() hash.Hash
	Name() string
	HasParent() bool
	Parent() hash.Hash
}
