package identities

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/identities/identities/authenticates"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/identities/identities/instances"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/identities/identities/lists"
)

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithAuthenticate(authenticate authenticates.Authenticate) Builder
	WithInstance(instance instances.Instance) Builder
	WithList(list lists.List) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	IsAuthenticate() bool
	Authenticate() authenticates.Authenticate
	IsInstance() bool
	Instance() instances.Instance
	IsList() bool
	List() lists.List
}
