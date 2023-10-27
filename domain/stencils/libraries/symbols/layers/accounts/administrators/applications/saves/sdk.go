package saves

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/constantvalues"
)

// Save represents a save
type Save interface {
	Instance() administrators.Administrator
	Password() constantvalues.ConstantValue
	HasNewPassword() bool
	NewPassword() constantvalues.ConstantValue
}
