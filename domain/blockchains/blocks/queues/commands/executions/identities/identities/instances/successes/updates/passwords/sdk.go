package passwords

import "github.com/steve-care-software/steve/domain/accounts/credentials"

// Builder represents a password builder
type Builder interface {
	Create() Builder
	WithSuccess(success credentials.Credentials) Builder
	WithFailure(failure credentials.Credentials) Builder
	Now() (Password, error)
}

// Password represents a password
type Password interface {
	IsSuccess() bool
	Success() credentials.Credentials
	IsFailure() bool
	Failure() credentials.Credentials
}
