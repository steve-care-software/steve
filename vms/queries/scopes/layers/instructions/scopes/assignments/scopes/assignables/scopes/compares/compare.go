package compares

import (
	bytes_vm "github.com/steve-care-software/steve/vms/bytes"
	bytes_programs "github.com/steve-care-software/steve/vms/bytes/programs"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/instructions/scopes/assignments/frames"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/instructions/scopes/assignments/scopes/assignables/scopes/compares/results"
	result_compares "github.com/steve-care-software/steve/vms/queries/scopes/layers/instructions/scopes/assignments/scopes/assignables/scopes/compares/results"
)

type compare struct {
	bytes          bytes_vm.Bytes
	resultBuilder  result_compares.Builder
	successBuilder result_compares.SuccessBuilder
	failureBuilder result_compares.FailureBuilder
}

func createCompare(
	bytes bytes_vm.Bytes,
	resultBuilder result_compares.Builder,
	successBuilder result_compares.SuccessBuilder,
	failureBuilder result_compares.FailureBuilder,
) Compare {
	out := compare{
		bytes:          bytes,
		resultBuilder:  resultBuilder,
		successBuilder: successBuilder,
		failureBuilder: failureBuilder,
	}

	return &out
}

// Execute executes the comparison
func (app *compare) Execute(input bytes_programs.Programs, frame frames.Frame) (results.Result, error) {
	exec, err := app.bytes.Programs(input, frame.Bytes())
	if err != nil {
		return nil, err
	}

	if !exec.IsSuccess() {
		failure, err := app.failureBuilder.Create().
			WithBytesFailed(exec).
			Now()

		if err != nil {
			return nil, err
		}

		return app.resultBuilder.Create().
			WithFailure(failure).
			Now()
	}

	success, err := app.successBuilder.Create().
		WithBytesResults(exec).
		Now()

	if err != nil {
		return nil, err
	}

	return app.resultBuilder.Create().
		WithSuccess(success).
		Now()
}
