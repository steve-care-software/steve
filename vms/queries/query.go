package queries

import (
	vm_bytes "github.com/steve-care-software/steve/vms/bytes"
	"github.com/steve-care-software/steve/vms/queries/programs"
	"github.com/steve-care-software/steve/vms/queries/results"
)

type query struct {
	vmBytes       vm_bytes.Bytes
	resultBuilder results.Builder
}

func createQuery(
	vmBytes vm_bytes.Bytes,
	resultBuilder results.Builder,
) Query {
	out := query{
		vmBytes:       vmBytes,
		resultBuilder: resultBuilder,
	}

	return &out
}

// Execute executes the query
func (app *query) Execute(input programs.Program) (results.Result, error) {
	return nil, nil
}
