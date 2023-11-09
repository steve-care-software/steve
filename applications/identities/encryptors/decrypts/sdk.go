package decrypts

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/encryptors"
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/encryptors/successes/decrypts"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/encryptors/contents/decrypts"
)

// Application represents the application
type Application interface {
	Execute(decrypt inputs.Decrypt, current encryptors.Encryptor) (executions.Decrypt, error)
}
