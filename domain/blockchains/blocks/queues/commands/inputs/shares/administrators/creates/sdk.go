package creates

import (
	"github.com/steve-care-software/steve/domain/accounts/credentials"
	"github.com/steve-care-software/steve/domain/dashboards"
)

// Create represents a create account
type Create interface {
	Credentials() credentials.Credentials
	Dashboard() dashboards.Dashboard
}
