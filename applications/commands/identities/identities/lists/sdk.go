package lists

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/identities/identities/lists"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/identities/identities/lists"
)

// Application represents the application
type Application interface {
	Execute(list inputs.List, frame frames.Frame) (executions.List, error)
}
