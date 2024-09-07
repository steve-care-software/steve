package activities

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits"
)

type activity struct {
	hash    hash.Hash
	commits commits.Commits
	head    hash.Hash
}

func createActivity(
	hash hash.Hash,
	commits commits.Commits,
	head hash.Hash,
) Activity {
	out := activity{
		hash:    hash,
		commits: commits,
		head:    head,
	}

	return &out
}

// Hash returns the hash
func (obj *activity) Hash() hash.Hash {
	return obj.hash
}

// Commits returns the commits
func (obj *activity) Commits() commits.Commits {
	return obj.commits
}

// Head returns the head
func (obj *activity) Head() hash.Hash {
	return obj.head
}
