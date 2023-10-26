package origins

import "errors"

type directionBuilder struct {
	next     Origin
	previous Origin
}

func createDirectionBuilder() DirectionBuilder {
	out := directionBuilder{
		next:     nil,
		previous: nil,
	}

	return &out
}

// Create initializes the builder
func (app *directionBuilder) Create() DirectionBuilder {
	return createDirectionBuilder()
}

// WithNext adds a next to the builder
func (app *directionBuilder) WithNext(next Origin) DirectionBuilder {
	app.next = next
	return app
}

// WithPrevious adds a previous to the builder
func (app *directionBuilder) WithPrevious(previous Origin) DirectionBuilder {
	app.previous = previous
	return app
}

// Now builds a new Direction instance
func (app *directionBuilder) Now() (Direction, error) {
	if app.next != nil {
		return createDirectionWithNext(app.next), nil
	}

	if app.previous != nil {
		return createDirectionWithPrevious(app.previous), nil
	}

	return nil, errors.New("the Direction is invalid")
}
