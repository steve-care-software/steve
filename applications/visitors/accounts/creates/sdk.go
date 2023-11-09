package creates

import (
	"github.com/steve-care-software/steve/domain/accounts/visitors"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/visitors/accounts/creates"
)

// Application represents the application
type Application interface {
	Execute(account inputs.Create) (visitors.Visitor, error)
}
