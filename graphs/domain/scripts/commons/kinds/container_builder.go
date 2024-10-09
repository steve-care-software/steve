package kinds

import "errors"

type containerBuilder struct {
	pFlag *uint8
	kind  Kind
}

func createContainerBuilder() ContainerBuilder {
	out := containerBuilder{
		pFlag: nil,
		kind:  nil,
	}

	return &out
}

// Create initializes the container builder
func (app *containerBuilder) Create() ContainerBuilder {
	return createContainerBuilder()
}

// WithFlag adds a flag to the container builder
func (app *containerBuilder) WithFlag(flag uint8) ContainerBuilder {
	app.pFlag = &flag
	return app
}

// WithKind adds a kind to the container builder
func (app *containerBuilder) WithKind(kind Kind) ContainerBuilder {
	app.kind = kind
	return app
}

// Now builds a new Container instance
func (app *containerBuilder) Now() (Container, error) {
	if app.pFlag == nil {
		return nil, errors.New("the flag is mandatory in order to build a Container instance")
	}

	if app.kind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Container instance")
	}

	flag := *app.pFlag
	if flag > ContainerSortedSet {
		return nil, errors.New("the flag is invalid while building a Container instance")
	}

	return createContainer(flag, app.kind), nil
}
