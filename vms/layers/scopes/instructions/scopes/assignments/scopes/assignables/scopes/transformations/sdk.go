package transformations

import (
	bytes_programs "github.com/steve-care-software/steve/vms/bytes/programs"
	"github.com/steve-care-software/steve/vms/layers/scopes/instructions/scopes/assignments/frames"
	"github.com/steve-care-software/steve/vms/layers/scopes/instructions/scopes/assignments/scopes/assignables/scopes/transformations/results"
)

// Transformation represents a transformation
type Transformation interface {
	Execute(input bytes_programs.Programs, frame frames.Frame) (results.Result, error)
}
