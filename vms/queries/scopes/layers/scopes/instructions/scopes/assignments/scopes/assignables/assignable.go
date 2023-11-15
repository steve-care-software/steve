package assignables

import (
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/frames"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/scopes/assignables/programs"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/scopes/assignables/results"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/scopes/assignables/scopes/reduces"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/scopes/assignables/scopes/transformations"
)

type assignable struct {
	compare       transformations.Transformation
	join          transformations.Transformation
	length        transformations.Transformation
	reduce        reduces.Reduce
	resultBuilder results.ResultBuilder
}

func createAssignable(
	compare transformations.Transformation,
	join transformations.Transformation,
	length transformations.Transformation,
	reduce reduces.Reduce,
	resultBuilder results.ResultBuilder,
) Assignable {
	out := assignable{
		compare:       compare,
		join:          join,
		length:        length,
		reduce:        reduce,
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

	if input.IsLength() {
		length := input.Length()
		exec, err := app.length.Execute(length, frame)
		if err != nil {
			return nil, err
		}

		builder.WithLength(exec)
	}

	if input.IsJoin() {
		join := input.Join()
		exec, err := app.join.Execute(join, frame)
		if err != nil {
			return nil, err
		}

		builder.WithJoin(exec)
	}

	if input.IsReduce() {
		reduce := input.Reduce()
		exec, err := app.reduce.Execute(reduce, frame)
		if err != nil {
			return nil, err
		}

		builder.WithReduce(exec)
	}

	return builder.Now()
}
