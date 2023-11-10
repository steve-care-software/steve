package identities

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/identities"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/administrators/identities"
)

// Application represents the identities application
type Application interface {
	Execute(instance inputs.Identities, frame frames.Frame) (executions.Identities, error)
}
