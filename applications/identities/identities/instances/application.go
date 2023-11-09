package instances

import (
	"github.com/steve-care-software/steve/applications/identities/identities/instances/deletes"
	"github.com/steve-care-software/steve/applications/identities/identities/instances/fetches"
	"github.com/steve-care-software/steve/applications/identities/identities/instances/updates"
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/identities/instances"
	"github.com/steve-care-software/steve/domain/commands/executions/identities/identities/instances/failures"
	"github.com/steve-care-software/steve/domain/commands/executions/identities/identities/instances/successes"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/identities/instances"
	"github.com/steve-care-software/steve/domain/stacks"
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
func (app *application) Execute(instance inputs.Instance, stack stacks.Stack) (executions.Instance, error) {
	name := instance.Name()
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

	retBadKindFailureFn := func() (executions.Instance, error) {
		failure, err := app.failureBuilder.Create().
			InstanceIsNotIdentity().
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
	if !assIdentity.IsInstance() {
		return retBadKindFailureFn()
	}

	identity := assIdentity.Instance()
	content := instance.Content()
	successBuilder := app.successBuilder.Create()
	if content.IsFetch() {
		fetch := content.Fetch()
		exec, err := app.fetchApp.Execute(fetch, identity)
		if err != nil {
			return nil, err
		}

		successBuilder.WithFetch(exec)
	}

	if content.IsUpdate() {
		update := content.Update()
		exec, err := app.updateApp.Execute(update, identity)
		if err != nil {
			return nil, err
		}

		successBuilder.WithUpdate(exec)
	}

	if content.IsDelete() {
		del := content.Delete()
		exec, err := app.delApp.Execute(del, identity)
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
