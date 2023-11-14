package bytevalues

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/bytevalues"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/bytevalues/failures"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/frames"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/frames/assignables"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/symbols/layers"
)

type application struct {
	executionsBuilder  executions.Builder
	executionBuilder   executions.ByteValueBuilder
	failureBuilder     failures.Builder
	assignablesBuilder assignables.Builder
	assignableBuilder  assignables.AssignableBuilder
}

func createApplication(
	executionsBuilder executions.Builder,
	executionBuilder executions.ByteValueBuilder,
	failureBuilder failures.Builder,
	assignablesBuilder assignables.Builder,
	assignableBuilder assignables.AssignableBuilder,
) Application {
	out := application{
		executionsBuilder:  executionsBuilder,
		executionBuilder:   executionBuilder,
		failureBuilder:     failureBuilder,
		assignablesBuilder: assignablesBuilder,
		assignableBuilder:  assignableBuilder,
	}

	return &out
}

// Values executes the values
func (app *application) Values(values layers.ByteValues, frame frames.Frame) (executions.ByteValues, error) {
	list := []executions.ByteValue{}
	valueList := values.List()
	for _, oneValue := range valueList {
		execution, err := app.Value(oneValue, frame)
		if err != nil {
			return nil, err
		}

		list = append(list, execution)
	}

	return app.executionsBuilder.Create().
		WithList(list).
		Now()
}

// Value executes the value
func (app *application) Value(value layers.ByteValue, frame frames.Frame) (executions.ByteValue, error) {
	builder := app.executionBuilder.Create()
	if value.IsBytes() {
		bytes := value.Bytes()
		return builder.WithSuccess(bytes).Now()
	}

	name := value.Variable()
	ins, err := frame.Fetch(name)
	if err != nil {
		failure, err := app.failureBuilder.Create().
			CouldNotFetchVariable().
			Now()

		if err != nil {
			return nil, err
		}

		return builder.WithFailure(failure).
			Now()
	}

	if !ins.IsBytes() {
		failure, err := app.failureBuilder.Create().
			VariableIsNotBytes().
			Now()

		if err != nil {
			return nil, err
		}

		return builder.WithFailure(failure).
			Now()
	}

	bytes := ins.Bytes()
	return builder.WithSuccess(bytes).
		Now()
}
