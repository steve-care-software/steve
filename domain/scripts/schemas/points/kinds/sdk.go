package kinds

import "github.com/steve-care-software/steve/domain/hash"

// Kind represents the point kind
type Kind interface {
	Hash() hash.Hash
	IsBytes() bool
	IsInt() bool
	IsUint() bool
	IsFloat() bool
	IsVector() bool
	Vector() Kind
}
