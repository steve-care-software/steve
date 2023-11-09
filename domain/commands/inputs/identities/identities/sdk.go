package identities

import (
	"github.com/steve-care-software/steve/domain/commands/inputs/identities/identities/authenticates"
	"github.com/steve-care-software/steve/domain/commands/inputs/identities/identities/instances"
	"github.com/steve-care-software/steve/domain/commands/inputs/identities/identities/lists"
)

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithAuthenticate(authenticate authenticates.Authenticate) Builder
	WithList(list lists.List) Builder
	WithInstance(instance instances.Instance) Builder
	Now() (Identity, error)
}

// Identity represents an identity input command
type Identity interface {
	IsAuthenticate() bool
	Authenticate() authenticates.Authenticate
	IsList() bool
	List() lists.List
	IsInstance() bool
	Instance() instances.Instance
}
