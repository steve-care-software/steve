package heads

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads/access"
)

type builder struct {
	name    string
	version uint
	access  access.Access
}

func createBuilder() Builder {
	out := builder{
		name:    "",
		version: 0,
		access:  nil,
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

// WithVersion adds a version to the builder
func (app *builder) WithVersion(version uint) Builder {
	app.version = version
	return app
}

// WithAccess adds an access to the builder
func (app *builder) WithAccess(access access.Access) Builder {
	app.access = access
	return app
}

// Now builds a new Head instance
func (app *builder) Now() (Head, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Head instance")
	}

	if app.version <= 0 {
		return nil, errors.New("the version is mandatory in order to build an Head instance")
	}

	if app.access == nil {
		return nil, errors.New("the access is mandatory in order to build an Head instance")
	}

	return createHead(
		app.name,
		app.version,
		app.access,
	), nil
}
