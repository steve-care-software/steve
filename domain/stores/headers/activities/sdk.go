package activities

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/commits"
)

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
