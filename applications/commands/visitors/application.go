package visitors

import (
	admin_applications "github.com/steve-care-software/steve/applications/commands/shares/administrators/creates"
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/visitors"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/visitors/failures"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/visitors/successes"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/visitors"
)

type application struct {
	adminApp         admin_applications.Application
	adminRepository  administrators.Repository
	failureBuilder   failures.Builder
	successBuilder   successes.Builder
	executionBuilder executions.Builder
}

func createApplication(
	adminApp admin_applications.Application,
	adminRepository administrators.Repository,
	failureBuilder failures.Builder,
	successBuilder successes.Builder,
	executionBuilder executions.Builder,
) Application {
	out := application{
		adminApp:         adminApp,
		failureBuilder:   failureBuilder,
		successBuilder:   successBuilder,
		executionBuilder: executionBuilder,
	}

	return &out
}

// Execute executes a visitor's application
func (app *application) Execute(visitor inputs.Visitor) (executions.Visitor, error) {
	builder := app.executionBuilder.Create()
	if visitor.IsAdministrator() {
		isInitialized, err := app.adminRepository.IsInitialized()
		if err != nil {
			return nil, err
		}

		if !isInitialized {
			failure, err := app.failureBuilder.Create().
				AdminAlreadyInitialized().
				Now()

			if err != nil {
				return nil, err
			}

			return app.executionBuilder.Create().
				WithFailure(failure).
				Now()
		}

		administrator := visitor.Administrator()
		exec, err := app.adminApp.Execute(administrator)
		if err != nil {
			return nil, err
		}

		success, err := app.successBuilder.Create().WithAdministrator(exec).Now()
		if err != nil {
			return nil, err
		}

		builder.WithSuccess(success)
	}

	return builder.Now()
}
