package results

import (
	bytes_results "github.com/steve-care-software/steve/vms/bytes/results"
	"github.com/steve-care-software/steve/vms/bytes/results/hash"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/instructions/frames"
	assignable_results "github.com/steve-care-software/steve/vms/queries/scopes/layers/instructions/scopes/assignments/scopes/assignables/results"
)

// Builder represents the result builder
type Builder interface {
	Create() Builder
	WithQuery(query Query) Builder
	WithResult(result frames.Frames) Builder
	Now() (Result, error)
}

// Result represens results
type Result interface {
	Query() Query
	Result() frames.Frames
}

// QueryBuilder represents the query builder
type QueryBuilder interface {
	Create() QueryBuilder
	WithLayer(layer assignable_results.Result) QueryBuilder
	WithBytes(bytes bytes_results.Result) QueryBuilder
	WithParams(params bytes_results.Results) QueryBuilder
	WithDependencies(dependencies assignable_results.Results) QueryBuilder
	Now() (Query, error)
}

// Query represents the query
type Query interface {
	Hash() hash.Hash
	Layer() assignable_results.Result
	Bytes() bytes_results.Result
	Params() bytes_results.Results
	HasDependencies() bool
	Dependencies() assignable_results.Results
}

// FetchLayer represents a fetch layer result
type FetchLayer interface {
	IsSuccess() bool
	Success() string
}
