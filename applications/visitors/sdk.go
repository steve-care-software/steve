package visitors

import "github.com/steve-care-software/steve/domain/commands/visitors/executions"

// Application represents the visitor application
type Application interface {
	Execute(message []byte) (executions.Execution, error)
}
