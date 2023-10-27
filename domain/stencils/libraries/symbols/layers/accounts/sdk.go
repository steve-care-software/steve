package accounts

import (
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/accounts/administrators"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/accounts/credentials"
)

// Account represents an account
type Account interface {
	IsAuthenticate() bool
	Authenticate() credentials.Credentials
	IsAdministrator() bool
	Administrator() administrators.Administrator
}
