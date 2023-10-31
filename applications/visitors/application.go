package visitors

import (
	account_applications "github.com/steve-care-software/steve/applications/visitors/accounts"
	admin_applications "github.com/steve-care-software/steve/applications/visitors/administrators"
	executions "github.com/steve-care-software/steve/domain/commands/executions/visitors"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/visitors"
)

type application struct {
	accountApp       account_applications.Application
	adminApp         admin_applications.Application
	executionBuilder executions.Builder
}

func createApplication(
	accountApp account_applications.Application,
	adminApp admin_applications.Application,
	executionBuilder executions.Builder,
) Application {
	out := application{
		accountApp:       accountApp,
		adminApp:         adminApp,
		executionBuilder: executionBuilder,
	}

	return &out
}

// Execute executes a visitor's application
func (app *application) Execute(visitor inputs.Visitor) (executions.Visitor, error) {
	builder := app.executionBuilder.Create()
	if visitor.IsAdministrator() {
		administrator := visitor.Administrator()
		exec, err := app.adminApp.Execute(administrator)
		if err != nil {
			return nil, err
		}

		builder.WithAdministrator(exec)
	}

	if visitor.IsAccount() {
		account := visitor.Account()
		exec, err := app.accountApp.Execute(account)
		if err != nil {
			return nil, err
		}

		builder.WithAccount(exec)
	}

	return builder.Now()
}
