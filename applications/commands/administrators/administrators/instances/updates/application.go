package updates

import (
	"github.com/steve-care-software/steve/applications/commands/administrators/administrators/instances/updates/passwords"
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/administrators/instances/successes/updates"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/administrators/administrators/instances/contents/updates"
)

type application struct {
	passwordApp      passwords.Application
	executionBuilder executions.Builder
}

func createApplication(
	passwordApp passwords.Application,
	executionBuilder executions.Builder,
	current administrators.Administrator,
) Application {
	out := application{
		passwordApp:      passwordApp,
		executionBuilder: executionBuilder,
	}

	return &out
}

// Execute executes an application
func (app *application) Execute(update inputs.Update, current administrators.Administrator) (executions.Update, error) {
	builder := app.executionBuilder.Create()
	if update.IsPassword() {
		password := update.Password()
		exec, err := app.passwordApp.Execute(password, current)
		if err != nil {
			return nil, err
		}

		builder.WithPassword(exec)
	}

	return builder.Now()
}
