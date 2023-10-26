package kinds

import "errors"

type builder struct {
	isContinue bool
	isPrompt   bool
	execute    []string
}

func createBuilder() Builder {
	out := builder{
		isContinue: false,
		isPrompt:   false,
		execute:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithExecute adds execute commands to the builder
func (app *builder) WithExecute(execute []string) Builder {
	app.execute = execute
	return app
}

// IsContinue flags the builder as a continue
func (app *builder) IsContinue() Builder {
	app.isContinue = true
	return app
}

// IsPrompt flags the builder as a prompt
func (app *builder) IsPrompt() Builder {
	app.isPrompt = true
	return app
}

// Now builds a new Kind instance
func (app *builder) Now() (Kind, error) {
	if app.isContinue {
		return createKindWithContinue(), nil
	}

	if app.isPrompt {
		return createKindWithPrompt(), nil
	}

	if app.execute != nil && len(app.execute) <= 0 {
		app.execute = nil
	}

	if app.execute != nil {
		return createKindWithExecute(app.execute), nil
	}

	return nil, errors.New("the Kind is invalid")
}
