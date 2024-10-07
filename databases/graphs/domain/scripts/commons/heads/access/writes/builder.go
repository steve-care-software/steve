package writes

import (
	"errors"

	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads/access/permissions"
)

type builder struct {
	modify permissions.Permissions
	review permissions.Permissions
}

func createBuilder() Builder {
	out := builder{
		modify: nil,
		review: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithModify adds a modify to the builder
func (app *builder) WithModify(modify permissions.Permissions) Builder {
	app.modify = modify
	return app
}

// WithReview adds a review to the builder
func (app *builder) WithReview(review permissions.Permissions) Builder {
	app.review = review
	return app
}

// Now builds a new Write instance
func (app *builder) Now() (Write, error) {
	if app.modify == nil {
		return nil, errors.New("the modify permissions is mandatory in order to build a Write instance")
	}

	if app.review != nil {
		return createWriteWithReview(app.modify, app.review), nil
	}

	return createWrite(app.modify), nil
}
