package selects

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/conditions"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references/externals"
)

type builder struct {
	externals externals.Externals
	condition conditions.Condition
	isDelete  bool
}

func createBuilder() Builder {
	return &builder{
		externals: nil,
		condition: nil,
		isDelete:  false,
	}
}

// Create initializes the select builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithExternals adds externals to the select builder
func (app *builder) WithExternals(externals externals.Externals) Builder {
	app.externals = externals
	return app
}

// WithCondition adds a condition to the select builder
func (app *builder) WithCondition(condition conditions.Condition) Builder {
	app.condition = condition
	return app
}

// IsDelete marks the select as a delete operation
func (app *builder) IsDelete() Builder {
	app.isDelete = true
	return app
}

// Now builds a new Select instance
func (app *builder) Now() (Select, error) {
	if app.externals == nil {
		return nil, errors.New("the externals are mandatory in order to build a Select instance")
	}

	if app.condition != nil {
		return createSelectWithCondition(app.externals, app.isDelete, app.condition), nil
	}

	return createSelect(app.externals, app.isDelete), nil
}
