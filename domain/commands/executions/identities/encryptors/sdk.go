package encryptors

import (
	"github.com/steve-care-software/steve/domain/commands/executions/identities/encryptors/failures"
	"github.com/steve-care-software/steve/domain/commands/executions/identities/encryptors/successes"
)

// Builder represents the encryptor builder
type Builder interface {
	Create() Builder
	WithSuccess(success successes.Success) Builder
	WithFailure(failure failures.Failure) Builder
	Now() (Encryptor, error)
}

// Encryptor represents an encryptor
type Encryptor interface {
	IsSuccess() bool
	Success() successes.Success
	IsFailure() bool
	Failure() failures.Failure
}
