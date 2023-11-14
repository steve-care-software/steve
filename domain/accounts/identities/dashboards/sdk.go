package dashboards

import "github.com/steve-care-software/steve/domain/dashboards/stencils"

// Dashboard represents a dashboard
type Dashboard interface {
	Stencils() stencils.Stencils
	Root() stencils.Stencil
}
