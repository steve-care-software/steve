package transformations

import (
	"bytes"

	bytes_vm "github.com/steve-care-software/steve/vms/bytes"
	bytes_programs "github.com/steve-care-software/steve/vms/bytes/programs"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/frames"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/scopes/assignables/scopes/transformations/results"
)

type compare struct {
	vmBytes       bytes_vm.Bytes
	resultBuilder results.Builder
	trueBytes     []byte
	falseBytes    []byte
}

func createCompare(
	vmBytes bytes_vm.Bytes,
	resultBuilder results.Builder,
	trueBytes []byte,
	falseBytes []byte,
) Transformation {
	out := compare{
		vmBytes:       vmBytes,
		resultBuilder: resultBuilder,
		trueBytes:     trueBytes,
		falseBytes:    falseBytes,
	}

	return &out
}

// Execute executes the comparison
func (app *compare) Execute(input bytes_programs.Programs, frame frames.Frame) (results.Result, error) {
	exec, err := app.vmBytes.Programs(input, frame.Bytes())
	if err != nil {
		return nil, err
	}

	if exec.HasFailure() {
		return app.resultBuilder.Create().
			WithFailure(exec).
			Now()
	}

	isEqual := true
	previous := []byte{}
	execList := exec.List()
	for _, oneExec := range execList {
		value := oneExec.Success()
		if len(previous) <= 0 {
			previous = value
			continue
		}

		if !bytes.Equal(previous, value) {
			isEqual = false
			break
		}
	}

	builder := app.resultBuilder.Create().WithSuccess(app.trueBytes)
	if !isEqual {
		builder.WithSuccess(app.falseBytes)
	}

	return builder.Now()
}
