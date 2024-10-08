package heads

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/components/compensations"
	"github.com/steve-care-software/steve/engine/domain/scripts/components/roles"
	"github.com/steve-care-software/steve/hash"
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
	WithRole(role roles.Role) Builder
	WithCompensation(compensation compensations.Compensation) Builder
	Now() (Head, error)
}

// Head represents the head
type Head interface {
	Hash() hash.Hash
	Name() string
	Version() uint
	HasRole() bool
	Role() roles.Role
	HasCompensation() bool
	Compensation() compensations.Compensation
}
