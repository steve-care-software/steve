package administrators

import (
	"github.com/steve-care-software/steve/applications/administrators/identities"
	"github.com/steve-care-software/steve/applications/administrators/stencils"
)

// Application represents the administrator application
type Application interface {
	Identity() (identities.Application, error)
	Stencil() (stencils.Application, error)
}
