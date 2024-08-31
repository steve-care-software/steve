package queries

import (
	"github.com/google/uuid"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewQueryBuilder creates a new query builder
func NewQueryBuilder() QueryBuilder {
	return createQueryBuilder()
}

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
	From(from uuid.UUID) QueryBuilder
	To(to uuid.UUID) QueryBuilder
	Now() (Query, error)
}

// Query represents the query
type Query interface {
	From() uuid.UUID
	To() uuid.UUID
}
