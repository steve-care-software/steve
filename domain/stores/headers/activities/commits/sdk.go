package commits

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications"
)

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	modificationsAdapter := modifications.NewAdapter()
	hashAdapter := hash.NewAdapter()
	builder := NewBuilder()
	commitBuilder := NewCommitBuilder()
	return createAdapter(
		modificationsAdapter,
		hashAdapter,
		builder,
		commitBuilder,
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewCommitBuilder creates a new builder
func NewCommitBuilder() CommitBuilder {
	hashAdapter := hash.NewAdapter()
	return createCommitBuilder(
		hashAdapter,
	)
}

// Adapter represents the commits adapter
type Adapter interface {
	InstancesToBytes(ins Commits) ([]byte, error)
	BytesToInstances(data []byte) (Commits, []byte, error)
	InstanceToBytes(ins Commit) ([]byte, error)
	BytesToInstance(data []byte) (Commit, []byte, error)
}

// Builder represents the commits builder
type Builder interface {
	Create() Builder
	WithList(list []Commit) Builder
	Now() (Commits, error)
}

// Commits represents commits
type Commits interface {
	Hash() hash.Hash
	List() []Commit
	Fetch(hash hash.Hash) (Commit, error)
}

// CommitBuilder represents a commit builder
type CommitBuilder interface {
	Create() CommitBuilder
	WithModifications(modifications modifications.Modifications) CommitBuilder
	WithParent(parent hash.Hash) CommitBuilder
	Now() (Commit, error)
}

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Modifications() modifications.Modifications
	HasParent() bool
	Parent() hash.Hash
}
