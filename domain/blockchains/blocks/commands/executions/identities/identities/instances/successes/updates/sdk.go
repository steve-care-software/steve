package updates

import "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/identities/identities/instances/successes/updates/passwords"

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
