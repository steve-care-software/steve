package frames

import (
	bytes_programs "github.com/steve-care-software/steve/vms/bytes/programs"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/programs"
)

// Frame represents a frame
type Frame interface {
	List() []Assignment
	Fetch(name string) (Assignment, error)
}

// Assignment represents a frame assignment
type Assignment interface {
	Name() string
	Assignable() Assignable
}

// Assignable represents an assignable
type Assignable interface {
	IsLayer() bool
	Layer() programs.Program
	IsBytes() bool
	Bytes() bytes_programs.Program
}
