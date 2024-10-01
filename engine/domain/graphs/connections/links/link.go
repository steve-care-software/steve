package links

import (
	"fmt"

	"github.com/steve-care-software/steve/commons/hash"
)

type link struct {
	hash    hash.Hash
	name    string
	reverse string
}

func createLink(
	hash hash.Hash,
	name string,
) Link {
	return createLinkInternally(hash, name, "")
}

func createLinkWithReverse(
	hash hash.Hash,
	name string,
	reverse string,
) Link {
	return createLinkInternally(hash, name, reverse)
}

func createLinkInternally(
	hash hash.Hash,
	name string,
	reverse string,
) Link {
	out := link{
		hash:    hash,
		name:    name,
		reverse: reverse,
	}

	return &out
}

// Hash returns the hash, if any
func (obj *link) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *link) Name() string {
	return obj.name
}

// HasReverse returns true if there is a reverse, false otherwise
func (obj *link) HasReverse() bool {
	return obj.reverse != ""
}

// Reverse returns the reverse, if any
func (obj *link) Reverse() string {
	return obj.reverse
}

// Debug returns the string debug representation of the link
func (obj *link) Debug() string {
	reverseStr := ""
	if obj.HasReverse() {
		reverseStr = fmt.Sprintf("reverse: %s, ", obj.Reverse())
	}
	return fmt.Sprintf("name: %s", obj.name, reverseStr)
}
