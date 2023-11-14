package stencils

import (
	"github.com/steve-care-software/steve/domain/dashboards/stencils"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/results"
)

// Application represents the stencil application
type Application interface {
	Execute(stencil stencils.Stencil, input []byte) (results.Result, error)
}
