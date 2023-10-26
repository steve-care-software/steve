package layers

import "github.com/steve-care-software/steve/domain/stencils/queries"

// Application represents a layer application
type Application interface {
	Execute(query queries.Query)
}
