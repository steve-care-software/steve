package dashboards

import "github.com/steve-care-software/steve/domain/stencils"

// Dashboard represents a dashboard
type Dashboard interface {
	Stencils() stencils.Stencils
	Root() stencils.Stencil
}
