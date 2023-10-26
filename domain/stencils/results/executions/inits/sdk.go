package inits

import (
	returns "github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns/expectations"
	"github.com/steve-care-software/steve/domain/stencils/results/executions/inits/inputs"
	"github.com/steve-care-software/steve/domain/stencils/results/executions/inits/values"
)

// Builder represents an init builder
type Builder interface {
	Create() Builder
	WithInput(input inputs.Input) Builder
	WithReturn(ret returns.Expectation) Builder
	WithValues(values values.Values) Builder
	Now() (Init, error)
}

// Init represents a layer init
type Init interface {
	Input() inputs.Input
	Return() returns.Expectation
	HasValues() bool
	Values() values.Values
}
