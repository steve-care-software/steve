package roots

import (
	"crypto/ed25519"
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	amount      uint64
	owner       ed25519.PublicKey
	commit      hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		amount:      0,
		owner:       nil,
		commit:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithAmount adds an amount to the builder
func (app *builder) WithAmount(amount uint64) Builder {
	app.amount = amount
	return app
}

// WithOwner adds an owner to the builder
func (app *builder) WithOwner(owner ed25519.PublicKey) Builder {
	app.owner = owner
	return app
}

// WithCommit adds a commit to the builder
func (app *builder) WithCommit(commit hash.Hash) Builder {
	app.commit = commit
	return app
}

// Now builds a new Root instance
func (app *builder) Now() (Root, error) {
	if app.amount <= 0 {
		return nil, errors.New("the amount is mandatory in order to build a Root instance")
	}

	if app.owner == nil {
		return nil, errors.New("the owner is mandatory in order to build a Root instance")
	}

	if app.commit == nil {
		return nil, errors.New("the commit is mandatory in order to build a Root instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(fmt.Sprintf("%d", app.amount)),
		app.owner,
		app.commit.Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createRoot(*pHash, app.amount, app.owner, app.commit), nil
}
