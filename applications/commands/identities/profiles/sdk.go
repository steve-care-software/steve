package profiles

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/identities/profiles"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/identities/profiles"
)

// Application represents the application
type Application interface {
	Execute(profile inputs.Profile, frame frames.Frame) (executions.Profile, error)
}
