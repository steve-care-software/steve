package primitives

const (
	// FlagInt represents an int
	FlagInt (uint8) = iota

	// FlagUint represents an uint
	FlagUint

	// FlagFloat represents a float
	FlagFloat

	// FlagBool represents a bool
	FlagBool

	// FlagString represents a string
	FlagString
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a primitive builder
type Builder interface {
	Create() Builder
	WithFlag(flag uint8) Builder
	Now() (Primitive, error)
}

// Primitive represents a primitive kind
type Primitive interface {
	Flag() uint8
	IsNumeric() bool
}
