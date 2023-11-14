package instructions

import (
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/frames"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/programs"
)

// Instruction represents an instruction
type Instruction interface {
	Instructions(programs programs.Programs, frame frames.Frame) (frames.Frames, error)
	Instruction(program programs.Program, frame frames.Frame) (frames.Frame, error)
	Condition(program programs.Condition, frame frames.Frame) (frames.Block, error)
}
