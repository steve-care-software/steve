package contents

import (
	"errors"

	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	trx         transactions.Transactions
	parent      hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		trx:         nil,
		parent:      nil,
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

// Now builds a new Content instance
func (app *builder) Now() (Content, error) {
	if app.trx == nil {
		return nil, errors.New("the transactions is mandatory in order to build a Content instance")
	}

	if app.parent == nil {
		return nil, errors.New("the parent hash is mandatory in order to build a Content instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.trx.Hash().Bytes(),
		app.parent.Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createContent(*pHash, app.trx, app.parent), nil
}
