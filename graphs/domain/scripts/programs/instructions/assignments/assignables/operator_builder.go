package assignables

import (
	"errors"
)

type operatorBuilder struct {
	arithmetic *uint8
	relational *uint8
	equal      *uint8
	logical    *uint8
}

func createOperatorBuilder() OperatorBuilder {
	return &operatorBuilder{
		arithmetic: nil,
		relational: nil,
		equal:      nil,
		logical:    nil,
	}
}

// Create initializes the operator builder
func (app *operatorBuilder) Create() OperatorBuilder {
	return createOperatorBuilder()
}

// WithArithmetic adds an arithmetic operator to the builder
func (app *operatorBuilder) WithArithmetic(arithmetic uint8) OperatorBuilder {
	app.arithmetic = &arithmetic
	return app
}

// WithRelational adds a relational operator to the builder
func (app *operatorBuilder) WithRelational(relational uint8) OperatorBuilder {
	app.relational = &relational
	return app
}

// WithEqual adds an equal operator to the builder
func (app *operatorBuilder) WithEqual(equal uint8) OperatorBuilder {
	app.equal = &equal
	return app
}

// WithLogical adds a logical operator to the builder
func (app *operatorBuilder) WithLogical(logical uint8) OperatorBuilder {
	app.logical = &logical
	return app
}

// Now builds a new Operator instance
func (app *operatorBuilder) Now() (Operator, error) {
	if app.arithmetic != nil {
		return createOperatorWithArithmetic(app.arithmetic), nil
	}

	if app.relational != nil {
		return createOperatorWithRelational(app.relational), nil
	}

	if app.equal != nil {
		return createOperatorWithEqual(app.equal), nil
	}

	if app.logical != nil {
		return createOperatorWithLogical(app.logical), nil
	}

	return nil, errors.New("the Operator is invalid")
}
