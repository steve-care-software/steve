package reduces

import (
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/frames"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/scopes/assignables/scopes/reduces/programs"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/scopes/assignables/scopes/reduces/results"
)

// Reduce represents a reduce
type Reduce interface {
	Execute(input programs.Program, frame frames.Frame) (results.Result, error)
}
