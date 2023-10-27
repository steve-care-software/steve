package administrators

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators/identities"
	"github.com/steve-care-software/steve/domain/dashboards"
)

// Administrator represents an administrator
type Administrator interface {
	Dashboard() dashboards.Dashboard
	HasIdentities() bool
	Identities() identities.Identities
}
