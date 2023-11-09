package encrypts

import (
	identity_publickeys "github.com/steve-care-software/steve/domain/accounts/identities/encryptors/publickeys"
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/encryptors/successes/publickeys/encrypts"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/encryptors/contents/publickeys/encrypts"
)

// Application represents the application
type Application interface {
	Execute(encryptor inputs.Encrypt, current identity_publickeys.PublicKey) (executions.Encrypt, error)
}
