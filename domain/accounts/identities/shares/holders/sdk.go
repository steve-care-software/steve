package holders

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/profiles"
	"github.com/steve-care-software/steve/domain/accounts/identities/shares/values"
)

// Holder represents a shareholder
type Holder interface {
	Profile() profiles.Profile
	Values() values.Values
	Amount() uint
}
