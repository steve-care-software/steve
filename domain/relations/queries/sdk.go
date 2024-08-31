package queries

import (
	"github.com/steve-care-software/steve/domain/relations/data/connections/contexts"
	"github.com/steve-care-software/steve/domain/relations/data/points"
	"github.com/steve-care-software/steve/domain/relations/queries/questions"
)

// Query represents a query
type Query interface {
	Contexts() contexts.Contexts
	Question() questions.Question
	HasPoints() bool
	Points() points.Points
}
