package visitors

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/visitors/failures"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/visitors/successes"
)

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithSuccess(success successes.Success) Builder
	WithFailure(failure failures.Failure) Builder
	Now() (Visitor, error)
}

// Visitor represents a visitor
type Visitor interface {
	IsSuccess() bool
	Success() successes.Success
	IsFailure() bool
	Failure() failures.Failure
}
