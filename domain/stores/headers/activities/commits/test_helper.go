package commits

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications"
)

// NewCommitsForTests creates a new commits for tests
func NewCommitsForTests(list []Commit) Commits {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewCommitForTests creates a new commit for tests
func NewCommitForTests(modifications modifications.Modifications) Commit {
	ins, err := NewCommitBuilder().Create().WithModifications(modifications).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewCommitWithParentForTests creates a new commit with parent for tests
func NewCommitWithParentForTests(modifications modifications.Modifications, parent hash.Hash) Commit {
	ins, err := NewCommitBuilder().Create().WithModifications(modifications).WithParent(parent).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
