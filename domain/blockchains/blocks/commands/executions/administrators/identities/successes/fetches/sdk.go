package fetches

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/identities/successes/fetches/failures"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/identities/successes/fetches/successes"
)

// Builder represents a fetch builder
type Builder interface {
	Create() Builder
	WithSuccess(success successes.Success) Builder
	WithFailure(failure failures.Failure) Builder
	Now() (Fetch, error)
}

// Fetch represents a fetch
type Fetch interface {
	IsSuccess() bool
	Success() successes.Success
	IsFailure() bool
	Failure() failures.Failure
}
