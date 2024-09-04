package identities

import (
	"crypto/ed25519"

	"github.com/steve-care-software/steve/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPK(pk ed25519.PrivateKey) Builder
	WithFlags(flags []hash.Hash) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	Name() string
	PK() ed25519.PrivateKey
	HasFlags() bool
	Flags() []hash.Hash
}
