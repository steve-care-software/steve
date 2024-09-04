package roots

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	amount      uint64
	owner       hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		amount:      0,
		owner:       nil,
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
func (app *builder) WithOwner(owner hash.Hash) Builder {
	app.owner = owner
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

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(fmt.Sprintf("%d", app.amount)),
		app.owner.Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createRoot(*pHash, app.amount, app.owner), nil
}
