package identities

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/identities/identities"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/identities/identities"
)

// Application represents the application
type Application interface {
	Execute(identity inputs.Identity, frame frames.Frame) (executions.Identity, error)
}
