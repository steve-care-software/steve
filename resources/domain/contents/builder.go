package contents

import (
	"errors"

	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits/modifications"
)

type builder struct {
	modification modifications.Modification
	data         []byte
}

func createBuilder() Builder {
	out := builder{
		modification: nil,
		data:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithModification adds a modification to the builder
func (app *builder) WithModification(modification modifications.Modification) Builder {
	app.modification = modification
	return app
}

// WithData adds data to the builder
func (app *builder) WithData(data []byte) Builder {
	app.data = data
	return app
}

// Now builds a new Content instance
func (app *builder) Now() (Content, error) {
	if app.modification == nil {
		return nil, errors.New("the modification is mandatory in order to build a Content instance")
	}

	if app.data != nil && len(app.data) <= 0 {
		app.data = nil
	}

	if app.data == nil {
		return nil, errors.New("the data is mandatory in order to build a Content instance")
	}

	return createContent(
		app.modification,
		app.data,
	), nil
}
