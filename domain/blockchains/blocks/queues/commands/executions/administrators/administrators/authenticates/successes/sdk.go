package successes

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
)

// Builder represents a success builder
type Builder interface {
	Create() Builder
	WithVariable(variable string) Builder
	WithInstance(instance administrators.Administrator) Builder
	Now() (Successful, error)
}

// Successful represents a successful authentication
type Successful interface {
	Variable() string
	Instance() administrators.Administrator
}
