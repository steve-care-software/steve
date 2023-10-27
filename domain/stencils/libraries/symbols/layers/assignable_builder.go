package layers

import (
	"errors"

	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/accounts"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/constantvalues"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/reduces"
)

type assignableBuilder struct {
	query   Query
	reduce  reduces.Reduce
	compare constantvalues.ConstantValues
	length  constantvalues.ConstantValue
	join    constantvalues.ConstantValues
	value   constantvalues.ConstantValue
	account accounts.Account
}

func createAssignableBuilder() AssignableBuilder {
	out := assignableBuilder{
		query:   nil,
		reduce:  nil,
		compare: nil,
		length:  nil,
		join:    nil,
		value:   nil,
		account: nil,
	}

	return &out
}

// Create initializes the builder
func (app *assignableBuilder) Create() AssignableBuilder {
	return createAssignableBuilder()
}

// WithQuery adds a query to the builder
func (app *assignableBuilder) WithQuery(query Query) AssignableBuilder {
	app.query = query
	return app
}

// WithReduce adds a reduce to the builder
func (app *assignableBuilder) WithReduce(reduce reduces.Reduce) AssignableBuilder {
	app.reduce = reduce
	return app
}

// WithCompare adds a compare to the builder
func (app *assignableBuilder) WithCompare(compare constantvalues.ConstantValues) AssignableBuilder {
	app.compare = compare
	return app
}

// WithLength adds a length to the builder
func (app *assignableBuilder) WithLength(length constantvalues.ConstantValue) AssignableBuilder {
	app.length = length
	return app
}

// WithJoin adds a join to the builder
func (app *assignableBuilder) WithJoin(join constantvalues.ConstantValues) AssignableBuilder {
	app.join = join
	return app
}

// WithValue adds a join to the builder
func (app *assignableBuilder) WithValue(value constantvalues.ConstantValue) AssignableBuilder {
	app.value = value
	return app
}

// WithAccount adds an account to the builder
func (app *assignableBuilder) WithAccount(account accounts.Account) AssignableBuilder {
	app.account = account
	return app
}

// Now builds a new Assignable instance
func (app *assignableBuilder) Now() (Assignable, error) {
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

	if app.account != nil {
		return createAssignableWithAccount(app.account), nil
	}

	return nil, errors.New("the Assignable is invalid")
}
