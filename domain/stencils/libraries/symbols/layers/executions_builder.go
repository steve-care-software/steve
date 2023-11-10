package layers

import (
	"errors"

	"github.com/steve-care-software/steve/domain/blockchains/hash"
)

type executionsBuilder struct {
	hashAdapter hash.Adapter
	list        []Execution
}

func createExecutionsBuilder(
	hashAdapter hash.Adapter,
) ExecutionsBuilder {
	out := executionsBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the executionsBuilder
func (app *executionsBuilder) Create() ExecutionsBuilder {
	return createExecutionsBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the executionsBuilder
func (app *executionsBuilder) WithList(list []Execution) ExecutionsBuilder {
	app.list = list
	return app
}

// Now builds a new Executions instance
func (app *executionsBuilder) Now() (Executions, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Execution in order to build a Executions instance")
	}

	data := [][]byte{}
	for _, oneExecution := range app.list {
		data = append(data, oneExecution.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createExecutions(*pHash, app.list), nil
}
