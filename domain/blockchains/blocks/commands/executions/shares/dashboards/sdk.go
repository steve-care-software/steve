package dashboards

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/shares/dashboards/failures"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/shares/dashboards/successes"
)

// Builder represents a dashboard builder
type Builder interface {
	Create() Builder
	WithSuccess(success successes.Success) Builder
	WithFailure(failure failures.Failure) Builder
	Now() (Dashboard, error)
}

// Dashboard represents a dashboard
type Dashboard interface {
	IsSuccess() bool
	Success() successes.Success
	IsFailure() bool
	Failure() failures.Failure
}
