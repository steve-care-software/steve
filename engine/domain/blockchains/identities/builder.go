package identities

import (
	"crypto/ed25519"
	"errors"

	"github.com/steve-care-software/steve/engine/domain/hash"
)

type builder struct {
	name  string
	pk    ed25519.PrivateKey
	flags []hash.Hash
}

func createBuilder() Builder {
	out := builder{
		name:  "",
		pk:    nil,
		flags: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithPK adds a pk to the builder
func (app *builder) WithPK(pk ed25519.PrivateKey) Builder {
	app.pk = pk
	return app
}

// WithFlags add flags to the builder
func (app *builder) WithFlags(flags []hash.Hash) Builder {
	app.flags = flags
	return app
}

// Now builds a new Identity instance
func (app *builder) Now() (Identity, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Identity instance")
	}

	if app.pk == nil {
		return nil, errors.New("the pk is mandatory in order to build an Identity instance")
	}

	if app.flags != nil && len(app.flags) <= 0 {
		app.flags = nil
	}

	if app.flags != nil {
		return createIdentityWithFlags(app.name, app.pk, app.flags), nil
	}

	return createIdentity(app.name, app.pk), nil
}
