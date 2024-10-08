package contents

import (
	"crypto/ed25519"
	"errors"

	"github.com/steve-care-software/steve/blockchains/domain/blocks/contents/transactions"
	"github.com/steve-care-software/steve/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	trx         transactions.Transactions
	parent      hash.Hash
	miner       ed25519.PublicKey
	commit      hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		trx:         nil,
		parent:      nil,
		miner:       nil,
		commit:      nil,
	}

	return &out
}

// Create intiializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithTransactions add transactions to the builder
func (app *builder) WithTransactions(trx transactions.Transactions) Builder {
	app.trx = trx
	return app
}

// WithParent add parent to the builder
func (app *builder) WithParent(parent hash.Hash) Builder {
	app.parent = parent
	return app
}

// WithMiner add miner to the builder
func (app *builder) WithMiner(miner ed25519.PublicKey) Builder {
	app.miner = miner
	return app
}

// WithCommit add commit to the builder
func (app *builder) WithCommit(commit hash.Hash) Builder {
	app.commit = commit
	return app
}

// Now builds a new Content instance
func (app *builder) Now() (Content, error) {
	if app.trx == nil {
		return nil, errors.New("the transactions is mandatory in order to build a Content instance")
	}

	if app.parent == nil {
		return nil, errors.New("the parent hash is mandatory in order to build a Content instance")
	}

	if app.miner == nil {
		return nil, errors.New("the miner is mandatory in order to build a Content instance")
	}

	if app.commit == nil {
		return nil, errors.New("the commit is mandatory in order to build a Content instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.trx.Hash().Bytes(),
		app.parent.Bytes(),
		app.miner,
		app.commit.Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createContent(*pHash, app.trx, app.parent, app.miner, app.commit), nil
}
