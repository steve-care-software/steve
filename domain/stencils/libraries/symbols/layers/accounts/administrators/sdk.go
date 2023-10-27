package administrators

import (
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/accounts/administrators/applications"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/accounts/administrators/instances"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/constantvalues"
)

// Administrator represents an administrator assignable
type Administrator interface {
	Application() constantvalues.ConstantValue
	Content() Content
}

// Content represents an administrator
type Content interface {
	IsApplication() bool
	Application() applications.Application
	IsInstance() bool
	Instance() instances.Instance
}
