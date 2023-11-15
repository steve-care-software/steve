package frames

import (
	bytes_programs "github.com/steve-care-software/steve/vms/children/bytes/programs"
)

// Builder represents a frame builder
type Builder interface {
	Create() Builder
	WithList(list []Assignment) Builder
	Now() (Frame, error)
}

// Frame represents a frame
type Frame interface {
	List() []Assignment
	Fetch(name string) (Assignment, error)
}

// Assignment represents a frame assignment
type Assignment interface {
	Name() string
	Assignable() bytes_programs.Program
}
