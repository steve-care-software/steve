package transfers

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

// Builder represents the transfer builder
type Builder interface {
	Create() Builder
	WithVersion(version uint) Builder
	WithAmount(amount uint64) Builder
	WithPublicKey(pubKey ed25519.PublicKey) Builder
	Now() (Transfer, error)
}

// Transfer represents a transfer
type Transfer interface {
	Hash() hash.Hash
	Version() uint
	Amount() uint64
	PublicKey() ed25519.PublicKey
}
