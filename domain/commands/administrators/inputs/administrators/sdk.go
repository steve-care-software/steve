package administrators

import (
	"github.com/steve-care-software/steve/domain/commands/administrators/inputs/administrators/authenticates"
	"github.com/steve-care-software/steve/domain/commands/administrators/inputs/administrators/instances"
)

// Builder represents an administrator's builder
type Builder interface {
	Create() Builder
	WithAuthenticate(authenticate authenticates.Authenticate) Builder
	WithInstance(instance instances.Instance) Builder
	Now() (Administrator, error)
}

// Administrator represents an administrator
type Administrator interface {
	IsAuthenticate() bool
	Authenticate() authenticates.Authenticate
	IsInstance() bool
	Instance() instances.Instance
}
