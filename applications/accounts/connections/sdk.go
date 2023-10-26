package connections

import "github.com/steve-care-software/steve/domain/accounts/identities/profiles"

// Application represents the account application
type Application interface {
	Retrieve() (profiles.Profile, error)
	Delete() error
}
