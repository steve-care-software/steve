package paths

import "github.com/steve-care-software/steve/domain/connections"

type pathBuilder struct {
	possibilities Paths
	destination   connections.Connection
}

func createPathBuilder() PathBuilder {
	out := pathBuilder{
		possibilities: nil,
		destination:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *pathBuilder) Create() PathBuilder {
	return createPathBuilder()
}

// WithPossibilities add possibilities to the builder
func (app *pathBuilder) WithPossibilities(paths Paths) PathBuilder {
	app.possibilities = paths
	return app
}

// WithDestination add destination to the builder
func (app *pathBuilder) WithDestination(destination connections.Connection) PathBuilder {
	app.destination = destination
	return app
}

// Now builds a new Path instance
func (app *pathBuilder) Now() (Path, error) {
	if app.possibilities == nil {

	}

	if app.destination != nil {

	}

	return createPath(app.possibilities), nil
}
