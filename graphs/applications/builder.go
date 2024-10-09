package applications

import (
	"errors"

	"github.com/steve-care-software/steve/hash"
	application_lists "github.com/steve-care-software/steve/lists/applications"
	application_resources "github.com/steve-care-software/steve/resources/applications"
)

type builder struct {
	storeListApp application_lists.Application
	resourceApp  application_resources.Application
	hashAdapter  hash.Adapter
	dbIdentifier string
}

func createBuilder(
	storeListApp application_lists.Application,
	resourceApp application_resources.Application,
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		storeListApp: storeListApp,
		resourceApp:  resourceApp,
		hashAdapter:  hashAdapter,
		dbIdentifier: "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.storeListApp,
		app.resourceApp,
		app.hashAdapter,
	)
}

// WithIdentifier adds an identitifer to the builder
func (app *builder) WithIdentifier(dbIdentifier string) Builder {
	app.dbIdentifier = dbIdentifier
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.dbIdentifier == "" {
		return nil, errors.New("the dbIdentifier is mandatory in order to build an Application instance")
	}

	return createApplication(
		app.storeListApp,
		app.resourceApp,
		app.hashAdapter,
		app.dbIdentifier,
	), nil
}
