package primitives

// Primitive represents a primitive kind
type Primitive interface {
	Flag() uint8
	IsNumeric() bool
}
