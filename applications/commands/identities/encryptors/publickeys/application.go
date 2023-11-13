package publickeys

import (
	"github.com/steve-care-software/steve/applications/commands/identities/encryptors/publickeys/encrypts"
	identity_publickeys "github.com/steve-care-software/steve/domain/accounts/identities/encryptors/publickeys"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/identities/encryptors/successes/publickeys"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/identities/encryptors/contents/publickeys"
)

type application struct {
	encryptApp       encrypts.Application
	executionBuilder executions.Builder
}

func createApplication(
	encryptApp encrypts.Application,
	executionBuilder executions.Builder,
) Application {
	out := application{
		encryptApp:       encryptApp,
		executionBuilder: executionBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(encryptor inputs.PublicKey, current identity_publickeys.PublicKey) (executions.PublicKey, error) {
	builder := app.executionBuilder.Create()
	if encryptor.IsBytes() {
		bytes := current.Bytes()
		builder.WithBytes(bytes)
	}

	if encryptor.IsEncrypt() {
		encrypt := encryptor.Encrypt()
		exec, err := app.encryptApp.Execute(encrypt, current)
		if err != nil {
			return nil, err
		}

		builder.WithEncrypt(exec)
	}

	return builder.Now()
}
