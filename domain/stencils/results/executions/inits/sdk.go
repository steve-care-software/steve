package inits

import (
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns"
	"github.com/steve-care-software/steve/domain/stencils/results/executions/inits/inputs"
	"github.com/steve-care-software/steve/domain/stencils/results/executions/inits/values"
)

// Builder represents an init builder
type Builder interface {
	Create() Builder
	WithInput(input inputs.Input) Builder
	WithReturn(ret returns.Return) Builder
	WithValues(values values.Values) Builder
	Now() (Init, error)
}

// Init represents a layer init
type Init interface {
	Input() inputs.Input
	Return() returns.Return
	HasValues() bool
	Values() values.Values
}
