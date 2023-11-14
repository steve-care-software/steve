package results

import (
	"github.com/steve-care-software/steve/vms/bytes/results"
	"github.com/steve-care-software/steve/vms/bytes/results/hash"
)

// Builder represents a compare builder
type Builder interface {
	Create() Builder
	WithSuccess(success Success) Builder
	WithFailure(failure Failure) Builder
	Now() (Result, error)
}

// Result represents a compare results
type Result interface {
	Hash() hash.Hash
	IsSuccess() bool
	Success() Success
	IsFailure() bool
	Failure() Failure
}

// SuccessBuilder represents a success builder
type SuccessBuilder interface {
	Create() SuccessBuilder
	WithBytesResults(bytesResults results.Results) SuccessBuilder
	Now() (Success, error)
}

// Success represents a success
type Success interface {
	Hash() hash.Hash
	Bytes() results.Results
	IsSame() bool
}

// FailureBuilder represents a failure builder
type FailureBuilder interface {
	Create() FailureBuilder
	WithBytesFailed(bytesFailed results.Results) FailureBuilder
	Now() (Failure, error)
}

// Failure represents a failure
type Failure interface {
	Hash() hash.Hash
	IsBytesFailed() bool
	BytesFailed() results.Results
}
