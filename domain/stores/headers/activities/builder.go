package activities

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits"
)

type builder struct {
	hashAdapter hash.Adapter
	commits     commits.Commits
	head        hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		commits:     nil,
		head:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithCommits add commits to the builder
func (app *builder) WithCommits(commits commits.Commits) Builder {
	app.commits = commits
	return app
}

// WithHead adds an head to the builder
func (app *builder) WithHead(head hash.Hash) Builder {
	app.head = head
	return app
}

// Now builds a new Activity instance
func (app *builder) Now() (Activity, error) {
	if app.commits == nil {
		return nil, errors.New("the commits is mandatory in order to build an Activity instance")
	}

	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build an Activity instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.commits.Hash().Bytes(),
		app.head.Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createActivity(
		*pHash,
		app.commits,
		app.head,
	), nil
}
