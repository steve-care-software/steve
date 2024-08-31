package queries

import "errors"

type builder struct {
	list []Query
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Query) Builder {
	app.list = list
	return app
}

// Now builds a new Queries instance
func (app *builder) Now() (Queries, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Query in order to build a Queries instance")
	}

	return createQueries(app.list), nil
}
