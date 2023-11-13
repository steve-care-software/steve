package successes

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/signers/publickeys"
	"github.com/steve-care-software/steve/domain/accounts/identities/signers/signatures"
	"github.com/steve-care-software/steve/domain/accounts/identities/signers/votes"
)

// Success represents success
type Success interface {
	IsPublicKey() bool
	PublicKey() publickeys.PublicKey
	IsSign() bool
	Sign() signatures.Signature
	IsVote() bool
	Vote() votes.Vote
	IsBytes() bool
	Bytes() []byte
}
