package applications

import (
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/accounts/administrators/applications/saves"
)

// Application represents application funcs
type Application interface {
	IsRetrieve() bool
	IsSave() bool
	Save() saves.Save
}
