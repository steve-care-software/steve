package encryptors

import (
	"github.com/steve-care-software/steve/applications/identities/encryptors/decrypts"
	"github.com/steve-care-software/steve/applications/identities/encryptors/publickeys"
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/encryptors"
	"github.com/steve-care-software/steve/domain/commands/executions/identities/encryptors/failures"
	"github.com/steve-care-software/steve/domain/commands/executions/identities/encryptors/successes"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/encryptors"
	"github.com/steve-care-software/steve/domain/stacks"
)

type application struct {
	decryptApp       decrypts.Application
	pubKeyApp        publickeys.Application
	failureBuilder   failures.Builder
	successBuilder   successes.Builder
	executionBuilder executions.Builder
}

func createApplication(
	decryptApp decrypts.Application,
	pubKeyApp publickeys.Application,
	failureBuilder failures.Builder,
	successBuilder successes.Builder,
	executionBuilder executions.Builder,
) Application {
	out := application{
		decryptApp:       decryptApp,
		pubKeyApp:        pubKeyApp,
		failureBuilder:   failureBuilder,
		successBuilder:   successBuilder,
		executionBuilder: executionBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(encryptor inputs.Encryptor, stack stacks.Stack) (executions.Encryptor, error) {
	name := encryptor.Name()
	assignable, err := stack.Fetch(name)
	if err != nil {
		failure, err := app.failureBuilder.Create().
			InstanceIsNotDeclared().
			WithName(name).
			Now()

		if err != nil {
			return nil, err
		}

		return app.executionBuilder.Create().
			WithFailure(failure).
			Now()
	}

	retBadKindFailureFn := func() (executions.Encryptor, error) {
		failure, err := app.failureBuilder.Create().
			InstanceIsNotEncryptor().
			WithName(name).
			Now()

		if err != nil {
			return nil, err
		}

		return app.executionBuilder.Create().
			WithFailure(failure).
			Now()
	}

	if !assignable.IsIdentity() {
		return retBadKindFailureFn()
	}

	assIdentity := assignable.Identity()
	if !assIdentity.IsEncryptor() {
		return retBadKindFailureFn()
	}

	current := assIdentity.Encryptor()
	content := encryptor.Content()
	variable := encryptor.AssignToVariable()
	successBuilder := app.successBuilder.Create().
		WithVariable(variable)

	if content.IsBytes() {
		bytes := current.Bytes()
		successBuilder.WithBytes(bytes)
	}

	if content.IsDecrypt() {
		decrypt := content.Decrypt()
		exec, err := app.decryptApp.Execute(decrypt, current)
		if err != nil {
			return nil, err
		}

		successBuilder.WithDecrypt(exec)
	}

	if content.IsPublicKey() {
		pubKey := content.PublicKey()
		exec, err := app.pubKeyApp.Execute(pubKey, current.Public())
		if err != nil {
			return nil, err
		}

		successBuilder.WithPublicKey(exec)
	}

	success, err := successBuilder.Now()
	if err != nil {
		return nil, err
	}

	return app.executionBuilder.Create().
		WithSuccess(success).
		Now()
}
