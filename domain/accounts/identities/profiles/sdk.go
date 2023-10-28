package profiles

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/profiles/authorizations"
	"github.com/steve-care-software/steve/domain/accounts/identities/profiles/roles"
	"github.com/steve-care-software/steve/domain/stencils"
)

// Profile represents a profile
type Profile interface {
	Name() string
	Description() string
	HasConnections() bool
	Connections() Connections
	HasParent() bool
	Parent() Profile
}

// Connections represents connections
type Connections interface {
	List() []Connection
}

// Connection represents a connection
type Connection interface {
	Rank() uint
	Profile() Profile
	Authorization() authorizations.Authorization
	Root() stencils.Stencil
	IsPublic() bool
	HasRole() bool
	Role() roles.Role
}
