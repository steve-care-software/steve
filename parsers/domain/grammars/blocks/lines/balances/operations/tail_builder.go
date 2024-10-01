package operations

import (
	"errors"
	"fmt"
)

type tailBuilder struct {
	pOperator *uint8
	actor     Actor
}

func createTailBuilder() TailBuilder {
	out := tailBuilder{
		pOperator: nil,
		actor:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *tailBuilder) Create() TailBuilder {
	return createTailBuilder()
}

// WithOperator adds an operator to the builder
func (app *tailBuilder) WithOperator(operator uint8) TailBuilder {
	app.pOperator = &operator
	return app
}

// WithActor adds an actor to the builder
func (app *tailBuilder) WithActor(actor Actor) TailBuilder {
	app.actor = actor
	return app
}

// Now builds a new Tail instance
func (app *tailBuilder) Now() (Tail, error) {
	if app.pOperator == nil {
		return nil, errors.New("the operator is mandatory in order to build a Tail instance")
	}

	if app.actor == nil {
		return nil, errors.New("the actor is mandatory in order to build a Tail instance")
	}

	operator := *app.pOperator
	if operator > OperatorXor {
		str := fmt.Sprintf("the operator (flag: %d) is invalid when building a Tail instance", operator)
		return nil, errors.New(str)
	}

	return createTail(
		operator,
		app.actor,
	), nil
}
