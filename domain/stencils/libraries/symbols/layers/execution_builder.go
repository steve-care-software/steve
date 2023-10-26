package layers

import (
	"errors"
)

type executionBuilder struct {
	isStop     bool
	assignment Assignment
	condition  Condition
}

func createExecutionBuilder() ExecutionBuilder {
	out := executionBuilder{
		isStop:     false,
		assignment: nil,
		condition:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *executionBuilder) Create() ExecutionBuilder {
	return createExecutionBuilder()
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
	if app.isStop {
		return createExecutionWithStop(), nil
	}

	if app.assignment != nil {
		return createExecutionWithAssignment(app.assignment), nil
	}

	if app.condition != nil {
		return createExecutionWithCondition(app.condition), nil
	}

	return nil, errors.New("the Execution is invalid")
}
