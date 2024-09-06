package resources

import (
	"errors"

	"github.com/steve-care-software/steve/domain/stores/resources/pointers"
)

type resourceBuilder struct {
	identifier string
	pointers   pointers.Pointers
}

func createResourceBuilder() ResourceBuilder {
	out := resourceBuilder{
		identifier: "",
		pointers:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *resourceBuilder) Create() ResourceBuilder {
	return createResourceBuilder()
}

// WithIdentifier adds an identifier to the builder
func (app *resourceBuilder) WithIdentifier(identifier string) ResourceBuilder {
	app.identifier = identifier
	return app
}

// WithPointers add pointers to the builder
func (app *resourceBuilder) WithPointers(pointers pointers.Pointers) ResourceBuilder {
	app.pointers = pointers
	return app
}

// Now builds a new Resource instance
func (app *resourceBuilder) Now() (Resource, error) {
	if app.identifier == "" {
		return nil, errors.New("the identifier is mandatory in order to build a Resource instance")
	}

	if app.pointers == nil {
		return nil, errors.New("the pointers is mandatory in order to build a Resource instance")
	}

	return createResource(
		app.identifier,
		app.pointers,
	), nil
}
