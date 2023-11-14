package results

import (
	"github.com/steve-care-software/steve/vms/bytes/results/hash"
	compare_results "github.com/steve-care-software/steve/vms/queries/scopes/layers/instructions/scopes/assignments/scopes/assignables/scopes/compares/results"
)

// Results represents results
type Results interface {
	Hash() hash.Hash
	List() []Result
}

// ResultBuilder represents the result builder
type ResultBuilder interface {
	Create() ResultBuilder
	WithCompare(compare compare_results.Result) ResultBuilder
	Now() (Result, error)
}

// Result represents an assignable result
type Result interface {
	Hash() hash.Hash
	IsCompare() bool
	Compare() compare_results.Result
}
