package profiles

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/profiles"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/profiles"
	"github.com/steve-care-software/steve/domain/stacks"
)

// Application represents the application
type Application interface {
	Execute(profile inputs.Profile, stack stacks.Stack) (executions.Profile, error)
}
