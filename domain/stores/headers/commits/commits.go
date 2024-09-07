package commits

import (
	"github.com/steve-care-software/steve/domain/hash"
)

type commits struct {
	hash hash.Hash
	list []Commit
}

func createCommits(
	hash hash.Hash,
	list []Commit,
) Commits {
	out := commits{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *commits) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *commits) List() []Commit {
	return obj.list
}
