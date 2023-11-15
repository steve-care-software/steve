package transformations

import (
	bytes_vm "github.com/steve-care-software/steve/vms/bytes"
	bytes_programs "github.com/steve-care-software/steve/vms/bytes/programs"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/frames"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/scopes/assignables/scopes/transformations/results"
)

type join struct {
	vmBytes       bytes_vm.Bytes
	resultBuilder results.Builder
}

func createJoin(
	vmBytes bytes_vm.Bytes,
	resultBuilder results.Builder,
) Transformation {
	out := join{
		vmBytes:       vmBytes,
		resultBuilder: resultBuilder,
	}

	return &out
}

// Execute executes the join
func (app *join) Execute(input bytes_programs.Programs, frame frames.Frame) (results.Result, error) {
	exec, err := app.vmBytes.Programs(input, frame.Bytes())
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
