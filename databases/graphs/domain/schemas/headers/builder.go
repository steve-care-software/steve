package headers

import "errors"

type builder struct {
	version uint
	name    string
}

func createBuilder() Builder {
	out := builder{
		version: 0,
		name:    "",
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

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// Now builds a new Header instance
func (app *builder) Now() (Header, error) {
	if app.version <= 0 {
		return nil, errors.New("the version must be greater than zero (0) in order to build an Header instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Header instance")
	}

	return createHeader(
		app.version,
		app.name,
	), nil
}
