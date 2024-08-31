package contexts

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
	WithName(name string) ContextBuilder
	WithParent(parent Context) ContextBuilder
	Now() (Context, error)
}

// Context represents a context
type Context interface {
	Name() string
	HasParent() bool
	Parent() Context
}
