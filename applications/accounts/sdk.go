package accounts

import (
	"github.com/steve-care-software/steve/applications/accounts/administrators"
	"github.com/steve-care-software/steve/applications/accounts/connections"
	"github.com/steve-care-software/steve/applications/accounts/identities"
	"github.com/steve-care-software/steve/applications/accounts/visitors"
	"github.com/steve-care-software/steve/domain/credentials"
)

// Application represents the application
type Application interface {
	Authenticate(credentials credentials.Credentials) (administrators.Application, error)
	Identify(credentials credentials.Credentials) (identities.Application, error)
	Connect() (connections.Application, error)
	Visitor() (visitors.Application, error)
}
