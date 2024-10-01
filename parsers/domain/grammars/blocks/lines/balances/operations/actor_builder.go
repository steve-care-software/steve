package operations

import (
	"errors"

	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/operations/selectors"
)

type actorBuilder struct {
	selector  selectors.Selector
	operation Operation
}

func createActorBuilder() ActorBuilder {
	out := actorBuilder{
		selector:  nil,
		operation: nil,
	}

	return &out
}

// Create initializes the builder
func (app *actorBuilder) Create() ActorBuilder {
	return createActorBuilder()
}

// WithSelector adds a selector to the builder
func (app *actorBuilder) WithSelector(selector selectors.Selector) ActorBuilder {
	app.selector = selector
	return app
}

// WithOperation adds an operation to the builder
func (app *actorBuilder) WithOperation(operation Operation) ActorBuilder {
	app.operation = operation
	return app
}

// Nwo builds a new Actor instance
func (app *actorBuilder) Now() (Actor, error) {
	if app.selector != nil {
		return createActorWithSelector(
			app.selector,
		), nil
	}

	if app.operation != nil {
		return createActorWithOperation(
			app.operation,
		), nil
	}

	return nil, errors.New("the Actor is invalid")
}
