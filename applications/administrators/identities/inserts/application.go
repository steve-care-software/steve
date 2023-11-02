package inserts

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators/identities"
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators/identities/inserts"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators/identities/contents/inserts"
)

type application struct {
	executionBuilder  executions.Builder
	identitiesBuilder identities.Builder
	identityBuilder   identities.IdentityBuilder
}

func createApplication(
	executionBuilder executions.Builder,
	identitiesBuilder identities.Builder,
	identityBuilder identities.IdentityBuilder,
) Application {
	out := application{
		executionBuilder:  executionBuilder,
		identitiesBuilder: identitiesBuilder,
		identityBuilder:   identityBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(instance inputs.Insert, current identities.Identities) (executions.Insert, error) {
	name := instance.Name()
	container := instance.Container()
	identity, err := app.identityBuilder.Create().
		WithName(name).
		WithContainer(container).
		Now()

	if err != nil {
		return nil, err
	}

	list := current.List()
	list = append(list, identity)
	identities, err := app.identitiesBuilder.Create().
		WithList(list).
		Now()

	if err != nil {
		return nil, err
	}

	return app.executionBuilder.Create().
		WithIdentities(identities).
		Now()
}
