package applications

import (
	"github.com/steve-care-software/steve/vms/children/bytes/frames"
	"github.com/steve-care-software/steve/vms/children/bytes/programs"
	"github.com/steve-care-software/steve/vms/children/bytes/results"
)

// Application represents the bytes application
type Application interface {
	Programs(programs programs.Programs, frame frames.Frame) (results.Results, error)
	Program(program programs.Program, frame frames.Frame) (results.Result, error)
}
