package administrators

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	"github.com/steve-care-software/steve/domain/accounts/administrators/identities"
)

// Administrator represents an administrator
type Administrator interface {
	IsInstance() bool
	Instance() administrators.Administrator
	IsIdentities() bool
	Identities() identities.Identities
}
