package deletes

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	"github.com/steve-care-software/steve/domain/commands/executions/administrators/instances/deletes/failures"
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
