package identities

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/identities/failures"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/identities/successes"
)

// Builder represents an identities builder
type Builder interface {
	Create() Builder
	WithSuccess(success successes.Success) Builder
	WithFailure(failure failures.Failure) Builder
	Now() (Identities, error)
}

// Identities represents identities
type Identities interface {
	IsSuccess() bool
	Success() successes.Success
	IsFailure() bool
	Failure() failures.Failure
}
