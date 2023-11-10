package deletes

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators/identities"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/identities/successes/deletes/failures"
)

// Builder represents a delete builder
type Builder interface {
	Create() Builder
	WithSuccess(success identities.Identities) Builder
	WithFailure(failure failures.Failure) Builder
	Now() (Delete, error)
}

// Delete repreents a delete
type Delete interface {
	IsSuccess() bool
	Success() identities.Identities
	IsFailure() bool
	Failure() failures.Failure
}
