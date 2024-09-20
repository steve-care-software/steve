package transfers

import "crypto"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the transfer builder
type Builder interface {
	Create() Builder
	WithVersion(version uint) Builder
	WithAmount(amount uint64) Builder
	WithPublicKey(pubKey crypto.PublicKey) Builder
	Now() (Transfer, error)
}

// Transfer represents a transfer
type Transfer interface {
	Version() uint
	Amount() uint64
	PublicKey() crypto.PublicKey
}
