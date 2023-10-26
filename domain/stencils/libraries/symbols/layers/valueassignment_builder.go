package layers

import "errors"

type valueAssignmentBuilder struct {
	name  string
	value Value
}

func createValueAssignmentBuilder() ValueAssignmentBuilder {
	out := valueAssignmentBuilder{
		name:  "",
		value: nil,
	}

	return &out
}

// Create initializes the builder
func (app *valueAssignmentBuilder) Create() ValueAssignmentBuilder {
	return createValueAssignmentBuilder()
}

// WithName adds a name to the builder
func (app *valueAssignmentBuilder) WithName(name string) ValueAssignmentBuilder {
	app.name = name
	return app
}

// WithValue adds a value to the builder
func (app *valueAssignmentBuilder) WithValue(value Value) ValueAssignmentBuilder {
	app.value = value
	return app
}

// Now builds a new ValueAssignment instance
func (app *valueAssignmentBuilder) Now() (ValueAssignment, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a ValueAssignment instance")
	}

	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build a VaueAssignment instance")
	}

	return createValueAssignment(app.name, app.value), nil
}
