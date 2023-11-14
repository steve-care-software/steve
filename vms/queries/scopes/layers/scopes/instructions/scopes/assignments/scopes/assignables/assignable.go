package assignables

import (
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/frames"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/scopes/assignables/programs"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/scopes/assignables/results"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/scopes/assignables/scopes/compares"
)

type assignable struct {
	compare       compares.Compare
	resultBuilder results.ResultBuilder
}

func createAssignable(
	compare compares.Compare,
	resultBuilder results.ResultBuilder,
) Assignable {
	out := assignable{
		compare:       compare,
		resultBuilder: resultBuilder,
	}

	return &out
}

// Execute executes an assignable
func (app *assignable) Execute(input programs.Program, frame frames.Frame) (results.Result, error) {
	builder := app.resultBuilder.Create()
	if input.IsCompare() {
		compare := input.Compare()
		exec, err := app.compare.Execute(compare, frame)
		if err != nil {
			return nil, err
		}

		builder.WithCompare(exec)
	}

	return builder.Now()
}
