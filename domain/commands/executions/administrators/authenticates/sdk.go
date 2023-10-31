package authenticates

import "github.com/steve-care-software/steve/domain/accounts/administrators"

// Builder represents an authenticate builder
type Builder interface {
	Create() Builder
	WithVariable(variable string) Builder
	WithInstance(instance administrators.Administrator) Builder
	Now() (Authenticate, error)
}

// Authenticate represents an authenticate
type Authenticate interface {
	Variable() string
	Instance() administrators.Administrator
}
