package tokens

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens/pointers"
	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens/updates"
)

type token struct {
	hash   hash.Hash
	update updates.Update
	del    pointers.Pointer
	insert pointers.Pointer
}

func createTokenWithUpdate(
	hash hash.Hash,
	update updates.Update,
) Token {
	return createTokenInternally(hash, update, nil, nil)
}

func createTokenWithDelete(
	hash hash.Hash,
	del pointers.Pointer,
) Token {
	return createTokenInternally(hash, nil, del, nil)
}

func createTokenWithInsert(
	hash hash.Hash,
	insert pointers.Pointer,
) Token {
	return createTokenInternally(hash, nil, nil, insert)
}

func createTokenInternally(
	hash hash.Hash,
	update updates.Update,
	del pointers.Pointer,
	insert pointers.Pointer,
) Token {
	out := token{
		hash:   hash,
		update: update,
		del:    del,
		insert: insert,
	}

	return &out
}

// Hash returns the hash
func (obj *token) Hash() hash.Hash {
	return obj.hash
}

// IsUpdate returns true if there is an update, false otherwise
func (obj *token) IsUpdate() bool {
	return obj.update != nil
}

// Update returns the update, if any
func (obj *token) Update() updates.Update {
	return obj.update
}

// IsDelete returns true if there is a delete, false otherwise
func (obj *token) IsDelete() bool {
	return obj.del != nil
}

// Delete returns the delete, if any
func (obj *token) Delete() pointers.Pointer {
	return obj.del
}

// IsInsert returns true if there is an insert, false otherwise
func (obj *token) IsInsert() bool {
	return obj.insert != nil
}

// Insert returns the insert, if any
func (obj *token) Insert() pointers.Pointer {
	return obj.insert
}
