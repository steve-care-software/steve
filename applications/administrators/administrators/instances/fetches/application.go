package fetches

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators/instances/fetches"
	"github.com/steve-care-software/steve/domain/commands/executions/administrators/instances/fetches/values"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators/instances/contents/fetches"
)

type application struct {
	executionBuilder executions.Builder
	valueBuilder     values.Builder
	current          administrators.Administrator
}

func createApplication(
	executionBuilder executions.Builder,
	valueBuilder values.Builder,
	current administrators.Administrator,
) Application {
	out := application{
		executionBuilder: executionBuilder,
		valueBuilder:     valueBuilder,
		current:          current,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(fetch inputs.Fetch) (executions.Fetch, error) {
	property := fetch.Property()
	valueBuilder := app.valueBuilder.Create()
	if property.IsUsername() {
		username := app.current.Username()
		valueBuilder.WithUsername(username)
	}

	if property.IsDashboard() {
		dashboard := app.current.Dashboard()
		valueBuilder.WithDashboard(dashboard)
	}

	if property.IsHasIdentities() {
		hasIdentities := app.current.HasIdentities()
		valueBuilder.WithHasIdentities(hasIdentities)
	}

	if property.IsIdentities() {
		identities := app.current.Identities()
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
