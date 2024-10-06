package walkers

import "errors"

type builder struct {
	fn   ElementFn
	list TokenList
}

func createBuilder() Builder {
	out := builder{
		fn:   nil,
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithFn adds an ElementFn to the builder
func (app *builder) WithFn(fn ElementFn) Builder {
	app.fn = fn
	return app
}

// WithList adds a list to the builder
func (app *builder) WithList(list TokenList) Builder {
	app.list = list
	return app
}

// Now builds a new Walker instance
func (app *builder) Now() (Walker, error) {
	if app.fn == nil {
		return nil, errors.New("the ElementFn is mandatory in order to build a Walker instance")
	}

	if app.list != nil {
		return createWalkerWithList(
			app.fn,
			app.list,
		), nil
	}

	return createWalker(
		app.fn,
	), nil
}
