package queries

import (
	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/data/connections/links/contexts"
)

// Builder represents the queries builder
type Builder interface {
	Create() Builder
	WithList(list []Query) Builder
	Now() (Queries, error)
}

// Queries represents queries
type Queries interface {
	List() []Query
}

// QueryBuilder represents the query builder
type QueryBuilder interface {
	Create() QueryBuilder
	WithContexts(contexts contexts.Contexts) QueryBuilder
	From(from uuid.UUID) QueryBuilder
	To(to uuid.UUID) QueryBuilder
	Now() (Query, error)
}

// Query represents the query
type Query interface {
	Contexts() contexts.Contexts
	From() uuid.UUID
	To() uuid.UUID
}
