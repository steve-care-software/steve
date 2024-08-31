package applications

import (
	"github.com/steve-care-software/steve/domain/relations/answers"
	"github.com/steve-care-software/steve/domain/relations/data/points"
	"github.com/steve-care-software/steve/domain/relations/queries"
)

// Factory represents the application factory
type Factory interface {
	Create() (Application, error)
}

// InMemoryBuilder represents the in-memory builder
type InMemoryBuilder interface {
	Create() InMemoryBuilder
	WithPoints(points points.Points) InMemoryBuilder
	Now() (Application, error)
}

// Application represents the application
type Application interface {
	Execute(query queries.Query) (answers.Answer, error)
}
