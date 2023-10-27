package shares

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/shares/holders"
	"github.com/steve-care-software/steve/domain/accounts/identities/shares/resolutions"
	"github.com/steve-care-software/steve/domain/accounts/identities/signers/votes"
)

// Shares represents shares
type Shares interface {
	List() []Share
}

// Share represents a share
type Share interface {
	Vote() votes.Vote
	Content() Content
}

// Content represents a share's content
type Content interface {
	Holder() holders.Holder
	Resolution() resolutions.Resolution
}
