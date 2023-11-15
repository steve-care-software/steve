package queries

import (
	bytes_applications "github.com/steve-care-software/steve/vms/bytes/applications"
	"github.com/steve-care-software/steve/vms/queries/programs"
	"github.com/steve-care-software/steve/vms/queries/results"
)

type query struct {
	bytesApp      bytes_applications.Application
	resultBuilder results.Builder
}

func createQuery(
	bytesApp bytes_applications.Application,
	resultBuilder results.Builder,
) Query {
	out := query{
		bytesApp:      bytesApp,
		resultBuilder: resultBuilder,
	}

	return &out
}

// Execute executes the query
func (app *query) Execute(input programs.Program) (results.Result, error) {
	return nil, nil
}
