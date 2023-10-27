package preparations

import (
	"errors"

	"github.com/steve-care-software/steve/domain/pointers"
)

type preparationBuilder struct {
	isStop    bool
	load      pointers.Pointer
	exists    pointers.Pointer
	condition Condition
}

func createPreparationBuilder() PreparationBuilder {
	out := preparationBuilder{
		isStop:    false,
		load:      nil,
		exists:    nil,
		condition: nil,
	}

	return &out
}

// Create initializes the builder
func (app *preparationBuilder) Create() PreparationBuilder {
	return createPreparationBuilder()
}

// WithLoad adds a load to the builder
func (app *preparationBuilder) WithLoad(load pointers.Pointer) PreparationBuilder {
	app.load = load
	return app
}

// WithExists add exists to the builder
func (app *preparationBuilder) WithExists(exists pointers.Pointer) PreparationBuilder {
	app.exists = exists
	return app
}

// WithCondition adds a condition to the builder
func (app *preparationBuilder) WithCondition(condition Condition) PreparationBuilder {
	app.condition = condition
	return app
}

// IsStop flags the builder as stop
func (app *preparationBuilder) IsStop() PreparationBuilder {
	app.isStop = true
	return app
}

// Now builds a new Preparation instance
func (app *preparationBuilder) Now() (Preparation, error) {
	if app.load != nil {
		return createPreparationWithLoad(app.load), nil
	}

	if app.exists != nil {
		return createPreparationWithExists(app.exists), nil
	}

	if app.condition != nil {
		return createPreparationWithCondition(app.condition), nil
	}

	if app.isStop {
		return createPreparationWithStop(), nil
	}

	return nil, errors.New("the Preparation is invalid")
}
