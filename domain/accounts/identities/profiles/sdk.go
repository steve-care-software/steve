package profiles

import "github.com/steve-care-software/steve/domain/accounts/identities/profiles/authorizations"

// Profile represents a profile
type Profile interface {
	Name() string
	Description() string
	HasConnections() bool
	Connections() Connections
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
	IsPublic() bool
}
