package contexts

type contextBuilder struct {
	name   string
	parent Context
}

func createContextBuilder() ContextBuilder {
	out := contextBuilder{
		name:   "",
		parent: nil,
	}

	return &out
}

// Create initializes the builder
func (app *contextBuilder) Create() ContextBuilder {
	return createContextBuilder()
}

// WithName adds a name to the builder
func (app *contextBuilder) WithName(name string) ContextBuilder {
	app.name = name
	return app
}

// WithParent adds a parent to the builder
func (app *contextBuilder) WithParent(parent Context) ContextBuilder {
	app.parent = parent
	return app
}

// Now builds a new Context instance
func (app *contextBuilder) Now() (Context, error) {
	if app.name == "" {

	}

	if app.parent != nil {
		return createContextWithParent(app.name, app.parent), nil
	}

	return createContext(app.name), nil
}
