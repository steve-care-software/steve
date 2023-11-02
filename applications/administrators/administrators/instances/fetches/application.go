package fetches

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators/administrators/instances/successes/fetches"
	"github.com/steve-care-software/steve/domain/commands/executions/administrators/administrators/instances/successes/fetches/values"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators/instances/contents/fetches"
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
func (app *application) Execute(fetch inputs.Fetch, current administrators.Administrator) (executions.Fetch, error) {
	property := fetch.Property()
	valueBuilder := app.valueBuilder.Create()
	if property.IsUsername() {
		username := current.Username()
		valueBuilder.WithUsername(username)
	}

	if property.IsDashboard() {
		dashboard := current.Dashboard()
		valueBuilder.WithDashboard(dashboard)
	}

	if property.IsHasIdentities() {
		hasIdentities := current.HasIdentities()
		valueBuilder.WithHasIdentities(hasIdentities)
	}

	if property.IsIdentities() {
		identities := current.Identities()
		valueBuilder.WithIdentities(identities)
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
