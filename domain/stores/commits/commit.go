package commits

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/commits/modifications"
)

type commit struct {
	hash          hash.Hash
	modifications modifications.Modifications
	parent        hash.Hash
}

func createCommit(
	hash hash.Hash,
	modifications modifications.Modifications,
) Commit {
	return createCommitInternally(hash, modifications, nil)
}

func createCommitWithParent(
	hash hash.Hash,
	modifications modifications.Modifications,
	parent hash.Hash,
) Commit {
	return createCommitInternally(hash, modifications, parent)
}

func createCommitInternally(
	hash hash.Hash,
	modifications modifications.Modifications,
	parent hash.Hash,
) Commit {
	out := commit{
		hash:          hash,
		modifications: modifications,
		parent:        parent,
	}

	return &out
}

// Hash returns the hash
func (obj *commit) Hash() hash.Hash {
	return obj.hash
}

// Modifications returns the modifications
func (obj *commit) Modifications() modifications.Modifications {
	return obj.modifications
}

// HasParent returns true if there is a parent, false otherwise
func (obj *commit) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *commit) Parent() hash.Hash {
	return obj.parent
}
