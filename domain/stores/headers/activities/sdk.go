package activities

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the activity adapter
type Adapter interface {
	ToBytes(ins Activity) ([]byte, error)
	ToInstance(data []byte) (Activity, error)
}

// Builder represents the header builder
type Builder interface {
	Create() Builder
	WithCommits(commits commits.Commits) Builder
	WithHead(head hash.Hash) Builder
	Now() (Activity, error)
}

// Activity represents the activity
type Activity interface {
	Hash() hash.Hash
	Commits() commits.Commits
	Head() hash.Hash
}