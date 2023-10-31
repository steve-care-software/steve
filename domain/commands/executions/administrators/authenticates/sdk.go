package authenticates

import (
	"github.com/steve-care-software/steve/domain/commands/executions/administrators/authenticates/successes"
	"github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators/authenticates/credentials"
)

// Builder represents an authenticate builder
type Builder interface {
	Create() Builder
	WithSuccess(success successes.Successful) Builder
	WithFailure(failure credentials.Credentials) Builder
	Now() (Authenticate, error)
}

// Authenticate represents an authenticate
type Authenticate interface {
	IsSuccess() bool
	Success() successes.Successful
	IsFailure() bool
	Failure() credentials.Credentials
}
