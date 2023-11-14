package compares

import (
	bytes_programs "github.com/steve-care-software/steve/vms/bytes/programs"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/frames"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/scopes/assignables/scopes/compares/results"
)

// Compare represents a compare
type Compare interface {
	Execute(input bytes_programs.Programs, frame frames.Frame) (results.Result, error)
}
