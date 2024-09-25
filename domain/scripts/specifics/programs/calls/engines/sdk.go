package engines

import (
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/calls/functions"
)

// Engine represents an engine call
type Engine interface {
	Scope() uint8 // role, identity, etc
	FuncCall() functions.Function
}
