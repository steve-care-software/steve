package layers

import (
	"github.com/steve-care-software/steve/vms/layers/frames"
	"github.com/steve-care-software/steve/vms/layers/programs"
	"github.com/steve-care-software/steve/vms/layers/results"
	vm_instructions "github.com/steve-care-software/steve/vms/layers/scopes/instructions"
	instruction_frames "github.com/steve-care-software/steve/vms/layers/scopes/instructions/frames"
)

type layer struct {
	vmInstructions          vm_instructions.Instruction
	instructionFrameBuilder instruction_frames.FrameBuilder
}

func createLayer(
	vmInstructions vm_instructions.Instruction,
	instructionFrameBuilder instruction_frames.FrameBuilder,
) Layer {
	out := layer{
		vmInstructions:          vmInstructions,
		instructionFrameBuilder: instructionFrameBuilder,
	}

	return &out
}

// Execute executes the layer
func (app *layer) Execute(program programs.Program, frame frames.Frame) (results.Result, error) {
	return nil, nil
}
