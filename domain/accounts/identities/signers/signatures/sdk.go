package signatures

import "github.com/steve-care-software/steve/domain/accounts/identities/encryptors/publickeys"

// Signature represents a signature
type Signature interface {
	PublicKey(msg []byte) publickeys.PublicKey
	Verify() bool
	Bytes() []byte
}
