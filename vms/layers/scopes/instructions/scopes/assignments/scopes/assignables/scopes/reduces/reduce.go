package reduces

import (
	bytes_applications "github.com/steve-care-software/steve/vms/bytes/applications"
	"github.com/steve-care-software/steve/vms/layers/scopes/instructions/scopes/assignments/frames"
	"github.com/steve-care-software/steve/vms/layers/scopes/instructions/scopes/assignments/scopes/assignables/scopes/reduces/programs"
	"github.com/steve-care-software/steve/vms/layers/scopes/instructions/scopes/assignments/scopes/assignables/scopes/reduces/results"
)

type reduce struct {
	bytesApp       bytes_applications.Application
	resultBuilder  results.Builder
	failureBuilder results.FailureBuilder
}

func createReduce(
	bytesApp bytes_applications.Application,
	resultBuilder results.Builder,
	failureBuilder results.FailureBuilder,
) Reduce {
	out := reduce{
		bytesApp:       bytesApp,
		resultBuilder:  resultBuilder,
		failureBuilder: failureBuilder,
	}

	return &out
}

// Execute executes the program
func (app *reduce) Execute(input programs.Program, frame frames.Frame) (results.Result, error) {
	bytes := input.Bytes()
	exec, err := app.bytesApp.Program(bytes, frame.Bytes())
	if err != nil {
		return nil, err
	}

	if exec.IsFailure() {
		failure, err := app.failureBuilder.Create().
			WithBytesFailure(exec).
			Now()

		if err != nil {
			return nil, err
		}

		return app.resultBuilder.Create().
			WithFailure(failure).
			Now()
	}

	success := exec.Success()
	total := uint(len(success))
	requested := input.Length()
	if total < requested {
		failure, err := app.failureBuilder.WithNotEnoughBytes(requested).Now()
		if err != nil {
			return nil, err
		}

		return app.resultBuilder.Create().
			WithFailure(failure).
			Now()
	}

	return app.resultBuilder.Create().
		WithSuccess(success[:int(requested)]).
		Now()
}
