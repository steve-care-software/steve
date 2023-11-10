package dashboards

import (
	"github.com/steve-care-software/steve/applications/commands/shares/dashboards/fetches"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/shares/dashboards"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/shares/dashboards/failures"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/shares/dashboards/successes"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/shares/dashboards"
)

type application struct {
	fetchApp         fetches.Application
	executionBuilder executions.Builder
	successBuilder   successes.Builder
	failureBuilder   failures.Builder
}

func createApplication(
	fetchApp fetches.Application,
	executionBuilder executions.Builder,
	successBuilder successes.Builder,
	failureBuilder failures.Builder,
) Application {
	out := application{
		fetchApp:         fetchApp,
		executionBuilder: executionBuilder,
		successBuilder:   successBuilder,
		failureBuilder:   failureBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(dashboard inputs.Dashboard, frame frames.Frame) (executions.Dashboard, error) {
	name := dashboard.Name()
	assignable, err := frame.Fetch(name)
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

	if !assignable.IsDashboard() {
		failure, err := app.failureBuilder.Create().
			InstanceIsNotDashboard().
			WithName(name).
			Now()

		if err != nil {
			return nil, err
		}

		return app.executionBuilder.Create().
			WithFailure(failure).
			Now()
	}

	current := assignable.Dashboard()
	content := dashboard.Content()
	builder := app.executionBuilder.Create()
	if content.IsFetch() {
		fetch := content.Fetch()
		exec, err := app.fetchApp.Execute(fetch, current)
		if err != nil {
			return nil, err
		}

		success, err := app.successBuilder.Create().
			WithFetch(exec).
			Now()

		if err != nil {
			return nil, err
		}

		builder.WithSuccess(success)
	}

	return builder.Now()
}
