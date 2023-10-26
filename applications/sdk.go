package applications

import (
	"github.com/steve-care-software/steve/applications/administrators"
	"github.com/steve-care-software/steve/applications/connections/authenticates"
	"github.com/steve-care-software/steve/applications/identities"
	"github.com/steve-care-software/steve/applications/visitors"
	"github.com/steve-care-software/steve/domain/credentials"
)

// Application represents the application
type Application interface {
	Authenticate(credentials credentials.Credentials) (administrators.Application, error)
	Identify(credentials credentials.Credentials) (identities.Application, error)
	Connect() (authenticates.Application, error)
	Visitor() (visitors.Application, error)
}
