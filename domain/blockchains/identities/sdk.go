package identities

import (
	"crypto/ed25519"

	"github.com/steve-care-software/steve/domain/hash"
)

// Identity represents an identity
type Identity interface {
	Name() string
	PK() ed25519.PrivateKey
	HasFlags() bool
	Flags() []hash.Hash
}
