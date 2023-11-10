package commands

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/frames"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs"
)

// Builder represents a commands builder
type Builder interface {
	Create() Builder
	WithList(list []Command) Builder
	Now() (Commands, error)
}

// Commands represents commands
type Commands interface {
	List() []Command
	Frames() frames.Frames
}

// CommandBuilder represents a command builder
type CommandBuilder interface {
	Create() CommandBuilder
	WithInput(input inputs.Input) CommandBuilder
	WithExecution(execution executions.Execution) CommandBuilder
	WithPreviousFrame(previousFrame frames.Frame) CommandBuilder
	Now() (Command, error)
}

// Command represents a command
type Command interface {
	Input() inputs.Input
	Execution() executions.Execution
	Frame() frames.Frame
}
