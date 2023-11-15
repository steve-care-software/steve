package assignments

import (
	"github.com/steve-care-software/steve/vms/layers/scopes/instructions/scopes/assignments/frames"
	"github.com/steve-care-software/steve/vms/layers/scopes/instructions/scopes/assignments/programs"
	"github.com/steve-care-software/steve/vms/layers/scopes/instructions/scopes/assignments/scopes/assignables"
)

type assignment struct {
	assignable        assignables.Assignable
	frameBuilder      frames.FrameBuilder
	assignmentBuilder frames.AssignmentBuilder
}

func createAssignment(
	assignable assignables.Assignable,
	frameBuilder frames.FrameBuilder,
	assignmentBuilder frames.AssignmentBuilder,
) Assignment {
	out := assignment{
		assignable:        assignable,
		frameBuilder:      frameBuilder,
		assignmentBuilder: assignmentBuilder,
	}

	return &out
}

// Execute executes an assignment
func (app *assignment) Execute(input programs.Program, frame frames.Frame) (frames.Frame, error) {
	assignable := input.Assignable()
	exec, err := app.assignable.Execute(assignable, frame)
	if err != nil {
		return nil, err
	}

	name := input.Name()
	assignment, err := app.assignmentBuilder.Create().
		WithName(name).
		WithValue(exec).
		Now()

	if err != nil {
		return nil, err
	}

	list := frame.List()
	list = append(list, assignment)
	return app.frameBuilder.Create().
		WithList(list).
		Now()
}
