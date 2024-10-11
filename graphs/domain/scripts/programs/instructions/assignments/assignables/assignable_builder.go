package assignables

import "errors"

type assignableBuilder struct {
	engine      AssignableEngine
	listMap     ListMap
	programCall ProgramCall
	primitive   PrimitiveValue
	casting     Casting
	expand      Iterable
	operation   Operation
}

func createAssignableBuilder() AssignableBuilder {
	return &assignableBuilder{
		engine:      nil,
		listMap:     nil,
		programCall: nil,
		primitive:   nil,
		casting:     nil,
		expand:      nil,
		operation:   nil,
	}
}

// Create initializes the AssignableBuilder
func (app *assignableBuilder) Create() AssignableBuilder {
	return createAssignableBuilder()
}

// WithEngine adds an engine to the builder
func (app *assignableBuilder) WithEngine(engine AssignableEngine) AssignableBuilder {
	app.engine = engine
	return app
}

// WithListMap adds a list map to the builder
func (app *assignableBuilder) WithListMap(listMap ListMap) AssignableBuilder {
	app.listMap = listMap
	return app
}

// WithProgramCall adds a program call to the builder
func (app *assignableBuilder) WithProgramCall(programCall ProgramCall) AssignableBuilder {
	app.programCall = programCall
	return app
}

// WithPrimitive adds a primitive value to the builder
func (app *assignableBuilder) WithPrimitive(primitive PrimitiveValue) AssignableBuilder {
	app.primitive = primitive
	return app
}

// WithCasting adds a casting to the builder
func (app *assignableBuilder) WithCasting(casting Casting) AssignableBuilder {
	app.casting = casting
	return app
}

// WithExpand adds an iterable to the builder
func (app *assignableBuilder) WithExpand(expand Iterable) AssignableBuilder {
	app.expand = expand
	return app
}

// WithOperation adds an operation to the builder
func (app *assignableBuilder) WithOperation(operation Operation) AssignableBuilder {
	app.operation = operation
	return app
}

// Now builds a new Assignable instance
func (app *assignableBuilder) Now() (Assignable, error) {
	if app.engine != nil {
		return createAssignableWithEngine(app.engine), nil
	}

	if app.listMap != nil {
		return createAssignableWithListMap(app.listMap), nil
	}

	if app.programCall != nil {
		return createAssignableWithProgramCall(app.programCall), nil
	}

	if app.primitive != nil {
		return createAssignableWithPrimitive(app.primitive), nil
	}

	if app.casting != nil {
		return createAssignableWithCasting(app.casting), nil
	}

	if app.expand != nil {
		return createAssignableWithExpand(app.expand), nil
	}

	if app.operation != nil {
		return createAssignableWithOperation(app.operation), nil
	}

	return nil, errors.New("the Assignable is invalid")
}
