package suites

import "github.com/steve-care-software/steve/domain/dashboards/stencils/symbols/layers"

// Application represents the application
type Application interface {
	Suites(suites layers.Suites, input []byte) error
	Suite(suite layers.Suite, input []byte) error
}
