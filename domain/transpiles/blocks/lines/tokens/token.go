package tokens

import (
	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens/pointers"
	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens/updates"
)

type token struct {
	update updates.Update
	del    pointers.Pointer
	insert pointers.Pointer
}

func createTokenWithUpdate(
	update updates.Update,
) Token {
	return createTokenInternally(update, nil, nil)
}

func createTokenWithDelete(
	del pointers.Pointer,
) Token {
	return createTokenInternally(nil, del, nil)
}

func createTokenWithInsert(
	insert pointers.Pointer,
) Token {
	return createTokenInternally(nil, nil, insert)
}

func createTokenInternally(
	update updates.Update,
	del pointers.Pointer,
	insert pointers.Pointer,
) Token {
	out := token{
		update: update,
		del:    del,
		insert: insert,
	}

	return &out
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
