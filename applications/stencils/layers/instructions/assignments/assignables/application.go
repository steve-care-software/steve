package assignables

import (
	"github.com/steve-care-software/steve/applications/stencils/layers/instructions/assignments/assignables/compares"
	"github.com/steve-care-software/steve/applications/stencils/queries"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/executions/instructions/assignments/assignables"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/frames"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/frames/assignables"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/symbols/layers"
)

type application struct {
	compareApp        compares.Application
	queryApp          queries.Application
	executionBuilder  executions.Builder
	assignableBuilder assignables.AssignableBuilder
}

func createApplication(
	compareApp compares.Application,
	queryApp queries.Application,
	executionBuilder executions.Builder,
	assignableBuilder assignables.AssignableBuilder,
) Application {
	out := application{
		compareApp:        compareApp,
		queryApp:          queryApp,
		executionBuilder:  executionBuilder,
		assignableBuilder: assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(assignable layers.Assignable, frame frames.Frame) (executions.Assignable, error) {
	builder := app.executionBuilder.Create()
	if assignable.IsCompare() {
		compare := assignable.Compare()
		exec, err := app.compareApp.Execute(compare, frame)
		if err != nil {
			return nil, err
		}

		builder.WithCompare(exec)
	}

	if assignable.IsJoin() {

	}

	if assignable.IsLength() {

	}

	if assignable.IsReduce() {

	}

	if assignable.IsValue() {

	}

	if assignable.IsQuery() {
		query := assignable.Query()
		exec, err := app.queryApp.Execute(query, frame)
		if err != nil {
			return nil, err
		}

		builder.WithQuery(exec)
	}

	return builder.Now()
}
