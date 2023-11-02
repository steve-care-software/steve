package instances

import (
	"github.com/steve-care-software/steve/applications/administrators/administrators/instances/deletes"
	"github.com/steve-care-software/steve/applications/administrators/administrators/instances/fetches"
	"github.com/steve-care-software/steve/applications/administrators/administrators/instances/updates"
	"github.com/steve-care-software/steve/applications/interpreters"
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators/instances"
	"github.com/steve-care-software/steve/domain/commands/executions/administrators/instances/failures"
	"github.com/steve-care-software/steve/domain/commands/executions/administrators/instances/successes"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators/instances"
)

type application struct {
	interpreterApp   interpreters.Application
	fetchApp         fetches.Application
	updateApp        updates.Application
	delApp           deletes.Application
	executionBuilder executions.Builder
	successBuilder   successes.Builder
	failureBuilder   failures.Builder
}

func createApplication(
	interpreterApp interpreters.Application,
	fetchApp fetches.Application,
	updateApp updates.Application,
	delApp deletes.Application,
	executionBuilder executions.Builder,
	successBuilder successes.Builder,
	failureBuilder failures.Builder,
) Application {
	out := application{
		interpreterApp:   interpreterApp,
		fetchApp:         fetchApp,
		updateApp:        updateApp,
		delApp:           delApp,
		executionBuilder: executionBuilder,
		successBuilder:   successBuilder,
		failureBuilder:   failureBuilder,
	}

	return &out
}

// Execute executes an application
func (app *application) Execute(instance inputs.Instance) (executions.Instance, error) {
	name := instance.Name()
	assignable, err := app.interpreterApp.Retrieve(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsAdministrator() {
		failure, err := app.failureBuilder.Create().
			InstanceIsNotAdministrator().
			WithName(name).
			Now()

		if err != nil {
			return nil, err
		}

		return app.executionBuilder.Create().
			WithFailure(failure).
			Now()
	}

	administrator := assignable.Administrator()
	content := instance.Content()
	successBuilder := app.successBuilder.Create()
	if content.IsFetch() {
		fetch := content.Fetch()
		exec, err := app.fetchApp.Execute(fetch, administrator)
		if err != nil {
			return nil, err
		}

		successBuilder.WithFetch(exec)
	}

	if content.IsUpdate() {
		update := content.Update()
		exec, err := app.updateApp.Execute(update, administrator)
		if err != nil {
			return nil, err
		}

		successBuilder.WithUpdate(exec)
	}

	if content.IsDelete() {
		del := content.Delete()
		exec, err := app.delApp.Execute(del, administrator)
		if err != nil {
			return nil, err
		}

		successBuilder.WithDelete(exec)
	}

	success, err := successBuilder.Now()
	if err != nil {
		return nil, err
	}

	return app.executionBuilder.Create().
		WithSuccess(success).
		Now()
}
