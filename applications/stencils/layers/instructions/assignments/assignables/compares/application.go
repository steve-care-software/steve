package compares

import (
	"github.com/steve-care-software/steve/applications/stencils/bytevalues"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/executions/instructions/assignments/assignables/compares"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/executions/instructions/assignments/assignables/compares/failures"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/frames"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/frames/assignables"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/symbols/layers"
)

type application struct {
	byteValuesApp     bytevalues.Application
	executionBuilder  executions.Builder
	failureBuilder    failures.Builder
	assignableBuilder assignables.AssignableBuilder
}

func createApplication(
	byteValuesApp bytevalues.Application,
	executionBuilder executions.Builder,
	failureBuilder failures.Builder,
	assignableBuilder assignables.AssignableBuilder,
) Application {
	out := application{
		byteValuesApp:     byteValuesApp,
		executionBuilder:  executionBuilder,
		failureBuilder:    failureBuilder,
		assignableBuilder: assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(values layers.ByteValues, frame frames.Frame) (executions.Compare, error) {
	assignables, err := app.byteValuesApp.Values(values, frame)
	if err != nil {
		return nil, err
	}

	builder := app.executionBuilder.Create()
	if !assignables.IsSuccess() {
		failure, err := app.failureBuilder.Create().
			WithCouldNotCompare(assignables).
			Now()

		if err != nil {
			return nil, err
		}

		return builder.WithFailure(failure).
			Now()
	}

	compare, err := assignables.Compare()
	if err != nil {
		return nil, err
	}

	assignableBuilder := app.assignableBuilder.Create().
		WithBytes([]byte{
			falseByte,
		})

	if compare {
		assignableBuilder.WithBytes([]byte{
			trueByte,
		})
	}

	ins, err := assignableBuilder.Now()
	if err != nil {
		return nil, err
	}

	return builder.WithSuccess(ins).
		Now()
}
