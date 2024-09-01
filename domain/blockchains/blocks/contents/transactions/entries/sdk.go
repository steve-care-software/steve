package entries

import "github.com/steve-care-software/steve/domain/hash"

// Entry represents an entry
type Entry interface {
	Hash() hash.Hash
	Script() hash.Hash
	Size() uint64
	Fees() uint64
}
