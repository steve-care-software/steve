package results

import bytes_results "github.com/steve-care-software/steve/vms/bytes/results"

// Builder represents a result builder
type Builder interface {
	Create() Builder
	WithSuccess(success []byte) Builder
	WithFailure(failure Failure) Builder
	Now() (Result, error)
}

// Result represents a result
type Result interface {
	IsSuccess() bool
	Success() []byte
	IsFailure() bool
	Failure() Failure
}

// FailureBuilder represents a failure builder
type FailureBuilder interface {
	Create() FailureBuilder
	WithBytesFailure(bytes bytes_results.Result) FailureBuilder
	WithNotEnoughBytes(requested uint) FailureBuilder
	Now() (Failure, error)
}

// Failure represents a failure
type Failure interface {
	IsBytesFailure() bool
	BytesFailure() bytes_results.Result
	IsNotEnoughBytes() bool
	NotEnoughBytes() *uint
}
