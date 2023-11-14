package conditions

import "github.com/steve-care-software/steve/domain/dashboards/stencils/symbols/layers"

// Application represents the application
type Application interface {
	Execute(condition layers.Condition, input []byte) error
}
