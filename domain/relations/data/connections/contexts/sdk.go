package contexts

// Contexts represents contexts
type Contexts interface {
	List() []Context
	Fetch(name string) (Context, error)
}

// Context represents a context
type Context interface {
	Name() string
	Value() string
}
