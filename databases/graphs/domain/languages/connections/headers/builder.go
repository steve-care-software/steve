package headers

import (
	"errors"

	"github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/headers/names"
)

type builder struct {
	name    names.Name
	reverse names.Name
}

func createBuilder() Builder {
	out := builder{
		name:    nil,
		reverse: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name names.Name) Builder {
	app.name = name
	return app
}

// WithReverse adds a reverse to the builder
func (app *builder) WithReverse(reverse names.Name) Builder {
	app.reverse = reverse
	return app
}

// Now builds a new Header instance
func (app *builder) Now() (Header, error) {
	if app.name == nil {
		return nil, errors.New("the name is mandatory in order to build an Header instance")
	}

	if app.reverse != nil {
		return createHeaderWithReverse(app.name, app.reverse), nil
	}

	return createHeader(app.name), nil
}
