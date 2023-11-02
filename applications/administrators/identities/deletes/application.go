package deletes

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators/identities"
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators/identities/deletes"
	"github.com/steve-care-software/steve/domain/commands/executions/administrators/identities/deletes/failures"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators/identities/contents/deletes"
)

type application struct {
	executionBuilder executions.Builder
	failureBuilder   failures.Builder
}

func createApplication(
	executionBuilder executions.Builder,
	failureBuilder failures.Builder,
) Application {
	out := application{
		executionBuilder: executionBuilder,
		failureBuilder:   failureBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(instance inputs.Delete, current identities.Identities) (executions.Delete, error) {
	index := instance.Index()
	if current.Exceeds(index) {
		failure, err := app.failureBuilder.Create().
			IsIndexExceedAmount().
			WithIndex(index).
			Now()

		if err != nil {
			return nil, err
		}

		return app.executionBuilder.Create().
			WithFailure(failure).
			Now()
	}

	updated, err := current.Delete(index)
	if err != nil {
		return nil, err
	}

	return app.executionBuilder.Create().
		WithSuccess(updated).
		Now()
}
