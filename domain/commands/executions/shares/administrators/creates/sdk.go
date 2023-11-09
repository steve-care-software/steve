package creates

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	"github.com/steve-care-software/steve/domain/commands/executions/shares/administrators/creates/failures"
)

// Builder represents a create builder
type Builder interface {
	Create() Builder
	WithSuccess(success administrators.Administrator) Builder
	WithFailure(failure failures.Failure) Builder
	Now() (Create, error)
}

// Create represents a create instance
type Create interface {
	IsSuccess() bool
	Success() administrators.Administrator
	IsFailure() bool
	Failure() failures.Failure
}
