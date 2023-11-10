package instances

import (
	"github.com/steve-care-software/steve/applications/commands/administrators/administrators/instances/deletes"
	"github.com/steve-care-software/steve/applications/commands/administrators/administrators/instances/fetches"
	"github.com/steve-care-software/steve/applications/commands/administrators/administrators/instances/updates"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/administrators/instances"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/administrators/instances/failures"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/administrators/instances/successes"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/administrators/administrators/instances"
)

type application struct {
	fetchApp         fetches.Application
	updateApp        updates.Application
	delApp           deletes.Application
	executionBuilder executions.Builder
	successBuilder   successes.Builder
	failureBuilder   failures.Builder
}

func createApplication(
	fetchApp fetches.Application,
	updateApp updates.Application,
	delApp deletes.Application,
	executionBuilder executions.Builder,
	successBuilder successes.Builder,
	failureBuilder failures.Builder,
) Application {
	out := application{
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
func (app *application) Execute(instance inputs.Instance, frame frames.Frame) (executions.Instance, error) {
	name := instance.Name()
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

	retBadKindFailureFn := func() (executions.Instance, error) {
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

	if !assignable.IsAdministrator() {
		return retBadKindFailureFn()
	}

	assAdmin := assignable.Administrator()
	if !assAdmin.IsInstance() {
		return retBadKindFailureFn()
	}

	administrator := assAdmin.Instance()
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
