package applications

import (
	"github.com/steve-care-software/steve/vms/domain/frames"
	"github.com/steve-care-software/steve/vms/domain/programs"
	"github.com/steve-care-software/steve/vms/domain/results"
)

// Application represents an application
type Application interface {
	Execute(programm programs.Program, frame frames.Frame) (results.Result, error)
}
