package decrypts

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/encryptors"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/identities/encryptors/successes/decrypts"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/identities/encryptors/contents/decrypts"
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
func (app *application) Execute(decrypt inputs.Decrypt, current encryptors.Encryptor) (executions.Decrypt, error) {
	cipher := decrypt.Cipher()
	msg, err := current.Decrypt(cipher)
	if err != nil {
		return nil, err
	}

	return app.executionBuilder.Create().
		WithMessage(msg).
		Now()
}
