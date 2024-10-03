package externals

import "errors"

type builder struct {
	schema string
	point  string
}

func createBuilder() Builder {
	out := builder{
		schema: "",
		point:  "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithSchema adds a schema to the builder
func (app *builder) WithSchema(schema string) Builder {
	app.schema = schema
	return app
}

// WithPoint adds a point to the builder
func (app *builder) WithPoint(point string) Builder {
	app.point = point
	return app
}

// Now builds a new External instance
func (app *builder) Now() (External, error) {
	if app.schema == "" {
		return nil, errors.New("the schema is mandatory in order to build an External instance")
	}

	if app.point == "" {
		return nil, errors.New("the point is mandatory in order to build an External instance")
	}

	return createExternal(
		app.schema,
		app.point,
	), nil
}
