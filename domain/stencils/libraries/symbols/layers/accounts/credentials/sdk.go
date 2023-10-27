package credentials

import "github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/constantvalues"

// Credentials represents credentials
type Credentials interface {
	Username() constantvalues.ConstantValue
	Password() constantvalues.ConstantValue
}
