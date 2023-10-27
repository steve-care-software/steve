package libraries

import (
	"github.com/steve-care-software/steve/domain/stencils/libraries"
	"github.com/steve-care-software/steve/domain/stencils/libraries/results"
)

type application struct {
	service libraries.Service
}

func createApplication(
	service libraries.Service,
) Application {
	out := application{
		service: service,
	}

	return &out
}

// Save saves a library to the database
func (app *application) Save(library libraries.Library) (results.Result, error) {
	return app.service.Save(library)
}
