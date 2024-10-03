package schemas

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/headers"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the schema adapter
type Adapter interface {
	ToLanguage(input []byte) (Schema, []byte, error)
}

// Builder represents the schema builder
type Builder interface {
	Create() Builder
	WithHeader(header headers.Header) Builder
	WithPoints(points []string) Builder
	WithConnections(connections connections.Connections) Builder
	Now() (Schema, error)
}

// Schema represents the schema
type Schema interface {
	Header() headers.Header
	Points() []string
	Connections() connections.Connections
}
