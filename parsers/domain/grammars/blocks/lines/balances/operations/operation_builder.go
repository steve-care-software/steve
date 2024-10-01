package operations

import "errors"

type operationBuilder struct {
	actor Actor
	tail  Tail
	isNot bool
}

func createOperationBuilder() OperationBuilder {
	out := operationBuilder{
		actor: nil,
		tail:  nil,
		isNot: false,
	}

	return &out
}

// Create initializes the builder
func (app *operationBuilder) Create() OperationBuilder {
	return createOperationBuilder()
}

// WithActor adds an actor to the builder
func (app *operationBuilder) WithActor(actor Actor) OperationBuilder {
	app.actor = actor
	return app
}

// WithTail adds a tail to the builder
func (app *operationBuilder) WithTail(tail Tail) OperationBuilder {
	app.tail = tail
	return app
}

// IsNot flags the builder as not
func (app *operationBuilder) IsNot() OperationBuilder {
	app.isNot = true
	return app
}

// Now builds a new Operation instance
func (app *operationBuilder) Now() (Operation, error) {
	if app.actor == nil {
		return nil, errors.New("the actor is mandatory in order to build an Operation instance")
	}

	if app.tail == nil {
		return nil, errors.New("the tail is mandatory in order to build an Operation instance")
	}

	return createOperation(
		app.actor,
		app.tail,
		app.isNot,
	), nil
}
