package results

import (
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
	"github.com/steve-care-software/steve/domain/stencils/results/inputs"
	"github.com/steve-care-software/steve/domain/stencils/results/preparations"
)

// Builder represents the result builder
type Builder interface {
	Create() Builder
	WithInput(input inputs.Input) Builder
	WithPreparations(preparations preparations.Preparations) Builder
	WithExecute(execute layers.Layer) Builder
	WithPrevious(previous Result) Builder
	Now() (Result, error)
}

// Result represents result
type Result interface {
	Input() inputs.Input
	Preparations() preparations.Preparations
	Execute() layers.Layer
	HasPrevious() bool
	Previous() Result
}
