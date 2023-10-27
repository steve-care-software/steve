package references

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/trees"
)

type actionBuilder struct {
	hashAdapter hash.Adapter
	insert      trees.HashTree
	del         trees.HashTree
}

func createActionBuilder(
	hashAdapter hash.Adapter,
) ActionBuilder {
	out := actionBuilder{
		hashAdapter: hashAdapter,
		insert:      nil,
		del:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *actionBuilder) Create() ActionBuilder {
	return createActionBuilder(
		app.hashAdapter,
	)
}

// WithInsert adds an insert to the builder
func (app *actionBuilder) WithInsert(insert trees.HashTree) ActionBuilder {
	app.insert = insert
	return app
}

// WithDelete adds a delete to the builder
func (app *actionBuilder) WithDelete(del trees.HashTree) ActionBuilder {
	app.del = del
	return app
}

// Now builds a new Action instance
func (app *actionBuilder) Now() (Action, error) {

	data := [][]byte{}
	if app.insert != nil {
		data = append(data, app.insert.Head().Bytes())
	}

	if app.del != nil {
		data = append(data, app.del.Head().Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the action is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.insert != nil && app.del != nil {
		return createActionWithInsertAndDelete(*pHash, app.insert, app.del), nil
	}

	if app.insert != nil {
		return createActionWithInsert(*pHash, app.insert), nil
	}

	return createActionWithDelete(*pHash, app.del), nil
}
