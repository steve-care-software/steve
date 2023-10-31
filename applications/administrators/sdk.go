package administrators

import (
	"github.com/steve-care-software/steve/domain/commands/administrators/executions"
)

// Application represents the administrator's application
type Application interface {
	Execute(message []byte, username string, password []byte) (executions.Execution, error)
}
