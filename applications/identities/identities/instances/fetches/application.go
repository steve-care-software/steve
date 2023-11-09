package fetches

import (
	"github.com/steve-care-software/steve/domain/accounts/identities"
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/identities/instances/successes/fetches"
	"github.com/steve-care-software/steve/domain/commands/executions/identities/identities/instances/successes/fetches/values"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/identities/instances/contents/fetches"
)

type application struct {
	executionBuilder executions.Builder
	valueBuilder     values.Builder
}

func createApplication(
	executionBuilder executions.Builder,
	valueBuilder values.Builder,
) Application {
	out := application{
		executionBuilder: executionBuilder,
		valueBuilder:     valueBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(fetch inputs.Fetch, current identities.Identity) (executions.Fetch, error) {
	property := fetch.Property()
	valueBuilder := app.valueBuilder.Create()

	if property.IsDashboard() {
		dashboard := current.Dashboard()
		valueBuilder.WithDashboard(dashboard)
	}

	if property.IsProfile() {
		profile := current.Profile()
		valueBuilder.WithProfile(profile)
	}

	if property.IsEncryptor() {
		encryptor := current.Encryptor()
		valueBuilder.WithEncryptor(encryptor)
	}

	if property.IsSigner() {
		signer := current.Signer()
		valueBuilder.WithSigner(signer)
	}

	if property.IsHasShares() {
		hasShares := current.HasShares()
		valueBuilder.WithHasShares(hasShares)
	}

	if property.IsShares() {
		shares := current.Shares()
		valueBuilder.WithShares(shares)
	}

	if property.IsHasConnections() {
		hasConnections := current.HasConnections()
		valueBuilder.WithHasConnections(hasConnections)
	}

	if property.IsConnections() {
		connections := current.Connections()
		valueBuilder.WithConnections(connections)
	}

	value, err := valueBuilder.Now()
	if err != nil {
		return nil, err
	}

	assignTo := fetch.AssignTo()
	return app.executionBuilder.Create().
		WithAssignTo(assignTo).
		WithValue(value).
		Now()
}
