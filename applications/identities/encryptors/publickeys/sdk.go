package publickeys

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/encryptors"
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/encryptors/successes/publickeys"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/encryptors/contents/publickeys"
)

// Application represents the application
type Application interface {
	Execute(encryptor inputs.PublicKey, current encryptors.Encryptor) (executions.PublicKey, error)
}
