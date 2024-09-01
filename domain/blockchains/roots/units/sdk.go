package units

import "github.com/steve-care-software/steve/domain/hash"

// Unit represents the root units
type Unit interface {
	Amount() uint64
	Owner() hash.Hash
}
