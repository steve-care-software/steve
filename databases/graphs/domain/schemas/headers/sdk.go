package headers

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the header builder
type Builder interface {
	Create() Builder
	WithVersion(version uint) Builder
	WithName(name string) Builder
	Now() (Header, error)
}

// Header represents the language header
type Header interface {
	Version() uint
	Name() string
}
