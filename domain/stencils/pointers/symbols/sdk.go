package symbols

const (
	// KindBytes represents the  bytes kind
	KindBytes (uint8) = 1

	// KindLayer represents the layer kind
	KindLayer

	// KindLink represents the link kind
	KindLink
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a symbol builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithKind(kind uint8) Builder
	Now() (Symbol, error)
}

// Symbol represents a symbol
type Symbol interface {
	Name() string
	Kind() uint8
}
