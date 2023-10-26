package visitors

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/profiles"
	"github.com/steve-care-software/steve/domain/accounts/identities/profiles/authorizations/requests"
	"github.com/steve-care-software/steve/domain/accounts/identities/profiles/authorizations/responses"
)

// Application represents the visitor's identity application
type Application interface {
	List() []string
	Retrieve(name string) (profiles.Profile, error)
	Initialize() ([]byte, error)
	Authorize(request requests.Request) (responses.Response, error)
}
