package deletes

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/administrators/instances/successes/deletes/failures"
)

// Builder represents a delete builder
type Builder interface {
	Create() Builder
	WithSuccess(admin administrators.Administrator) Builder
	WithFailure(failure failures.Failure) Builder
	Now() (Delete, error)
}

// Delete represents a delete administration
type Delete interface {
	IsSuccess() bool
	Success() administrators.Administrator
	IsFailure() bool
	Failure() failures.Failure
}
