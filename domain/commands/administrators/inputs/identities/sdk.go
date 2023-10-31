package identities

import (
	"github.com/steve-care-software/steve/domain/commands/administrators/inputs/identities/instances"
)

// Builder represents an identities builder
type Builder interface {
	Create() Builder
	WithInstance(instance instances.Instance) Builder
	Now() (Identities, error)
}

// Identities represents identities
type Identities interface {
	IsInstance() bool
	Instance() instances.Instance
}
