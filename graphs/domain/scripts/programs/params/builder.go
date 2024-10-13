package params

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/kinds"
)

type builder struct {
	kind        kinds.Kind
	internal    string
	external    string
	isMandatory bool
}

func createBuilder() Builder {
	return &builder{
		kind:        nil,
		internal:    "",
		external:    "",
		isMandatory: false,
	}
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

func (app *builder) WithKind(kind kinds.Kind) Builder {
	app.kind = kind
	return app
}

func (app *builder) WithInternal(internal string) Builder {
	app.internal = internal
	return app
}

func (app *builder) WithExternal(external string) Builder {
	app.external = external
	return app
}

func (app *builder) IsMandatory() Builder {
	app.isMandatory = true
	return app
}

func (app *builder) Now() (Params, error) {
	if app.kind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Params instance")
	}

	if app.internal == "" {
		return nil, errors.New("the internal is mandatory in order to build a Params instance")
	}

	if app.external == "" {
		return nil, errors.New("the external is mandatory in order to build a Params instance")
	}

	return createParams(
		app.kind,
		app.internal,
		app.external,
		app.isMandatory,
	), nil
}
