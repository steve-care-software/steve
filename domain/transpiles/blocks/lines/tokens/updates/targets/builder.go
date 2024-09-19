package targets

import "errors"

type builder struct {
	constant string
	rule     string
}

func createBuilder() Builder {
	out := builder{
		constant: "",
		rule:     "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithConstant adds a constant to the builder
func (app *builder) WithConstant(constant string) Builder {
	app.constant = constant
	return app
}

// WithRule adds a rule to the builder
func (app *builder) WithRule(rule string) Builder {
	app.rule = rule
	return app
}

// Now builds a new Target instance
func (app *builder) Now() (Target, error) {
	if app.constant != "" {
		return createTargetWithConstant(app.constant), nil
	}

	if app.rule != "" {
		return createTargetWithRule(app.rule), nil
	}

	return nil, errors.New("the Target is invalid")
}
