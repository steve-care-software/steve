package profiles

import (
	"crypto/ed25519"

	"github.com/steve-care-software/steve/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the profile builder
type Builder interface {
	Create() Builder
	WithVersion(version uint) Builder
	WithHandle(handle string) Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	WithPublicKey(pubKey ed25519.PublicKey) Builder
	Now() (Profile, error)
}

// Profile represents a profile
type Profile interface {
	Hash() hash.Hash
	Version() uint
	Handle() string
	Name() string
	Description() string
	PublicKey() ed25519.PublicKey
}
