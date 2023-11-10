package identities

import (
	"github.com/steve-care-software/steve/applications/commands/administrators/identities/deletes"
	"github.com/steve-care-software/steve/applications/commands/administrators/identities/fetches"
	"github.com/steve-care-software/steve/applications/commands/administrators/identities/inserts"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/identities"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/identities/failures"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/identities/successes"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/administrators/identities"
)

type application struct {
	fetchApp         fetches.Application
	insertApp        inserts.Application
	delApp           deletes.Application
	executionBuilder executions.Builder
	successBuilder   successes.Builder
	failureBuilder   failures.Builder
}

func createApplication(
	fetchApp fetches.Application,
	insertApp inserts.Application,
	delApp deletes.Application,
	executionBuilder executions.Builder,
	successBuilder successes.Builder,
	failureBuilder failures.Builder,
) Application {
	out := application{
		fetchApp:         fetchApp,
		insertApp:        insertApp,
		delApp:           delApp,
		executionBuilder: executionBuilder,
		successBuilder:   successBuilder,
		failureBuilder:   failureBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(instance inputs.Identities, frame frames.Frame) (executions.Identities, error) {
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

	retBadKindFailureFn := func() (executions.Identities, error) {
		failure, err := app.failureBuilder.Create().
			InstanceIsNotIdentities().
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
	if !assAdmin.IsIdentities() {
		return retBadKindFailureFn()
	}

	identities := assAdmin.Identities()
	content := instance.Content()
	successBuilder := app.successBuilder.Create()
	if content.IsFetch() {
		fetch := content.Fetch()
		exec, err := app.fetchApp.Execute(fetch, identities)
		if err != nil {
			return nil, err
		}

		successBuilder.WithFetch(exec)
	}

	if content.IsInsert() {
		insert := content.Insert()
		exec, err := app.insertApp.Execute(insert, identities)
		if err != nil {
			return nil, err
		}

		successBuilder.WithInsert(exec)
	}

	if content.IsDelete() {
		del := content.Delete()
		exec, err := app.delApp.Execute(del, identities)
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
