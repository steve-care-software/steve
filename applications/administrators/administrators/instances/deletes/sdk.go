package deletes

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators/instances/deletes"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators/instances/contents/deletes"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithInstance(instance administrators.Administrator) Builder
	Now() (Application, error)
}

// Application represents the delete application
type Application interface {
	Execute(delete inputs.Delete) (executions.Delete, error)
}
