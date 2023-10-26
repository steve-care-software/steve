package reduces

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a reduce
type Builder interface {
	Create() Builder
	WithVariable(variable string) Builder
	WithLength(length uint8) Builder
	Now() (Reduce, error)
}

// Reduce represents a reduce
type Reduce interface {
	Variable() string
	Length() uint8
}
