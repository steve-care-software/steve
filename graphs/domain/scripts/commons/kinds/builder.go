package kinds

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/kinds/primitives"
)

type builder struct {
	container Container
	pEngine   *uint8
	primitive primitives.Primitive
	isMap     bool
}

func createBuilder() Builder {
	out := builder{
		container: nil,
		pEngine:   nil,
		primitive: nil,
		isMap:     false,
	}

	return &out
}

// Create initializes the kind builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithContainer adds a container to the kind builder
func (app *builder) WithContainer(container Container) Builder {
	app.container = container
	return app
}

// WithEngine adds an engine to the kind builder
func (app *builder) WithEngine(engine uint8) Builder {
	app.pEngine = &engine
	return app
}

// WithPrimitive adds a primitive to the kind builder
func (app *builder) WithPrimitive(primitive primitives.Primitive) Builder {
	app.primitive = primitive
	return app
}

// IsMap marks the kind as a map
func (app *builder) IsMap() Builder {
	app.isMap = true
	return app
}

// Now builds a new Kind instance
func (app *builder) Now() (Kind, error) {
	if app.container != nil {
		return createKindWithContainer(app.container), nil
	}

	if app.pEngine != nil {
		flag := *app.pEngine
		if flag > EngineBridges {
			str := fmt.Sprintf("the engine flag (%d) is invalid while building the Kind instance", EngineBridges)
			return nil, errors.New(str)
		}

		return createKindWithEngine(app.pEngine), nil
	}

	if app.primitive != nil {
		return createKindWithPrimitive(app.primitive), nil
	}

	if app.isMap {
		return createKindWithMap(), nil
	}

	return nil, errors.New("the Kind is invalid")

}
