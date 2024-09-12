package lists

import (
	"errors"

	"github.com/steve-care-software/steve/applications/resources"
	"github.com/steve-care-software/steve/domain/stores/lists"
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
	if app.resourceApp != nil {
		return nil, errors.New("the resource application is mandatory in order to build an APplication instance")
	}

	return createApplication(
		app.resourceApp,
		app.listAdapter,
	), nil
}
