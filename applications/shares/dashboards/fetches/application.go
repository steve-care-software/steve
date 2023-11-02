package fetches

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/shares/dashboards/successes/fetches"
	"github.com/steve-care-software/steve/domain/commands/executions/shares/dashboards/successes/fetches/values"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/shares/dashboards/contents/fetches"
	"github.com/steve-care-software/steve/domain/dashboards"
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
func (app *application) Execute(fetch inputs.Fetch, current dashboards.Dashboard) (executions.Fetch, error) {
	content := fetch.Content()
	valueBuilder := app.valueBuilder.Create()
	if content.IsRoot() {
		root := current.Root()
		valueBuilder.WithRoot(root)
	}

	if content.IsStencils() {
		stencils := current.Stencils()
		valueBuilder.WithStencils(stencils)
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
