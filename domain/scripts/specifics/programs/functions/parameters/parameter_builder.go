package parameters

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/containers"
)

type parameterBuilder struct {
	hashAdapter hash.Adapter
	name        string
	container   containers.Container
	isMandatory bool
}

func createParameterBuilder(
	hashAdapter hash.Adapter,
) ParameterBuilder {
	out := parameterBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		container:   nil,
		isMandatory: false,
	}

	return &out
}

// Create initializes the builder
func (app *parameterBuilder) Create() ParameterBuilder {
	return createParameterBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *parameterBuilder) WithName(name string) ParameterBuilder {
	app.name = name
	return app
}

// WithContainer adds a container to the builder
func (app *parameterBuilder) WithContainer(container containers.Container) ParameterBuilder {
	app.container = container
	return app
}

// IsMandatory flags the builder as mandatory
func (app *parameterBuilder) IsMandatory() ParameterBuilder {
	app.isMandatory = true
	return app
}

// Now builds a new Parameter instance
func (app *parameterBuilder) Now() (Parameter, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Parameter instance")
	}

	if app.container == nil {
		return nil, errors.New("the container is mandatory in order to build a Parameter instance")
	}

	isMandatory := "false"
	if app.isMandatory {
		isMandatory = "true"
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		app.container.Hash().Bytes(),
		[]byte(isMandatory),
	})

	if err != nil {
		return nil, err
	}

	return createParameter(
		*pHash,
		app.name,
		app.container,
		app.isMandatory,
	), nil
}
