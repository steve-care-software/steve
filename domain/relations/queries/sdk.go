package queries

import (
	"github.com/steve-care-software/steve/domain/relations/data/connections/links/contexts"
	"github.com/steve-care-software/steve/domain/relations/data/points"
	"github.com/steve-care-software/steve/domain/relations/queries/questions"
)

// Builder represents the query builder
type Builder interface {
	Create() Builder
	WithContexts(contexts contexts.Contexts) Builder
	WithQuestion(question questions.Question) Builder
	WithPoints(points points.Points) Builder
	Now() (Query, error)
}

// Query represents a query
type Query interface {
	Contexts() contexts.Contexts
	Question() questions.Question
	HasPoints() bool
	Points() points.Points
}
