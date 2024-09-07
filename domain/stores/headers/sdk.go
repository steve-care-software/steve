package headers

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/commits"
	"github.com/steve-care-software/steve/domain/stores/headers/commits/modifications/resources"
)

// Builder represents the header builder
type Builder interface {
	Create() Builder
	WithRoot(root resources.Resources) Builder
	WithCommits(commits commits.Commits) Builder
	WithHead(head hash.Hash) Builder
	Now() (Header, error)
}

// Header represents the header
type Header interface {
	Hash() hash.Hash
	Root() resources.Resources
	Commits() commits.Commits
	Head() hash.Hash
}
