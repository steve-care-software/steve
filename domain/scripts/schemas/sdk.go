package schemas

import (
	"github.com/steve-care-software/steve/domain/graphs/connections"
	"github.com/steve-care-software/steve/domain/graphs/points"
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/schemas/roles"
)

// ParserAdapter represents a schema parser adapter
type ParserAdapter interface {
	ToSchema(input []byte) (Schema, error)
}

// Schema represents the schema
type Schema interface {
	Hash() hash.Hash
	Name() string
	Version() uint
	Points() points.Points
	Connections() connections.Connections
	HasRole() bool
	Role() roles.Role
}
