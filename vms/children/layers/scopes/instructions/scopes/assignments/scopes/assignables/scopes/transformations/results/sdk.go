package results

import (
	bytes_results "github.com/steve-care-software/steve/vms/children/bytes/results"
	"github.com/steve-care-software/steve/vms/libraries/hash"
)

// Builder represents the result builder
type Builder interface {
	Create() Builder
	WithSuccess(success []byte) Builder
	WithFailure(failure bytes_results.Results) Builder
	Now() (Result, error)
}

// Result represents the result
type Result interface {
	Hash() hash.Hash
	IsSuccess() bool
	Success() []byte
	IsFailure() bool
	Failure() bytes_results.Results
}
