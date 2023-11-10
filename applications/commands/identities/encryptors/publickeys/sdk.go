package publickeys

import (
	identity_publickeys "github.com/steve-care-software/steve/domain/accounts/identities/encryptors/publickeys"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/identities/encryptors/successes/publickeys"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/identities/encryptors/contents/publickeys"
)

// Application represents the application
type Application interface {
	Execute(encryptor inputs.PublicKey, current identity_publickeys.PublicKey) (executions.PublicKey, error)
}
