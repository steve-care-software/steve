package creates

import "github.com/steve-care-software/steve/domain/stencils"

// Create represents a create account command
type Create interface {
	Stencil() stencils.Stencil
}
