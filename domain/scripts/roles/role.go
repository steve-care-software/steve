package roles

import "github.com/steve-care-software/steve/domain/hash"

type role struct {
	hash    hash.Hash
	version uint
	name    string
	insert  []hash.Hash
	del     []hash.Hash
}

func createRoleWithInsert(
	hash hash.Hash,
	version uint,
	name string,
	insert []hash.Hash,
) Role {
	return createRoleInternally(hash, version, name, insert, nil)
}

func createRoleWithDelete(
	hash hash.Hash,
	version uint,
	name string,
	del []hash.Hash,
) Role {
	return createRoleInternally(hash, version, name, nil, del)
}

func createRoleWithInsertAndDelete(
	hash hash.Hash,
	version uint,
	name string,
	insert []hash.Hash,
	del []hash.Hash,
) Role {
	return createRoleInternally(hash, version, name, insert, del)
}

func createRoleInternally(
	hash hash.Hash,
	version uint,
	name string,
	insert []hash.Hash,
	del []hash.Hash,
) Role {
	out := role{
		hash:    hash,
		version: version,
		name:    name,
		insert:  insert,
		del:     del,
	}

	return &out
}

// Hash returns the hash
func (obj *role) Hash() hash.Hash {
	return obj.hash
}

// Version returns the hash
func (obj *role) Version() uint {
	return obj.version
}

// Name returns the name
func (obj *role) Name() string {
	return obj.name
}

// HasInsert returns true if there is an insert, false otherwise
func (obj *role) HasInsert() bool {
	return obj.insert != nil
}

// Insert returns the insert, if any
func (obj *role) Insert() []hash.Hash {
	return obj.insert
}

// HasDelete returns true if there is a delete, false otherwise
func (obj *role) HasDelete() bool {
	return obj.del != nil
}

// Delete returns the delete, if any
func (obj *role) Delete() []hash.Hash {
	return obj.del
}
