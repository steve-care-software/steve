package assignables

import (
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/instructions/scopes/assignments/frames"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/instructions/scopes/assignments/scopes/assignables/programs"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/instructions/scopes/assignments/scopes/assignables/results"
)

// Assignable represents an assignable
type Assignable interface {
	Execute(input programs.Program, frame frames.Frame) (results.Result, error)
}
