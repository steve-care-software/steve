package preparations

import (
	"errors"

	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/pointers"
)

type preparationBuilder struct {
	hashAdapter hash.Adapter
	isStop      bool
	load        pointers.Pointer
	exists      pointers.Pointer
	condition   Condition
}

func createPreparationBuilder(
	hashAdapter hash.Adapter,
) PreparationBuilder {
	out := preparationBuilder{
		hashAdapter: hashAdapter,
		isStop:      false,
		load:        nil,
		exists:      nil,
		condition:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *preparationBuilder) Create() PreparationBuilder {
	return createPreparationBuilder(
		app.hashAdapter,
	)
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

	data := [][]byte{}
	if app.load != nil {
		data = append(data, app.load.Hash().Bytes())
	}

	if app.exists != nil {
		data = append(data, app.exists.Hash().Bytes())
	}

	if app.condition != nil {
		data = append(data, app.condition.Hash().Bytes())
	}

	if app.isStop {
		data = append(data, []byte{0})
	}

	if len(data) <= 0 {
		return nil, errors.New("the Preparation is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.load != nil {
		return createPreparationWithLoad(*pHash, app.load), nil
	}

	if app.exists != nil {
		return createPreparationWithExists(*pHash, app.exists), nil
	}

	if app.condition != nil {
		return createPreparationWithCondition(*pHash, app.condition), nil
	}

	return createPreparationWithStop(*pHash), nil
}
