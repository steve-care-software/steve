package activities

import (
	"github.com/steve-care-software/steve/hash"
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits"
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits/modifications/resources/pointers"
)

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	commitsAdapter := commits.NewAdapter()
	hashAdapter := hash.NewAdapter()
	builder := NewBuilder()
	return createAdapter(
		commitsAdapter,
		hashAdapter,
		builder,
	)
}

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
	ToInstance(data []byte) (Activity, []byte, error)
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
	Map(rootHash hash.Hash) (map[string]pointers.Pointer, []string, error)
}
