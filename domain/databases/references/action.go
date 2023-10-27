package references

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/trees"
)

type action struct {
	hash   hash.Hash
	insert trees.HashTree
	del    trees.HashTree
}

func createActionWithInsertAndDelete(
	hash hash.Hash,
	insert trees.HashTree,
	del trees.HashTree,
) Action {
	return createActionInternally(hash, insert, del)
}

func createActionWithInsert(
	hash hash.Hash,
	insert trees.HashTree,
) Action {
	return createActionInternally(hash, insert, nil)
}

func createActionWithDelete(
	hash hash.Hash,
	del trees.HashTree,
) Action {
	return createActionInternally(hash, nil, del)
}

func createActionInternally(
	hash hash.Hash,
	insert trees.HashTree,
	del trees.HashTree,
) Action {
	out := action{
		hash:   hash,
		insert: insert,
		del:    del,
	}

	return &out
}

// Hash returns the hash
func (obj *action) Hash() hash.Hash {
	return obj.hash
}

// HasInsert retruns true if insert, false otherwise
func (obj *action) HasInsert() bool {
	return obj.insert != nil
}

// Insert returns the insert, if any
func (obj *action) Insert() trees.HashTree {
	return obj.insert
}

// HasDelete retruns true if delete, false otherwise
func (obj *action) HasDelete() bool {
	return obj.del != nil
}

// Delete returns the delete, if any
func (obj *action) Delete() trees.HashTree {
	return obj.del
}
