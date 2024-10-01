package lists

import (
	"errors"

	lists "github.com/steve-care-software/steve/databases/lists/domain"
	resources "github.com/steve-care-software/steve/databases/resources/applications"
)

type builder struct {
	listAdapter lists.Adapter
	resourceApp resources.Application
}

func createBuilder(
	listAdapter lists.Adapter,
) Builder {
	out := builder{
		listAdapter: listAdapter,
		resourceApp: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.listAdapter,
	)
}

// WithResource add resource to the builder
func (app *builder) WithResource(resource resources.Application) Builder {
	app.resourceApp = resource
	return app
}

// Now builds a new Application
func (app *builder) Now() (Application, error) {
	if app.resourceApp == nil {
		return nil, errors.New("the resource application is mandatory in order to build an APplication instance")
	}

	return createApplication(
		app.resourceApp,
		app.listAdapter,
	), nil
}
