package queues

import (
	"github.com/steve-care-software/steve/vms/children/commands/domain/resources"
)

// Builder represents a queue builder
type Builder interface {
	Create() Builder
	WithPath(path string) Builder
	WithCommands(commands resources.Resources) Builder
	Now() (Queue, error)
}

// Queue represents a queue
type Queue interface {
	Path() string
	Commands() resources.Resources
}

// Repository represents the queue repository
type Repository interface {
	Init(path string) (*uint, error)
	Retrieve(context uint) (Queue, error)
}

// Service represents the queue service
type Service interface {
	Append(context uint, cmd resources.Resource) error
	Replace(context uint, queue Queue) error
	Clear(context uint) error
}
