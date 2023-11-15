package results

import (
	reduce_results "github.com/steve-care-software/steve/vms/children/layers/scopes/instructions/scopes/assignments/scopes/assignables/scopes/reduces/results"
	transformation_results "github.com/steve-care-software/steve/vms/children/layers/scopes/instructions/scopes/assignments/scopes/assignables/scopes/transformations/results"
	"github.com/steve-care-software/steve/vms/libraries/hash"
)

// Results represents results
type Results interface {
	Hash() hash.Hash
	List() []Result
}

// ResultBuilder represents the result builder
type ResultBuilder interface {
	Create() ResultBuilder
	WithCompare(compare transformation_results.Result) ResultBuilder
	WithLength(length transformation_results.Result) ResultBuilder
	WithJoin(join transformation_results.Result) ResultBuilder
	WithReduce(reduce reduce_results.Result) ResultBuilder
	Now() (Result, error)
}

// Result represents an assignable result
type Result interface {
	Hash() hash.Hash
	IsCompare() bool
	Compare() transformation_results.Result
	IsLength() bool
	Length() transformation_results.Result
	IsJoin() bool
	Join() transformation_results.Result
	IsReduce() bool
	Reduce() reduce_results.Result
}
