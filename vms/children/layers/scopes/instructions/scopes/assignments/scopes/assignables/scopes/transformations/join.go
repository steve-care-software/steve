package transformations

import (
	bytes_applications "github.com/steve-care-software/steve/vms/children/bytes/applications"
	bytes_programs "github.com/steve-care-software/steve/vms/children/bytes/programs"
	"github.com/steve-care-software/steve/vms/children/layers/scopes/instructions/scopes/assignments/frames"
	"github.com/steve-care-software/steve/vms/children/layers/scopes/instructions/scopes/assignments/scopes/assignables/scopes/transformations/results"
)

type join struct {
	bytesApp      bytes_applications.Application
	resultBuilder results.Builder
}

func createJoin(
	bytesApp bytes_applications.Application,
	resultBuilder results.Builder,
) Transformation {
	out := join{
		bytesApp:      bytesApp,
		resultBuilder: resultBuilder,
	}

	return &out
}

// Execute executes the join
func (app *join) Execute(input bytes_programs.Programs, frame frames.Frame) (results.Result, error) {
	exec, err := app.bytesApp.Programs(input, frame.Bytes())
	if err != nil {
		return nil, err
	}

	if exec.HasFailure() {
		return app.resultBuilder.Create().
			WithFailure(exec).
			Now()
	}

	bytes := []byte{}
	list := exec.List()
	for _, oneResult := range list {
		bytes = append(bytes, oneResult.Success()...)
	}

	return app.resultBuilder.Create().
		WithSuccess(bytes).
		Now()
}
