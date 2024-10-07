package permissions

import "errors"

type builder struct {
	names        []string
	compensation float64
}

func createBuilder() Builder {
	out := builder{
		names:        nil,
		compensation: 0.0,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithNames adds namess to the builder
func (app *builder) WithNames(names []string) Builder {
	app.names = names
	return app
}

// WithCompensation adds a compensation to the builder
func (app *builder) WithCompensation(compensation float64) Builder {
	app.compensation = compensation
	return app
}

// Now builds a new Permission instance
func (app *builder) Now() (Permission, error) {
	if app.names != nil && len(app.names) <= 0 {
		app.names = nil
	}

	if app.names == nil {
		return nil, errors.New("there must be at least 1 permission names in order to build a Permission instance")
	}

	if app.compensation > 0.0 {
		return createPermissionWithCompensation(
			app.names,
			app.compensation,
		), nil
	}

	return createPermission(
		app.names,
	), nil
}
