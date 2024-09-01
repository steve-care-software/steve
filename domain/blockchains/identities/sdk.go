package identities

import (
	"crypto/ed25519"
)

// Identity represents an identity
type Identity interface {
	Name() string
	PK() ed25519.PrivateKey
}
