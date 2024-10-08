package commits

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/hash"
)

type commits struct {
	hash hash.Hash
	list []Commit
	mp   map[string]Commit
}

func createCommits(
	hash hash.Hash,
	list []Commit,
	mp map[string]Commit,
) Commits {
	out := commits{
		hash: hash,
		list: list,
		mp:   mp,
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

// Fetch fetches a commit by hash
func (obj *commits) Fetch(hash hash.Hash) (Commit, error) {
	if ins, ok := obj.mp[hash.String()]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the commit (hash: %s) could not be found", hash.String())
	return nil, errors.New(str)
}
