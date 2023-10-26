package votes

import "github.com/steve-care-software/steve/domain/accounts/identities/signers/publickeys"

// Vote represents a vote
type Vote interface {
	Ring() []publickeys.PublicKey
	Verify(msg []byte) bool
	String() []byte
}
