package creates

import (
	"github.com/steve-care-software/steve/domain/dashboards"
	"github.com/steve-care-software/steve/domain/stencils"
)

// Create represents a create account
type Create interface {
	Username() string
	Password() []byte
	Dashboard() dashboards.Dashboard
	Visitor() stencils.Stencil
}
