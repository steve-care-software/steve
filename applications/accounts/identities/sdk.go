package identities

import (
	"github.com/steve-care-software/steve/domain/accounts/identities"
	layer_identities "github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/identities"
)

// Application represents the identity application
type Application interface {
	List() []string
	Retrieve(retrieve layer_identities.Retrieve) (identities.Identity, error)
	Insert(insert layer_identities.Insert) error
	Update(update layer_identities.Update) error
	Delete(del layer_identities.Delete) error
}
