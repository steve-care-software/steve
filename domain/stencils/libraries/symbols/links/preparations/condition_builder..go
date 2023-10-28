package preparations

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type conditionBuilder struct {
	hashAdapter  hash.Adapter
	variable     string
	preparations Preparations
}

func createConditionBuilder(
	hashAdapter hash.Adapter,
) ConditionBuilder {
	out := conditionBuilder{
		hashAdapter:  hashAdapter,
		variable:     "",
		preparations: nil,
	}

	return &out
}

// Create initializes the builder
func (app *conditionBuilder) Create() ConditionBuilder {
	return createConditionBuilder(
		app.hashAdapter,
	)
}

// WithVariable adds a variable to the builder
func (app *conditionBuilder) WithVariable(variable string) ConditionBuilder {
	app.variable = variable
	return app
}

// WithPreparations add preparations to the builder
func (app *conditionBuilder) WithPreparations(preparations Preparations) ConditionBuilder {
	app.preparations = preparations
	return app
}

// Now builds a new Condition instance
func (app *conditionBuilder) Now() (Condition, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a Condition instance")
	}

	if app.preparations == nil {
		return nil, errors.New("the preparations is mandatory in order to build a Condition instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.variable),
		app.preparations.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createCondition(*pHash, app.variable, app.preparations), nil
}
