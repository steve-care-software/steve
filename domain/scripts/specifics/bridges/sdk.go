package bridges

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/components/compensations"
	"github.com/steve-care-software/steve/domain/scripts/components/roles"
	"github.com/steve-care-software/steve/domain/scripts/specifics/bridges/connections"
)

// Builder represents the bridge builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithVersion(version uint) Builder
	WithOrigin(origin string) Builder
	WithTarget(target string) Builder
	WithConnections(connections connections.Connections) Builder
	WithRole(role roles.Role) Builder
	WithCompensation(compensation compensations.Compensation) Builder
	Now() (Bridge, error)
}

// Bridge represents a bridge
type Bridge interface {
	Hash() hash.Hash
	Name() string
	Version() uint
	Origin() string
	Target() string
	Connections() connections.Connections
	HasRole() bool
	Role() roles.Role
	HasCompensation() bool
	Compensation() compensations.Compensation
}
