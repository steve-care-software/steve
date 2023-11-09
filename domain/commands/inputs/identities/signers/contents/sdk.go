package contents

import (
	"github.com/steve-care-software/steve/domain/commands/inputs/identities/signers/contents/signs"
	"github.com/steve-care-software/steve/domain/commands/inputs/identities/signers/contents/votes"
)

// Builder represents a content builder
type Builder interface {
	Create() Builder
	WithSign(sign signs.Sign) Builder
	WithVote(vote votes.Vote) Builder
	IsPublicKey() Builder
	IsBytes() Builder
	Now() (Content, error)
}

// Content represents a content
type Content interface {
	IsBytes() bool
	IsPublicKey() bool
	IsSign() bool
	Sign() signs.Sign
	IsVote() bool
	Vote() votes.Vote
}
