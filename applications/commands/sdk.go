package commands

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions"
)

// Application represents the command application
type Application interface {
	Execute(input []byte) (executions.Execution, error)
}
