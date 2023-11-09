package successes

import "github.com/steve-care-software/steve/domain/accounts/identities"

// Builder represents a success builder
type Builder interface {
	Create() Builder
	WithVariable(variable string) Builder
	WithInstance(instance identities.Identity) Builder
	Now() (Successful, error)
}

// Successful represents a successful authentication
type Successful interface {
	Variable() string
	Instance() identities.Identity
}
