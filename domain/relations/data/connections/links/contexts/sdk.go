package contexts

// Builder represents the contexts builder
type Builder interface {
	Create() Builder
	WithList(list []Context) Builder
	Now() (Contexts, error)
}

// Contexts represents contexts
type Contexts interface {
	List() []Context
	Fetch(name string) (Context, error)
}

// ContextBuilder represents the context builder
type ContextBuilder interface {
	Create() ContextBuilder
	WithName(name string) ContextBuilder
	WithValue(value string) ContextBuilder
	Now() (Context, error)
}

// Context represents a context
type Context interface {
	Name() string
	Value() string
}
