package references

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/grammars/constants/tokens/elements/references/values"
	"github.com/steve-care-software/steve/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	grammar     string
	value       values.Value
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		grammar:     "",
		value:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithGrammar adds a grammar to the builder
func (app *builder) WithGrammar(grammar string) Builder {
	app.grammar = grammar
	return app
}

// WithValue adds a value to the builder
func (app *builder) WithValue(value values.Value) Builder {
	app.value = value
	return app
}

// Now builds a new Reference instance
func (app *builder) Now() (Reference, error) {
	if app.grammar == "" {
		return nil, errors.New("the grammar is mandatory in order to build a Reference instance")
	}

	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build a Reference instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.grammar),
		app.value.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createReference(
		*pHash,
		app.grammar,
		app.value,
	), nil
}
