package commits

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/commits/modifications"
)

type builder struct {
	hashAdapter   hash.Adapter
	modifications modifications.Modifications
	parent        hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:   hashAdapter,
		modifications: nil,
		parent:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithModifications add modifications to the builder
func (app *builder) WithModifications(modifications modifications.Modifications) Builder {
	app.modifications = modifications
	return app
}

// WithParent add parent to the builder
func (app *builder) WithParent(parent hash.Hash) Builder {
	app.parent = parent
	return app
}

// Now builds a new Commit instance
func (app *builder) Now() (Commit, error) {
	if app.modifications == nil {
		return nil, errors.New("the modifications is mandatory in order to build a Commit instance")
	}

	data := [][]byte{
		app.modifications.Hash().Bytes(),
	}

	if app.parent != nil {
		data = append(data, app.parent.Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.parent != nil {
		return createCommitWithParent(*pHash, app.modifications, app.parent), nil
	}

	return createCommit(*pHash, app.modifications), nil
}
