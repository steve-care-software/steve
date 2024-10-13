package instructions

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/assignments/assignables"
)

type forKeyValueBuilder struct {
	key          string
	value        string
	iterable     assignables.Iterable
	instructions ForInstructions
}

func createForKeyValueBuilder() ForKeyValueBuilder {
	return &forKeyValueBuilder{
		key:          "",
		value:        "",
		iterable:     nil,
		instructions: nil,
	}
}

// Create initializes the builder
func (app *forKeyValueBuilder) Create() ForKeyValueBuilder {
	return createForKeyValueBuilder()
}

// WithKey adds a key to the builder
func (app *forKeyValueBuilder) WithKey(key string) ForKeyValueBuilder {
	app.key = key
	return app
}

// WithValue adds a value to the builder
func (app *forKeyValueBuilder) WithValue(value string) ForKeyValueBuilder {
	app.value = value
	return app
}

// WithIterable adds an iterable to the builder
func (app *forKeyValueBuilder) WithIterable(iterable assignables.Iterable) ForKeyValueBuilder {
	app.iterable = iterable
	return app
}

// WithInstructions adds instructions to the builder
func (app *forKeyValueBuilder) WithInstructions(instructions ForInstructions) ForKeyValueBuilder {
	app.instructions = instructions
	return app
}

// Now builds and returns a ForKeyValue instance
func (app *forKeyValueBuilder) Now() (ForKeyValue, error) {
	if app.key == "" {
		return nil, errors.New("the key is mandatory to build a ForKeyValue instance")
	}

	if app.value == "" {
		return nil, errors.New("the value is mandatory to build a ForKeyValue instance")
	}

	if app.iterable == nil {
		return nil, errors.New("an iterable is mandatory to build a ForKeyValue instance")
	}

	if app.instructions == nil {
		return nil, errors.New("instructions are mandatory to build a ForKeyValue instance")
	}

	return createForKeyValue(app.key, app.value, app.iterable, app.instructions), nil
}
