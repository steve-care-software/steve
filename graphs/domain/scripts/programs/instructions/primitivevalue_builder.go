package instructions

import (
	"errors"
)

type primitiveValueBuilder struct {
	numeric   NumericValue
	boolValue *bool
	strValue  *string
}

func createPrimitiveValueBuilder() PrimitiveValueBuilder {
	return &primitiveValueBuilder{
		numeric:   nil,
		boolValue: nil,
		strValue:  nil,
	}
}

// Create initializes the primitive value builder
func (app *primitiveValueBuilder) Create() PrimitiveValueBuilder {
	return createPrimitiveValueBuilder()
}

// WithNumeric adds a numeric value to the builder
func (app *primitiveValueBuilder) WithNumeric(numeric NumericValue) PrimitiveValueBuilder {
	app.numeric = numeric
	return app
}

// WithBool adds a boolean value to the builder
func (app *primitiveValueBuilder) WithBool(boolValue bool) PrimitiveValueBuilder {
	app.boolValue = &boolValue
	return app
}

// WithString adds a string value to the builder
func (app *primitiveValueBuilder) WithString(strValue string) PrimitiveValueBuilder {
	app.strValue = &strValue
	return app
}

// Now builds a new PrimitiveValue instance
func (app *primitiveValueBuilder) Now() (PrimitiveValue, error) {
	if app.numeric != nil {
		return createPrimitiveValueWithNumeric(app.numeric), nil
	}

	if app.boolValue != nil {
		return createPrimitiveValueWithBool(app.boolValue), nil
	}

	if app.strValue != nil {
		return createPrimitiveValueWithString(app.strValue), nil
	}

	return nil, errors.New("the PrimitiveValue is invalid")
}
