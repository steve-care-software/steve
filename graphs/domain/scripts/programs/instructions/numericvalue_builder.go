package instructions

import (
	"errors"
)

type numericValueBuilder struct {
	flValue  *float64
	uiValue  *uint64
	intValue *int64
}

func createNumericValueBuilder() NumericValueBuilder {
	return &numericValueBuilder{
		flValue:  nil,
		uiValue:  nil,
		intValue: nil,
	}
}

// Create initializes the numeric value builder
func (app *numericValueBuilder) Create() NumericValueBuilder {
	return createNumericValueBuilder()
}

// WithFloat adds a float value to the numeric value builder
func (app *numericValueBuilder) WithFloat(flValue float64) NumericValueBuilder {
	app.flValue = &flValue
	return app
}

// WithUint adds an unsigned integer value to the numeric value builder
func (app *numericValueBuilder) WithUint(uiValue uint64) NumericValueBuilder {
	app.uiValue = &uiValue
	return app
}

// WithInt adds an integer value to the numeric value builder
func (app *numericValueBuilder) WithInt(intValue int64) NumericValueBuilder {
	app.intValue = &intValue
	return app
}

// Now builds a new NumericValue instance
func (app *numericValueBuilder) Now() (NumericValue, error) {
	if app.flValue != nil {
		return createNumericValueWithFloat(app.flValue), nil
	}

	if app.uiValue != nil {
		return createNumericValueWithUint(app.uiValue), nil
	}

	if app.intValue != nil {
		return createNumericValueWithInt(app.intValue), nil
	}

	return nil, errors.New("the NumericValue is invalid")
}
