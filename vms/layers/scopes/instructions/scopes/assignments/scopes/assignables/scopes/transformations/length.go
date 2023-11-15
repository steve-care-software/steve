package transformations

import (
	"encoding/binary"

	bytes_vm "github.com/steve-care-software/steve/vms/bytes"
	bytes_programs "github.com/steve-care-software/steve/vms/bytes/programs"
	"github.com/steve-care-software/steve/vms/layers/scopes/instructions/scopes/assignments/frames"
	"github.com/steve-care-software/steve/vms/layers/scopes/instructions/scopes/assignments/scopes/assignables/scopes/transformations/results"
)

type length struct {
	vmBytes       bytes_vm.Bytes
	resultBuilder results.Builder
}

func createLength(
	vmBytes bytes_vm.Bytes,
	resultBuilder results.Builder,
) Transformation {
	out := length{
		vmBytes:       vmBytes,
		resultBuilder: resultBuilder,
	}

	return &out
}

// Execute executes the length
func (app *length) Execute(input bytes_programs.Programs, frame frames.Frame) (results.Result, error) {
	exec, err := app.vmBytes.Programs(input, frame.Bytes())
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
