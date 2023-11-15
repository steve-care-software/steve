package instructions

import (
	"bytes"

	bytes_applications "github.com/steve-care-software/steve/vms/bytes/applications"
	"github.com/steve-care-software/steve/vms/layers/scopes/instructions/frames"
	"github.com/steve-care-software/steve/vms/layers/scopes/instructions/programs"
	"github.com/steve-care-software/steve/vms/layers/scopes/instructions/scopes/assignments"
)

type instruction struct {
	bytesApp            bytes_applications.Application
	vmAssignment        assignments.Assignment
	framesBuilder       frames.Builder
	frameBuilder        frames.FrameBuilder
	entryBuilder        frames.EntryBuilder
	blockBuilder        frames.BlockBuilder
	blockFailureBuilder frames.BlockFailureBuilder
	trueBytes           []byte
}

func createInstruction(
	bytesApp bytes_applications.Application,
	vmAssignment assignments.Assignment,
	framesBuilder frames.Builder,
	frameBuilder frames.FrameBuilder,
	entryBuilder frames.EntryBuilder,
	blockBuilder frames.BlockBuilder,
	blockFailureBuilder frames.BlockFailureBuilder,
	trueBytes []byte,
) Instruction {
	out := instruction{
		bytesApp:            bytesApp,
		vmAssignment:        vmAssignment,
		framesBuilder:       framesBuilder,
		frameBuilder:        frameBuilder,
		entryBuilder:        entryBuilder,
		blockBuilder:        blockBuilder,
		blockFailureBuilder: blockFailureBuilder,
		trueBytes:           trueBytes,
	}

	return &out
}

// Instructions executes the instructions
func (app *instruction) Instructions(programs programs.Programs, frame frames.Frame) (frames.Frames, error) {
	framesList := []frames.Frame{}
	list := programs.List()
	for _, oneProgram := range list {
		frame, err := app.Instruction(oneProgram, frame)
		if err != nil {
			return nil, err
		}

		framesList = append(framesList, frame)
	}

	return app.framesBuilder.Create().
		WithList(framesList).
		Now()
}

// Instruction executes the instruction
func (app *instruction) Instruction(program programs.Program, frame frames.Frame) (frames.Frame, error) {
	if program.IsStop() {
		return frame, nil
	}

	builder := app.frameBuilder.Create()
	if program.IsAssignment() {
		assignment := program.Assignment()
		exec, err := app.vmAssignment.Execute(assignment, frame.Assignment())
		if err != nil {
			return nil, err
		}

		entry, err := app.entryBuilder.Create().
			WithFrame(exec).
			Now()

		if err != nil {
			return nil, err
		}

		list := frame.List()
		list = append(list, entry)
		builder.WithList(list)
	}

	if program.IsCondition() {
		condition := program.Condition()
		exec, err := app.Condition(condition, frame)
		if err != nil {
			return nil, err
		}

		entry, err := app.entryBuilder.Create().
			WithBlock(exec).
			Now()

		if err != nil {
			return nil, err
		}

		list := frame.List()
		if entry != nil {
			list = append(list, entry)
		}

		builder.WithList(list)
	}

	return builder.Now()
}

// Condition executes the condition
func (app *instruction) Condition(program programs.Condition, frame frames.Frame) (frames.Block, error) {
	constraint := program.Constraint()
	execCons, err := app.bytesApp.Program(constraint, frame.Bytes())
	if err != nil {
		return nil, err
	}

	if execCons.IsFailure() {
		failure, err := app.blockFailureBuilder.Create().
			WithConditionFailed(execCons).
			Now()

		if err != nil {
			return nil, err
		}

		return app.blockBuilder.Create().
			WithFailure(failure).
			Now()
	}

	success := execCons.Success()
	if !bytes.Equal(success, app.trueBytes) {
		return nil, nil
	}

	programs := program.Programs()
	frames, err := app.Instructions(programs, frame)
	if err != nil {
		return nil, err
	}

	return app.blockBuilder.Create().
		WithSuccess(frames).
		Now()
}
