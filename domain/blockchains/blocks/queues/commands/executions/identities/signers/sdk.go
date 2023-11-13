package signers

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/identities/signers/failures"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/identities/signers/successes"
)

// Signer represents a signer
type Signer interface {
	IsSuccess() bool
	Success() successes.Success
	IsFailure() bool
	Failure() failures.Failure
}
