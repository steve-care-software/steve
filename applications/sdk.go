package applications

import (
	"github.com/steve-care-software/steve/applications/accounts"
)

// Application represents the application
type Application interface {
	Account() accounts.Application
}
