package assignments

import (
	"github.com/steve-care-software/steve/vms/layers/scopes/instructions/scopes/assignments/frames"
	"github.com/steve-care-software/steve/vms/layers/scopes/instructions/scopes/assignments/programs"
)

// Assignment represents an assignment
type Assignment interface {
	Execute(input programs.Program, frame frames.Frame) (frames.Frame, error)
}
