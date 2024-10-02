package cardinalities

// Cardinality represents a cardinality
type Cardinality interface {
	Min() uint
	HaxMax() bool
	Max() *uint
}
