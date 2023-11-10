package votes

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/signers/publickeys"
	"github.com/steve-care-software/steve/domain/blockchains/hash"
)

// Vote represents a vote
type Vote interface {
	Hash() hash.Hash
	Ring() []publickeys.PublicKey
	Verify(msg []byte) bool
	String() []byte
}
