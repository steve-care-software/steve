package profiles

import (
	"crypto/ed25519"
	"errors"
	"strconv"

	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	version     uint
	handle      string
	name        string
	description string
	pubKey      ed25519.PublicKey
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		version:     0,
		handle:      "",
		name:        "",
		description: "",
		pubKey:      nil,
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

// WithHandle adds an handle to the builder
func (app *builder) WithHandle(handle string) Builder {
	app.handle = handle
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithDescription adds a description to the builder
func (app *builder) WithDescription(description string) Builder {
	app.description = description
	return app
}

// WithPublicKey adds a publicKey to the builder
func (app *builder) WithPublicKey(pubKey ed25519.PublicKey) Builder {
	app.pubKey = pubKey
	return app
}

// Now builds a new Profile instance
func (app *builder) Now() (Profile, error) {
	if app.version <= 0 {
		return nil, errors.New("the version must be greater than zero (0) in order to build a Profile instance")
	}

	if app.handle == "" {
		return nil, errors.New("the handle is mandatory in order to create a Profile instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to create a Profile instance")
	}

	if app.pubKey == nil {
		return nil, errors.New("the publicKey is mandatory in order to create a Profile instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strconv.Itoa(int(app.version))),
		[]byte(app.handle),
		[]byte(app.name),
		[]byte(app.description),
		app.pubKey,
	})

	if err != nil {
		return nil, err
	}

	return createProfile(
		*pHash,
		app.version,
		app.handle,
		app.name,
		app.description,
		app.pubKey,
	), nil
}
