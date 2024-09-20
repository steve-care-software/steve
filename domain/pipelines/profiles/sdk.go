package profiles

import "crypto"

// Profile represents a profile
type Profile interface {
	Version() uint
	Handle() string
	Name() string
	Description() string
	PublicKey() crypto.PublicKey
}
