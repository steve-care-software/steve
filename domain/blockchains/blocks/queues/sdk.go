package queues

import "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands"

// Builder represents a queue builder
type Builder interface {
	Create() Builder
	WithPath(path string) Builder
	WithCommands(commands commands.Commands) Builder
	Now() (Queue, error)
}

// Queue represents a queue
type Queue interface {
	Path() string
	Commands() commands.Commands
}

// Repository represents the queue repository
type Repository interface {
	Init(path string) (*uint, error)
	Retrieve(context uint) (Queue, error)
}

// Service represents the queue service
type Service interface {
	Append(context uint, cmd commands.Command) error
	Replace(context uint, queue Queue) error
	Clear(context uint) error
}
