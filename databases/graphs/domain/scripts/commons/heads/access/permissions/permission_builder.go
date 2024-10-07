package permissions

import "errors"

type permissionBuilder struct {
	name         string
	compensation float64
}

func createPermissionBuilder() PermissionBuilder {
	out := permissionBuilder{
		name:         "",
		compensation: 0.0,
	}

	return &out
}

// Create initializes the builder
func (app *permissionBuilder) Create() PermissionBuilder {
	return createPermissionBuilder()
}

// WithName adds a name to the builder
func (app *permissionBuilder) WithName(name string) PermissionBuilder {
	app.name = name
	return app
}

// WithCompensation adds a compensation to the builder
func (app *permissionBuilder) WithCompensation(compensation float64) PermissionBuilder {
	app.compensation = compensation
	return app
}

// Now builds a new Permission instance
func (app *permissionBuilder) Now() (Permission, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Permission instance")
	}

	if app.compensation <= 0.0 {
		return nil, errors.New("the compensation is mandatory in order to build a Permission instance")
	}

	return createPermission(
		app.name,
		app.compensation,
	), nil
}
