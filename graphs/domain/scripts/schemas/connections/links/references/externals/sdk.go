package externals

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Externals represents externals
type Externals interface {
	List() []External
}

// Builder represents the external builder
type Builder interface {
	Create() Builder
	WithSchema(schema string) Builder
	WithPoint(point string) Builder
	Now() (External, error)
}

// External represents an external reference
type External interface {
	Schema() string
	Point() string
}
