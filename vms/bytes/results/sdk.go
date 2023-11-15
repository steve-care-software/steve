package results

import (
	"github.com/steve-care-software/steve/vms/bytes/results/hash"
)

// Builder represents the results builder
type Builder interface {
	Create() Builder
	WithList(list []Result) Builder
	Now() (Results, error)
}

// Results represents results
type Results interface {
	Hash() hash.Hash
	List() []Result
	HasFailure() bool
	Failure() Failure
}

// ResultBuilder represents the result builder
type ResultBuilder interface {
	Create() ResultBuilder
	WithSuccess(success []byte) ResultBuilder
	WithFailure(failure Failure) ResultBuilder
	Now() (Result, error)
}

// Result represents result
type Result interface {
	Hash() hash.Hash
	IsSuccess() bool
	Success() []byte
	IsFailure() bool
	Failure() Failure
}

// FailureBuilder represents the failure builder
type FailureBuilder interface {
	Create() FailureBuilder
	WithUndefined(undefined string) FailureBuilder
	Now() (Failure, error)
}

// Failure represents failure
type Failure interface {
	Hash() hash.Hash
	IsUndefined() bool
	Undefined() string
}
