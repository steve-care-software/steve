package layers

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type executionBuilder struct {
	hashAdapter hash.Adapter
	isStop      bool
	assignment  Assignment
	condition   Condition
}

func createExecutionBuilder(
	hashAdapter hash.Adapter,
) ExecutionBuilder {
	out := executionBuilder{
		hashAdapter: hashAdapter,
		isStop:      false,
		assignment:  nil,
		condition:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *executionBuilder) Create() ExecutionBuilder {
	return createExecutionBuilder(
		app.hashAdapter,
	)
}

// WithAssignment adds an assignment to the builder
func (app *executionBuilder) WithAssignment(assignmnet Assignment) ExecutionBuilder {
	app.assignment = assignmnet
	return app
}

// WithCondition adds a condition to the builder
func (app *executionBuilder) WithCondition(condition Condition) ExecutionBuilder {
	app.condition = condition
	return app
}

// IsStop flags the builder as a stop
func (app *executionBuilder) IsStop() ExecutionBuilder {
	app.isStop = true
	return app
}

// Now builds a new Execution instance
func (app *executionBuilder) Now() (Execution, error) {
	data := [][]byte{}
	if app.isStop {
		data = append(data, []byte{0})
	}

	if app.assignment != nil {
		data = append(data, app.assignment.Hash().Bytes())
	}

	if app.condition != nil {
		data = append(data, app.condition.Hash().Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the Execution is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.isStop {
		return createExecutionWithStop(*pHash), nil
	}

	if app.assignment != nil {
		return createExecutionWithAssignment(*pHash, app.assignment), nil
	}

	return createExecutionWithCondition(*pHash, app.condition), nil
}
