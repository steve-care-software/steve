package commits

import "github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources"

// Builder represents a commits builder
type Builder interface {
	Create() Builder
	WithList(list []Commit) Builder
	Now() (Commits, error)
}

// Commits represents commits
type Commits interface {
	List() []Commit
}

// CommitBuilder represents a commit builder
type CommitBuilder interface {
	Create() CommitBuilder
	WithMessage(msg string) CommitBuilder
	WithResources(resources resources.Resources) CommitBuilder
	Now() (Commit, error)
}

// Commit represents a commit
type Commit interface {
	Message() string
	Resources() resources.Resources
}
