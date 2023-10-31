package authenticates

import "github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators/authenticates/credentials"

// Builder represents an authenticate builder
type Builder interface {
	Create() Builder
	WithAssignToVariable(assignToVariable string) Builder
	WithUsername(username string) Builder
	WithPassword(password []byte) Builder
	Now() (Authenticate, error)
}

// Authenticate represents an authenticate
type Authenticate interface {
	AssignToVariable() string
	Credentials() credentials.Credentials
}
