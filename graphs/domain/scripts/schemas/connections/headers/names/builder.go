package names

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/headers/names/cardinalities"
)

type builder struct {
	name        string
	cardinality cardinalities.Cardinality
}

func createBuilder() Builder {
	out := builder{
		name:        "",
		cardinality: nil,
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

// WithCardinality adds a cardinality to the builder
func (app *builder) WithCardinality(cardinality cardinalities.Cardinality) Builder {
	app.cardinality = cardinality
	return app
}

// Now builds a new Name instance
func (app *builder) Now() (Name, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Name instance")
	}

	if app.cardinality == nil {
		return nil, errors.New("the cardinality is mandatory in order to build a Name instance")
	}

	return createName(
		app.name,
		app.cardinality,
	), nil
}
