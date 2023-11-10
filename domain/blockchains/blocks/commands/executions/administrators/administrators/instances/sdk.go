package instances

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/administrators/instances/failures"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/administrators/instances/successes"
)

// Builder represents an instance builder
type Builder interface {
	Create() Builder
	WithSuccess(success successes.Success) Builder
	WithFailure(failure failures.Failure) Builder
	Now() (Instance, error)
}

// Instance represents an instance
type Instance interface {
	IsSuccess() bool
	Success() successes.Success
	IsFailure() bool
	Failure() failures.Failure
}
