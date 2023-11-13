package updates

import (
	"github.com/steve-care-software/steve/applications/commands/identities/identities/instances/updates/passwords"
	"github.com/steve-care-software/steve/domain/accounts/identities"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/identities/identities/instances/successes/updates"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/identities/identities/instances/contents/updates"
)

type application struct {
	passApp          passwords.Application
	executionBuilder executions.Builder
}

func createApplication(
	passApp passwords.Application,
	executionBuilder executions.Builder,
) Application {
	out := application{
		passApp:          passApp,
		executionBuilder: executionBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(update inputs.Update, current identities.Identity) (executions.Update, error) {
	builder := app.executionBuilder.Create()
	if update.IsPassword() {
		pass := update.Password()
		exec, err := app.passApp.Execute(pass, current)
		if err != nil {
			return nil, err
		}

		builder.WithPassword(exec)
	}

	return builder.Now()
}
