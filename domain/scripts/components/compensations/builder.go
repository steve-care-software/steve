package compensations

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	pWrite      *float64
	pReview     *float64
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		pWrite:      nil,
		pReview:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithWrite adds a write to the builder
func (app *builder) WithWrite(write float64) Builder {
	app.pWrite = &write
	return app
}

// WithReview adds a review to the builder
func (app *builder) WithReview(review float64) Builder {
	app.pReview = &review
	return app
}

// Now builds a new Compensation instance
func (app *builder) Now() (Compensation, error) {
	data := [][]byte{}
	if app.pWrite != nil {
		data = append(data, []byte(fmt.Sprintf("%f", *app.pWrite)))
	}

	if app.pReview != nil {
		data = append(data, []byte(fmt.Sprintf("%f", *app.pReview)))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.pWrite != nil && app.pReview != nil {
		return createCompensationWithWriteAndReview(
			*pHash,
			app.pWrite,
			app.pReview,
		), nil
	}

	if app.pWrite != nil {
		return createCompensationWithWrite(
			*pHash,
			app.pWrite,
		), nil
	}

	if app.pReview != nil {
		return createCompensationWithReview(
			*pHash,
			app.pReview,
		), nil
	}

	return nil, errors.New("the Compensation is invalid")
}
