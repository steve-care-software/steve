package bytes

import (
	"github.com/steve-care-software/steve/vms/bytes/frames"
	"github.com/steve-care-software/steve/vms/bytes/programs"
	"github.com/steve-care-software/steve/vms/bytes/results"
)

type bytes struct {
	resultsBuilder results.Builder
	resultBuilder  results.ResultBuilder
	failureBuilder results.FailureBuilder
}

func createBytes(
	resultsBuilder results.Builder,
	resultBuilder results.ResultBuilder,
	failureBuilder results.FailureBuilder,
) Bytes {
	out := bytes{
		resultsBuilder: resultsBuilder,
		resultBuilder:  resultBuilder,
		failureBuilder: failureBuilder,
	}

	return &out
}

// Programs execute programs
func (app *bytes) Programs(programs programs.Programs, frame frames.Frame) (results.Results, error) {
	list := []results.Result{}
	programsList := programs.List()
	for _, oneProgram := range programsList {
		ins, err := app.Program(oneProgram, frame)
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
		if ins.IsFailure() {
			break
		}
	}

	return app.resultsBuilder.Create().
		WithList(list).
		Now()
}

// Program execute program
func (app *bytes) Program(program programs.Program, frame frames.Frame) (results.Result, error) {
	builder := app.resultBuilder.Create()
	if program.IsVariable() {
		variable := program.Variable()
		value, err := frame.Fetch(variable)
		if err != nil {
			failure, err := app.failureBuilder.Create().
				WithUndefined(variable).
				Now()

			if err != nil {
				return nil, err
			}

			return builder.WithFailure(failure).
				Now()
		}

		builder.WithSuccess(value)
	}

	if program.IsValue() {
		value := program.Value()
		builder.WithSuccess(value)
	}

	return builder.Now()
}
