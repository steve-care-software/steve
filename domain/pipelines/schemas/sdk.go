package schemas

import (
	"github.com/steve-care-software/steve/domain/graphs/connections"
	"github.com/steve-care-software/steve/domain/graphs/points"
)

// ParserAdapter represents a schema parser adapter
type ParserAdapter interface {
	ToSchema(input []byte) (Schema, error)
}

// Schema represents the schema
type Schema interface {
	Name() string
	Version() uint
	Points() points.Points
	Connections() connections.Connections
}
