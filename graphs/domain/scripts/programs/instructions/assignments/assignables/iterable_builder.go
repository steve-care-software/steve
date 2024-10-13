package assignables

import "errors"

type iterableBuilder struct {
	listMap  ListMap
	variable string
}

func createIterableBuilder() IterableBuilder {
	return &iterableBuilder{
		listMap:  nil,
		variable: "",
	}
}

// Create initializes the Iterable builder
func (app *iterableBuilder) Create() IterableBuilder {
	return createIterableBuilder()
}

// WithListMap adds a ListMap to the builder
func (app *iterableBuilder) WithListMap(listMap ListMap) IterableBuilder {
	app.listMap = listMap
	return app
}

// WithVariable adds a variable to the builder
func (app *iterableBuilder) WithVariable(variable string) IterableBuilder {
	app.variable = variable
	return app
}

// Now builds a new Iterable instance
func (app *iterableBuilder) Now() (Iterable, error) {
	if app.listMap != nil {
		return createIterableWithListMap(app.listMap), nil
	}

	if app.variable != "" {
		return createIterableWithVariable(app.variable), nil
	}

	return nil, errors.New("the Iterable is invalid")
}
