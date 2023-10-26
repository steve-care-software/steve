package layers

import (
	"errors"

	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/reduces"
)

type asignableBuilder struct {
	query   Query
	reduce  reduces.Reduce
	compare ConstantValues
	length  ConstantValue
	join    ConstantValues
	value   ConstantValue
}

func createAssignableBuilder() AssignableBuilder {
	out := asignableBuilder{
		query:   nil,
		reduce:  nil,
		compare: nil,
		length:  nil,
		join:    nil,
		value:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *asignableBuilder) Create() AssignableBuilder {
	return createAssignableBuilder()
}

// WithQuery adds a query to the builder
func (app *asignableBuilder) WithQuery(query Query) AssignableBuilder {
	app.query = query
	return app
}

// WithReduce adds a reduce to the builder
func (app *asignableBuilder) WithReduce(reduce reduces.Reduce) AssignableBuilder {
	app.reduce = reduce
	return app
}

// WithCompare adds a compare to the builder
func (app *asignableBuilder) WithCompare(compare ConstantValues) AssignableBuilder {
	app.compare = compare
	return app
}

// WithLength adds a length to the builder
func (app *asignableBuilder) WithLength(length ConstantValue) AssignableBuilder {
	app.length = length
	return app
}

// WithJoin adds a join to the builder
func (app *asignableBuilder) WithJoin(join ConstantValues) AssignableBuilder {
	app.join = join
	return app
}

// WithValue adds a join to the builder
func (app *asignableBuilder) WithValue(value ConstantValue) AssignableBuilder {
	app.value = value
	return app
}

// Now builds a new Assignable instance
func (app *asignableBuilder) Now() (Assignable, error) {
	if app.query != nil {
		return createAssignableWithQuery(app.query), nil
	}

	if app.reduce != nil {
		return createAssignableWithReduce(app.reduce), nil
	}

	if app.compare != nil {
		return createAssignableWithCompare(app.compare), nil
	}

	if app.length != nil {
		return createAssignableWithLength(app.length), nil
	}

	if app.join != nil {
		return createAssignableWithJoin(app.join), nil
	}

	if app.value != nil {
		return createAssignableWithValue(app.value), nil
	}

	return nil, errors.New("the Assignable is invalid")
}
