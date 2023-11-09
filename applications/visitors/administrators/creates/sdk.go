package creates

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/visitors/administrators/creates"
)

// Application represents the visitor application
type Application interface {
	Execute(administrator inputs.Create) (administrators.Administrator, error)
}
