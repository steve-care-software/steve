package identities

import (
	"crypto/ed25519"

	"github.com/steve-care-software/steve/engine/domain/hash"
)

const dataLengthTooSmallErrPattern = "(identity) the data length was expected to be at least %d bytes, %d returned"

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	hashAdapter := hash.NewAdapter()
	builder := NewBuilder()
	return createAdapter(
		hashAdapter,
		builder,
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents an identity adapter
type Adapter interface {
	ToBytes(ins Identity) ([]byte, error)
	ToInstance(bytes []byte) (Identity, []byte, error)
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
