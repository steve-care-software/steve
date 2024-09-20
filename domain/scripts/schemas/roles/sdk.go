package roles

import "github.com/steve-care-software/steve/domain/hash"

// Role represents role
type Role interface {
	Hash() hash.Hash
	HasRead() bool
	Read() []string
	HasWrite() bool
	Write() []string
	HasReview() bool
	Review() []string
}
