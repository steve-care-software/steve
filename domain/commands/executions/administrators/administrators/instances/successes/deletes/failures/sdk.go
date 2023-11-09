package failures

import (
	"github.com/steve-care-software/steve/domain/accounts/credentials"
)

// Builder represents a failure builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials credentials.Credentials) Builder
	CouldNotRetrieve() Builder
	CouldNotDelete() Builder
	Now() (Failure, error)
}

// Failure represents a failure
type Failure interface {
	Credentials() credentials.Credentials
	Content() Content
}

// Content represents the failure content
type Content interface {
	CouldNotRetrieve() bool
	CouldNotDelete() bool
}
