package encrypts

import (
	identity_publickeys "github.com/steve-care-software/steve/domain/accounts/identities/encryptors/publickeys"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/identities/encryptors/successes/publickeys/encrypts"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/identities/encryptors/contents/publickeys/encrypts"
)

type application struct {
	executionBuilder executions.Builder
}

func createApplication(
	executionBuilder executions.Builder,
) Application {
	out := application{
		executionBuilder: executionBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(encryptor inputs.Encrypt, current identity_publickeys.PublicKey) (executions.Encrypt, error) {
	msg := encryptor.Message()
	cipher, err := current.Encrypt(msg)
	if err != nil {
		return nil, err
	}

	return app.executionBuilder.Create().
		WithCipher(cipher).
		Now()
}
