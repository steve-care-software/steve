package applications

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/responses"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Schema saves a schema on the database
func (app *application) Schema(schema schemas.Schema) error {
	return nil
}

// Execute executes a script on the database
func (app *application) Execute(script scripts.Script) (responses.Response, error) {
	return nil, nil
}
