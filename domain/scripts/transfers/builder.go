package transfers

import (
	"crypto"
	"errors"
)

type builder struct {
	version   uint
	amount    uint64
	publicKey crypto.PublicKey
}

func createBuilder() Builder {
	out := builder{
		version:   0,
		amount:    0,
		publicKey: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithVersion adds a version to the builder
func (app *builder) WithVersion(version uint) Builder {
	app.version = version
	return app
}

// WithAmount adds an amount to the builder
func (app *builder) WithAmount(amount uint64) Builder {
	app.amount = amount
	return app
}

// WithPublicKey adds a pubKey to the builder
func (app *builder) WithPublicKey(pubKey crypto.PublicKey) Builder {
	app.publicKey = pubKey
	return app
}

// Now builds a new Transfer instance
func (app *builder) Now() (Transfer, error) {
	if app.version <= 0 {
		return nil, errors.New("the version is mandatory in order to build a Transfer instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount is mandatory in order to build a Transfer instance")
	}

	if app.publicKey == nil {
		return nil, errors.New("the publicKey is mandatory in order to build a Transfer instance")
	}

	return createTransfer(
		app.version,
		app.amount,
		app.publicKey,
	), nil
}
