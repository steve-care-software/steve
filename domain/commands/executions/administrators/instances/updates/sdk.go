package updates

import "github.com/steve-care-software/steve/domain/commands/executions/administrators/instances/updates/passwords"

// Builder represents an update builder
type Builder interface {
	Create() Builder
	WithPassword(password passwords.Password) Builder
	Now() (Update, error)
}

// Update represents an update
type Update interface {
	IsPassword() bool
	Password() passwords.Password
}
