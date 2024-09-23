package transfers

import (
	"crypto/ed25519"
	"errors"
	"strconv"

	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	version     uint
	amount      uint64
	publicKey   ed25519.PublicKey
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		version:     0,
		amount:      0,
		publicKey:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
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
func (app *builder) WithPublicKey(pubKey ed25519.PublicKey) Builder {
	app.publicKey = pubKey
	return app
}

// Now builds a new Transfer instance
func (app *builder) Now() (Transfer, error) {
	if app.version <= 0 {
		return nil, errors.New("the version must be greater than zero (0) in order to build a Transfer instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount is mandatory in order to build a Transfer instance")
	}

	if app.publicKey == nil {
		return nil, errors.New("the publicKey is mandatory in order to build a Transfer instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strconv.Itoa(int(app.version))),
		[]byte(strconv.Itoa(int(app.amount))),
		[]byte(app.publicKey),
	})

	if err != nil {
		return nil, err
	}

	return createTransfer(
		*pHash,
		app.version,
		app.amount,
		app.publicKey,
	), nil
}
