package applications

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/responses"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Execute executes a script on the database
func (app *application) Execute(script scripts.Script) (responses.Response, error) {
	return nil, nil
}
