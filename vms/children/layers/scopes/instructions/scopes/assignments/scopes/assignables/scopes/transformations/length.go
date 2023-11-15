package transformations

import (
	"encoding/binary"

	bytes_applications "github.com/steve-care-software/steve/vms/children/bytes/applications"
	bytes_programs "github.com/steve-care-software/steve/vms/children/bytes/programs"
	"github.com/steve-care-software/steve/vms/children/layers/scopes/instructions/scopes/assignments/frames"
	"github.com/steve-care-software/steve/vms/children/layers/scopes/instructions/scopes/assignments/scopes/assignables/scopes/transformations/results"
)

type length struct {
	bytesApp      bytes_applications.Application
	resultBuilder results.Builder
}

func createLength(
	bytesApp bytes_applications.Application,
	resultBuilder results.Builder,
) Transformation {
	out := length{
		bytesApp:      bytesApp,
		resultBuilder: resultBuilder,
	}

	return &out
}

// Execute executes the length
func (app *length) Execute(input bytes_programs.Programs, frame frames.Frame) (results.Result, error) {
	exec, err := app.bytesApp.Programs(input, frame.Bytes())
	if err != nil {
		return nil, err
	}

	if exec.HasFailure() {
		return app.resultBuilder.Create().
			WithFailure(exec).
			Now()
	}

	list := exec.List()
	amount := len(list)

	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(amount))
	return app.resultBuilder.Create().
		WithSuccess(bytes).
		Now()
}
