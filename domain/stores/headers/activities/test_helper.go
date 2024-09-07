package activities

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits"
)

// NewActivityForTests creates a new activity for tests
func NewActivityForTests(commits commits.Commits, head hash.Hash) Activity {
	ins, err := NewBuilder().Create().WithCommits(commits).WithHead(head).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
