package references

import (
	"time"

	"github.com/steve-care-software/steve/domain/hash"
)

type commit struct {
	hash      hash.Hash
	action    Action
	createdOn time.Time
	pParent   *hash.Hash
}

func createCommit(
	hash hash.Hash,
	action Action,
	createdOn time.Time,
) Commit {
	return createCommitInternally(hash, action, createdOn, nil)
}

func createCommitWithParent(
	hash hash.Hash,
	action Action,
	createdOn time.Time,
	pParent *hash.Hash,
) Commit {
	return createCommitInternally(hash, action, createdOn, pParent)
}

func createCommitInternally(
	hash hash.Hash,
	action Action,
	createdOn time.Time,
	pParent *hash.Hash,
) Commit {
	out := commit{
		hash:      hash,
		action:    action,
		createdOn: createdOn,
		pParent:   pParent,
	}

	return &out
}

// Hash returns the hash
func (obj *commit) Hash() hash.Hash {
	return obj.hash
}

// Action returns the action
func (obj *commit) Action() Action {
	return obj.action
}

// CreatedOn returns the creation time
func (obj *commit) CreatedOn() time.Time {
	return obj.createdOn
}

// HasParent returns true if there is a parent, false otherwise
func (obj *commit) HasParent() bool {
	return obj.pParent != nil
}

// Parent returns the parent, if any
func (obj *commit) Parent() *hash.Hash {
	return obj.pParent
}
