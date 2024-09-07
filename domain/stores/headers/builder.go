package headers

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/activities"
	"github.com/steve-care-software/steve/domain/stores/headers/commits/modifications/resources"
)

type builder struct {
	hashAdapter hash.Adapter
	root        resources.Resources
	activity    activities.Activity
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		root:        nil,
		activity:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root resources.Resources) Builder {
	app.root = root
	return app
}

// WithActivity adds an activity to the builder
func (app *builder) WithActivity(activity activities.Activity) Builder {
	app.activity = activity
	return app
}

// Now builds a new Header instance
func (app *builder) Now() (Header, error) {
	if app.root == nil {
		return nil, errors.New("the root is mandatory in order to build an Header instance")
	}

	data := [][]byte{
		app.root.Hash().Bytes(),
	}

	if app.activity != nil {
		data = append(data, app.activity.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.activity != nil {
		return createHeaderWithActivity(
			*pHash,
			app.root,
			app.activity,
		), nil
	}

	return createHeader(
		*pHash,
		app.root,
	), nil
}
