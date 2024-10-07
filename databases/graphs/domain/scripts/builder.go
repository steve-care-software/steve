package scripts

import (
	"errors"

	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas"
)

type builder struct {
	schema schemas.Schema
}

func createBuilder() Builder {
	out := builder{
		schema: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithSchema adds a schema to the builder
func (app *builder) WithSchema(schema schemas.Schema) Builder {
	app.schema = schema
	return app
}

// Now builds a new Script instance
func (app *builder) Now() (Script, error) {
	if app.schema != nil {
		return createScriptWithSchema(app.schema), nil
	}

	return nil, errors.New("the Script is invalid")
}
