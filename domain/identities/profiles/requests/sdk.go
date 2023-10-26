package requests

import "github.com/steve-care-software/steve/domain/accounts/identities/profiles/authorizations"

// Request represents an authorization request
type Request interface {
	Name() string
	Authorization() authorizations.Authorization
}
