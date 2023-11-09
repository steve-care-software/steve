package accounts

import (
	"github.com/steve-care-software/steve/applications/visitors/accounts/creates"
	execution_accounts "github.com/steve-care-software/steve/domain/commands/executions/visitors/accounts"
	executions "github.com/steve-care-software/steve/domain/commands/executions/visitors/accounts"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/visitors/accounts"
)

type application struct {
	createApp        creates.Application
	executionBuilder execution_accounts.Builder
}

func createApplication(
	createApp creates.Application,
	executionBuilder execution_accounts.Builder,
) Application {
	out := application{
		createApp:        createApp,
		executionBuilder: executionBuilder,
	}

	return &out
}

// Execute executes an application
func (app *application) Execute(account inputs.Account) (executions.Account, error) {
	builder := app.executionBuilder.Create()
	if account.IsCreate() {
		create := account.Create()
		exec, err := app.createApp.Execute(create)
		if err != nil {
			return nil, err
		}

		builder.WithCreate(exec)
	}

	return builder.Now()
}
