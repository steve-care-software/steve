package blockchains

import (
	"github.com/steve-care-software/steve/applications/blockchains/actions"
	"github.com/steve-care-software/steve/domain/accounts/identities/signers"
	"github.com/steve-care-software/steve/domain/blockchains"
	"github.com/steve-care-software/steve/domain/hash"
)

type application struct {
	actionBuilder        actions.Builder
	blockchainRepository blockchains.Repository
	hashAdapter          hash.Adapter
}

func createApplication(
	actionBuilder actions.Builder,
	blockchainRepository blockchains.Repository,
	hashAdapter hash.Adapter,
) Application {
	out := application{
		actionBuilder:        actionBuilder,
		blockchainRepository: blockchainRepository,
		hashAdapter:          hashAdapter,
	}

	return &out
}

// Action returns the action application
func (app *application) Action(signer signers.Signer) (actions.Application, error) {
	identifierBytes := signer.PublicKey().Bytes()
	pHash, err := app.hashAdapter.FromBytes(identifierBytes)
	if err != nil {
		return nil, err
	}

	blockchain, err := app.Retrieve(*pHash)
	if err != nil {
		return nil, err
	}

	return app.actionBuilder.Create().
		WithSigner(signer).
		WithBlockchain(blockchain).
		Now()
}

// Retrieve retrieves a blockchain
func (app *application) Retrieve(identifier hash.Hash) (blockchains.Blockchain, error) {
	return app.blockchainRepository.Retrieve(identifier)
}
