package weights

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/steve/engine/domain/scripts/components/suites"
	"github.com/steve-care-software/steve/hash"
)

type weightBuilder struct {
	hashAdapter hash.Adapter
	name        string
	pValue      *uint
	reverse     string
	suites      suites.Suites
}

func createWeightBuilder(
	hashAdapter hash.Adapter,
) WeightBuilder {
	out := weightBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		pValue:      nil,
		reverse:     "",
		suites:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *weightBuilder) Create() WeightBuilder {
	return createWeightBuilder(
		app.hashAdapter,
	)
}

// WithName adds a weight to the builder
func (app *weightBuilder) WithName(name string) WeightBuilder {
	app.name = name
	return app
}

// WithValue adds a value to the builder
func (app *weightBuilder) WithValue(value uint) WeightBuilder {
	app.pValue = &value
	return app
}

// WithReverse adds a reverse to the builder
func (app *weightBuilder) WithReverse(reverse string) WeightBuilder {
	app.reverse = reverse
	return app
}

// WithSuites adds a suites to the builder
func (app *weightBuilder) WithSuites(suites suites.Suites) WeightBuilder {
	app.suites = suites
	return app
}

// Now builds a new Weight instance
func (app *weightBuilder) Now() (Weight, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Weight instance")
	}

	if app.pValue == nil {
		return nil, errors.New("the value is mandatory in order to build a Weight instance")
	}

	data := [][]byte{
		[]byte(app.name),
		[]byte(strconv.Itoa(int(*app.pValue))),
	}

	if app.reverse != "" {
		data = append(data, []byte(app.reverse))
	}

	if app.suites != nil {
		data = append(data, []byte(app.suites.Hash().Bytes()))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.reverse != "" && app.suites != nil {
		return createWeightWithReverseAndSuites(
			*pHash,
			app.name,
			*app.pValue,
			app.reverse,
			app.suites,
		), nil
	}

	if app.reverse != "" {
		return createWeightWithReverse(
			*pHash,
			app.name,
			*app.pValue,
			app.reverse,
		), nil
	}

	if app.suites != nil {
		return createWeightWithSuites(
			*pHash,
			app.name,
			*app.pValue,
			app.suites,
		), nil
	}

	return createWeight(*pHash, app.name, *app.pValue), nil
}
