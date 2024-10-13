package instructions

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/assignments/assignables"
)

type returnInstructionBuilder struct {
	assignable assignables.Assignable
}

func createReturnInstructionBuilder() ReturnInstructionBuilder {
	return &returnInstructionBuilder{
		assignable: nil,
	}
}

func (app *returnInstructionBuilder) Create() ReturnInstructionBuilder {
	return createReturnInstructionBuilder()
}

func (app *returnInstructionBuilder) WithAssignable(assignable assignables.Assignable) ReturnInstructionBuilder {
	app.assignable = assignable
	return app
}

func (app *returnInstructionBuilder) Now() (ReturnInstruction, error) {
	if app.assignable != nil {
		return createReturnInstructionWithAssignable(app.assignable), nil
	}

	return createReturnInstruction(), nil
}
