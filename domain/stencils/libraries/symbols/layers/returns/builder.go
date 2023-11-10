package returns

import (
	"errors"

	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns/kinds"
)

type builder struct {
	hashAdapter hash.Adapter
	output      []byte
	kind        kinds.Kind
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		output:      nil,
		kind:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithOutput adds an output to the builder
func (app *builder) WithOutput(output []byte) Builder {
	app.output = output
	return app
}

// WithKind adds a kind to the builder
func (app *builder) WithKind(kind kinds.Kind) Builder {
	app.kind = kind
	return app
}

// Now builds a new Return instance
func (app *builder) Now() (Return, error) {
	if app.output != nil && len(app.output) <= 0 {
		app.output = nil
	}

	if app.output == nil {
		return nil, errors.New("the output is mandatory in order to build a Return return instance")
	}

	if app.kind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Return return instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.output,
		app.kind.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createReturn(*pHash, app.output, app.kind), nil
}
