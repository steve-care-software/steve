package signers

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/signers/publickeys"
	"github.com/steve-care-software/steve/domain/accounts/identities/signers/signatures"
	"github.com/steve-care-software/steve/domain/accounts/identities/signers/votes"
)

// SignerFactory represents a signer factory
type SignerFactory interface {
	Create() Signer
}

// Signer represents a signer
type Signer interface {
	PublicKey() publickeys.PublicKey
	Sign(msg []byte) (signatures.Signature, error)
	Vote(msg []byte, ring []publickeys.PublicKey) (votes.Vote, error)
	Bytes() []byte
}
