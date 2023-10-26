package parameters

import "errors"

type parameterBuilder struct {
	name  string
	pKind *uint8
}

func createParameterBuilder() ParameterBuilder {
	out := parameterBuilder{
		name:  "",
		pKind: nil,
	}

	return &out
}

// Create initializes the builder
func (app *parameterBuilder) Create() ParameterBuilder {
	return createParameterBuilder()
}

// WithName adds a name to the builder
func (app *parameterBuilder) WithName(name string) ParameterBuilder {
	app.name = name
	return app
}

// WithKind adds a kind to the builder
func (app *parameterBuilder) WithKind(kind uint8) ParameterBuilder {
	app.pKind = &kind
	return app
}

// Now builds a new Parameter instance
func (app *parameterBuilder) Now() (Parameter, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Parameter instance")
	}

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Parameter instance")
	}

	kind := *app.pKind
	if !Validate(kind) {
		return nil, errors.New("the kind was expected to be one of these: tree, token, layer, link")
	}

	return createParameter(app.name, kind), nil
}
