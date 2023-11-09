package commands

import (
	"github.com/steve-care-software/steve/domain/commands/executions"
	"github.com/steve-care-software/steve/domain/commands/inputs"
)

// Builder represents a command builder
type Builder interface {
	Create() Builder
	WithInput(input inputs.Input) Builder
	WithExecution(execution executions.Execution) Builder
	Now() (Command, error)
}

// Command represents a command
type Command interface {
	Input() inputs.Input
	Execution() executions.Execution
}
