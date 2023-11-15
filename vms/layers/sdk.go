package layers

import (
	"github.com/steve-care-software/steve/vms/layers/frames"
	"github.com/steve-care-software/steve/vms/layers/programs"
	"github.com/steve-care-software/steve/vms/layers/results"
)

// Layer represents a layer
type Layer interface {
	Execute(program programs.Program, frame frames.Frame) (results.Result, error)
}
