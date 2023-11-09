package deletes

import (
	"github.com/steve-care-software/steve/domain/accounts/identities"
	"github.com/steve-care-software/steve/domain/commands/executions/identities/identities/instances/successes/deletes/failures"
)

// Builder represents a delete builder
type Builder interface {
	Create() Builder
	WithSuccess(admin identities.Identity) Builder
	WithFailure(failure failures.Failure) Builder
	Now() (Delete, error)
}

// Delete represents a delete administration
type Delete interface {
	IsSuccess() bool
	Success() identities.Identity
	IsFailure() bool
	Failure() failures.Failure
}
