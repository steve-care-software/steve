package names

import "github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/headers/names/cardinalities"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the name builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithCardinality(cardinality cardinalities.Cardinality) Builder
	Now() (Name, error)
}

// Name represents an header name
type Name interface {
	Name() string
	Cardinality() cardinalities.Cardinality
}
