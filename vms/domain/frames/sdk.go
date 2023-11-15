package frames

import "github.com/steve-care-software/steve/vms/domain/resources/blocks/queues"

// Factory represents a frame factory
type Factory interface {
	Create() Frame
}

// Builder represents a frame builder
type Builder interface {
	Create() Builder
	WithAssignables(assignables Assignables) Builder
	Now() (Assignables, error)
}

// Frame represents a frame
type Frame interface {
	Save(name string, assignable Assignable) error
	Fetch(name string) (Assignable, error)
}

// AssignablesBuilder represents the assignables builder
type AssignablesBuilder interface {
	Create() AssignablesBuilder
	WithList(list []Assignable) AssignablesBuilder
	Now() (Assignables, error)
}

// Assignables represents assignables
type Assignables interface {
	List() []Assignable
}

// AssignableBuilder represents the assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	WithBool(boolean bool) AssignableBuilder
	WithBytes(bytes []byte) AssignableBuilder
	WithContext(context uint) AssignableBuilder
	WithQueue(queue queues.Queue) AssignableBuilder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	IsBool() bool
	Bool() *bool
	IsBytes() bool
	Bytes() []byte
	IsContext() bool
	Context() *uint
	IsQueue() bool
	Queue() queues.Queue
}
