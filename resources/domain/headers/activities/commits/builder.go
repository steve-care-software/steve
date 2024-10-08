package commits

import (
	"errors"

	"github.com/steve-care-software/steve/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Commit
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Commit) Builder {
	app.list = list
	return app
}

// Now builds a new Commits instance
func (app *builder) Now() (Commits, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Commit in order to build a Commits instance")
	}

	data := [][]byte{}
	mp := map[string]Commit{}
	for _, oneCommit := range app.list {
		data = append(data, oneCommit.Hash().Bytes())
		mp[oneCommit.Hash().String()] = oneCommit
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createCommits(
		*pHash,
		app.list,
		mp,
	), nil
}
