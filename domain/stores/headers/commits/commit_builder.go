package commits

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/commits/modifications"
)

type commitBuilder struct {
	hashAdapter   hash.Adapter
	modifications modifications.Modifications
	parent        hash.Hash
}

func createCommitBuilder(
	hashAdapter hash.Adapter,
) CommitBuilder {
	out := commitBuilder{
		hashAdapter:   hashAdapter,
		modifications: nil,
		parent:        nil,
	}

	return &out
}

// Create initializes the commitBuilder
func (app *commitBuilder) Create() CommitBuilder {
	return createCommitBuilder(
		app.hashAdapter,
	)
}

// WithModifications add modifications to the commitBuilder
func (app *commitBuilder) WithModifications(modifications modifications.Modifications) CommitBuilder {
	app.modifications = modifications
	return app
}

// WithParent add parent to the commitBuilder
func (app *commitBuilder) WithParent(parent hash.Hash) CommitBuilder {
	app.parent = parent
	return app
}

// Now builds a new Commit instance
func (app *commitBuilder) Now() (Commit, error) {
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
