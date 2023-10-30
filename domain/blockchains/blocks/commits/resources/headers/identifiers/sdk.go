package identifiers

import "github.com/steve-care-software/steve/domain/hash"

// Builder represents an identifier's builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithKind(kind uint) Builder
	Now() (Identifier, error)
}

// Identifier represents an identifier
type Identifier interface {
	Hash() hash.Hash
	Kind() uint
}
