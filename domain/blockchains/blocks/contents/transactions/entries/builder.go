package entries

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	flag        hash.Hash
	script      hash.Hash
	pFees       *uint64
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		flag:        nil,
		script:      nil,
		pFees:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithFlag adds a flag to the builder
func (app *builder) WithFlag(flag hash.Hash) Builder {
	app.flag = flag
	return app
}

// WithScript adds a script to the builder
func (app *builder) WithScript(script hash.Hash) Builder {
	app.script = script
	return app
}

// WithFees adds a fees to the builder
func (app *builder) WithFees(fees uint64) Builder {
	app.pFees = &fees
	return app
}

// Now builds a new Entry
func (app *builder) Now() (Entry, error) {
	if app.flag == nil {
		return nil, errors.New("the flag is mandatory in order to build an Entry instance")
	}

	if app.script == nil {
		return nil, errors.New("the script is mandatory in order to build an Entry instance")
	}

	if app.pFees == nil {
		return nil, errors.New("the fees is mandatory in order to build an Entry instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.flag.Bytes(),
		app.script.Bytes(),
		[]byte(fmt.Sprintf("%d", *app.pFees)),
	})

	if err != nil {
		return nil, err
	}

	return createEntry(
		*pHash,
		app.flag,
		app.script,
		*app.pFees,
	), nil

}
