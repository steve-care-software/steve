package primitives

import "errors"

type builder struct {
	pFlag *uint8
}

func createBuilder() Builder {
	out := builder{
		pFlag: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithFlag adds a flag to the builder
func (app *builder) WithFlag(flag uint8) Builder {
	app.pFlag = &flag
	return app
}

// Now builds a new Primitive instance
func (app *builder) Now() (Primitive, error) {
	if app.pFlag == nil {
		return nil, errors.New("the flag is mandatory in order to build a Primitive instance")
	}

	flag := *app.pFlag
	if flag > FlagString {
		return nil, errors.New("the flag is invalid when building a Primitive instance")
	}

	return createPrimitive(
		flag,
	), nil
}
