package conditions

import (
	"errors"
)

type clausesBuilder struct {
	list []Clause
}

func createClausesBuilder() ClausesBuilder {
	out := clausesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *clausesBuilder) Create() ClausesBuilder {
	return createClausesBuilder()
}

// WithList adds a list to the builder
func (app *clausesBuilder) WithList(list []Clause) ClausesBuilder {
	app.list = list
	return app
}

// Now builds a new Clauses instance
func (app *clausesBuilder) Now() (Clauses, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Clause in order to build a Clauses instance")
	}

	return createClauses(app.list), nil
}
