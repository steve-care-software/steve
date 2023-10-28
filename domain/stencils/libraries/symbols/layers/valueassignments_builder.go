package layers

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type valueAssignmentsBuilder struct {
	hashAdapter hash.Adapter
	list        []ValueAssignment
}

func createValueAssignmentsBuilder(
	hashAdapter hash.Adapter,
) ValueAssignmentsBuilder {
	out := valueAssignmentsBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the valueAssignmentsBuilder
func (app *valueAssignmentsBuilder) Create() ValueAssignmentsBuilder {
	return createValueAssignmentsBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the valueAssignmentsBuilder
func (app *valueAssignmentsBuilder) WithList(list []ValueAssignment) ValueAssignmentsBuilder {
	app.list = list
	return app
}

// Now builds a new ValueAssignments instance
func (app *valueAssignmentsBuilder) Now() (ValueAssignments, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 ValueAssignment in order to build a ValueAssignments instance")
	}

	data := [][]byte{}
	for _, oneAssignment := range app.list {
		data = append(data, oneAssignment.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createValueAssignments(*pHash, app.list), nil
}
