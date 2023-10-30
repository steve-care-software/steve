package blockchains

import (
	"github.com/steve-care-software/steve/applications/blockchains/actions"
	"github.com/steve-care-software/steve/domain/accounts/identities/signers"
	"github.com/steve-care-software/steve/domain/blockchains"
	"github.com/steve-care-software/steve/domain/hash"
)

// Application represents a blockchain application
type Application interface {
	Action(signer signers.Signer) (actions.Application, error)
	Retrieve(identifier hash.Hash) (blockchains.Blockchain, error)
}
