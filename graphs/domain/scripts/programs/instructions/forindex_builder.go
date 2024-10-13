package instructions

import "errors"

type forIndexBuilder struct {
	clause       ForUntilClause
	instructions ForInstructions
}

func createForIndexBuilder() ForIndexBuilder {
	out := forIndexBuilder{
		clause:       nil,
		instructions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *forIndexBuilder) Create() ForIndexBuilder {
	return createForIndexBuilder()
}

// WithClause adds the clause to the builder
func (app *forIndexBuilder) WithClause(clause ForUntilClause) ForIndexBuilder {
	app.clause = clause
	return app
}

// WithInstructions adds instructions to the builder
func (app *forIndexBuilder) WithInstructions(instructions ForInstructions) ForIndexBuilder {
	app.instructions = instructions
	return app
}

// Now builds and returns the ForIndex instance
func (app *forIndexBuilder) Now() (ForIndex, error) {
	if app.clause == nil {
		return nil, errors.New("the clause is mandatory in order to build a ForIndex instance")
	}
	if app.instructions == nil {
		return nil, errors.New("the instructions are mandatory in order to build a ForIndex instance")
	}

	return createForIndex(
		app.clause,
		app.instructions,
	), nil
}
