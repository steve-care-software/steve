package compares

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/executions/instructions/assignments/assignables/compares/failures"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/frames/assignables"
)

// Builder represents the compare builder
type Builder interface {
	Create() Builder
	WithSuccess(success assignables.Assignable) Builder
	WithFailure(failure failures.Failure) Builder
	Now() (Compare, error)
}

// Compare represents a compare
type Compare interface {
	IsSuccess() bool
	Success() assignables.Assignable
	IsFailure() bool
	Failure() failures.Failure
}
