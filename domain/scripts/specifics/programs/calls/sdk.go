package calls

import (
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/calls/engines"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/calls/functions"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/calls/programs"
)

// Call represents a call
type Call interface {
	IsProgram() bool
	Program() programs.Program
	IsEngine() bool
	Engine() engines.Engine
	IsFunc() bool
	Func() functions.Function
}
