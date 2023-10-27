package libraries

import (
	"github.com/steve-care-software/steve/domain/stencils/libraries"
	"github.com/steve-care-software/steve/domain/stencils/libraries/results"
)

// NewApplication creates a new application
func NewApplication(
	service libraries.Service,
) Application {
	return createApplication(
		service,
	)
}

// Application represents the library application
type Application interface {
	Save(library libraries.Library) (results.Result, error)
}
