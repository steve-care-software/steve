package references

import (
	"errors"
	"fmt"
	"time"

	"github.com/steve-care-software/steve/domain/hash"
)

type commitBuilder struct {
	hashAdapter hash.Adapter
	action      Action
	pParent     *hash.Hash
	pCreatedOn  *time.Time
}

func createCommitBuilder(
	hashAdapter hash.Adapter,
) CommitBuilder {
	out := commitBuilder{
		hashAdapter: hashAdapter,
		action:      nil,
		pParent:     nil,
		pCreatedOn:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *commitBuilder) Create() CommitBuilder {
	return createCommitBuilder(
		app.hashAdapter,
	)
}

// WithAction adds an action to the builder
func (app *commitBuilder) WithAction(action Action) CommitBuilder {
	app.action = action
	return app
}

// WithParent adds a parent to the builder
func (app *commitBuilder) WithParent(parent hash.Hash) CommitBuilder {
	app.pParent = &parent
	return app
}

// CreatedOn adds a creation time to the builder
func (app *commitBuilder) CreatedOn(createdOn time.Time) CommitBuilder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Commit instance
func (app *commitBuilder) Now() (Commit, error) {
	if app.action == nil {
		return nil, errors.New("the action is mandatory in order to build a Commit instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Commit instance")
	}

	data := [][]byte{
		app.action.Hash().Bytes(),
		[]byte(fmt.Sprintf("%d", app.pCreatedOn.UnixNano())),
	}

	if app.pParent != nil {
		data = append(data, app.pParent.Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.pParent != nil {
		return createCommitWithParent(*pHash, app.action, *app.pCreatedOn, app.pParent), nil
	}

	return createCommit(*pHash, app.action, *app.pCreatedOn), nil
}
