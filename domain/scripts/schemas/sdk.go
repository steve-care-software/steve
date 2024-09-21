package schemas

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/schemas/compensations"
	"github.com/steve-care-software/steve/domain/scripts/schemas/connections"
	"github.com/steve-care-software/steve/domain/scripts/schemas/points"
	"github.com/steve-care-software/steve/domain/scripts/schemas/roles"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithVersion(version uint) Builder
	WithPoints(points points.Points) Builder
	WithConnections(connections connections.Connections) Builder
	WithRole(role roles.Role) Builder
	WithCompensation(compensation compensations.Compensation) Builder
	Now() (Schema, error)
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
	HasCompensation() bool
	Compensation() compensations.Compensation
}
