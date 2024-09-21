package roles

import "github.com/steve-care-software/steve/domain/hash"

type role struct {
	hash   hash.Hash
	read   []string
	write  []string
	review []string
}

func createRoleWithRead(
	hash hash.Hash,
	read []string,
) Role {
	return createRoleInternally(hash, read, nil, nil)
}

func createRoleWithWrite(
	hash hash.Hash,
	write []string,
) Role {
	return createRoleInternally(hash, nil, write, nil)
}

func createRoleWithReview(
	hash hash.Hash,
	review []string,
) Role {
	return createRoleInternally(hash, nil, nil, review)
}

func createRoleWithReadAndWrite(
	hash hash.Hash,
	read []string,
	write []string,
) Role {
	return createRoleInternally(hash, read, write, nil)
}

func createRoleWithReadAndReview(
	hash hash.Hash,
	read []string,
	review []string,
) Role {
	return createRoleInternally(hash, read, nil, review)
}

func createRoleWithWriteAndReview(
	hash hash.Hash,
	write []string,
	review []string,
) Role {
	return createRoleInternally(hash, nil, write, review)
}

func createRoleWithReadAndWriteAndReview(
	hash hash.Hash,
	read []string,
	write []string,
	review []string,
) Role {
	return createRoleInternally(hash, read, write, review)
}

func createRoleInternally(
	hash hash.Hash,
	read []string,
	write []string,
	review []string,
) Role {
	out := role{
		hash:   hash,
		read:   read,
		write:  write,
		review: review,
	}

	return &out
}

// Hash returns the hash
func (obj *role) Hash() hash.Hash {
	return obj.hash
}

// HasRead returns true if there is read, false otherwise
func (obj *role) HasRead() bool {
	return obj.read != nil
}

// Read returns the read, if any
func (obj *role) Read() []string {
	return obj.read
}

// HasWrite returns true if there is write, false otherwise
func (obj *role) HasWrite() bool {
	return obj.write != nil
}

// Write returns the write, if any
func (obj *role) Write() []string {
	return obj.write
}

// HasReview returns true if there is review, false otherwise
func (obj *role) HasReview() bool {
	return obj.review != nil
}

// Review returns the review, if any
func (obj *role) Review() []string {
	return obj.review
}
