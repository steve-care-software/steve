package administrators

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/administrators/authenticates"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/administrators/instances"
)

// Builder represents the builder
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
