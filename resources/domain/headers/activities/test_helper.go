package activities

import (
	"github.com/steve-care-software/steve/hash"
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits"
)

// NewActivityForTests creates a new activity for tests
func NewActivityForTests(commits commits.Commits, head hash.Hash) Activity {
	ins, err := NewBuilder().Create().WithCommits(commits).WithHead(head).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
