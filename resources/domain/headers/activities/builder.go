package activities

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/hash"
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits"
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

	_, err := app.commits.Fetch(app.head)
	if err != nil {
		str := fmt.Sprintf("the head commit (hash: %s) could not be found in the commits", app.head.String())
		return nil, errors.New(str)
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
