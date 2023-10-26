package administrators

import "github.com/steve-care-software/steve/domain/accounts/administrators/identities"

// Administrator represents an administrator
type Administrator interface {
	HasIdentities() bool
	Identities() identities.Identities
}
