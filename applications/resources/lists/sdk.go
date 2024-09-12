package lists

import (
	"github.com/steve-care-software/steve/applications/resources"
	"github.com/steve-care-software/steve/domain/stores/lists"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	listAdapter := lists.NewAdapter()
	return createBuilder(
		listAdapter,
	)
}

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithResource(resource resources.Application) Builder
	Now() (Application, error)
}

// Application represents the list application
type Application interface {
	Amount(name string) (*uint, error)
	Retrieve(name string, index uint, amount uint) ([][]byte, error)
	RetrieveAll(name string) ([][]byte, error)
	Append(name string, values [][]byte) error
}
