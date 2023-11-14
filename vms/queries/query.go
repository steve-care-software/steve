package queries

import (
	vm_bytes "github.com/steve-care-software/steve/vms/bytes"
	"github.com/steve-care-software/steve/vms/queries/programs"
	"github.com/steve-care-software/steve/vms/queries/results"
	vm_layers "github.com/steve-care-software/steve/vms/queries/scopes/layers"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/frames"
)

type query struct {
	vmBytes       vm_bytes.Bytes
	vmLayer       vm_layers.Layer
	resultBuilder results.Builder
	queryBuilder  results.QueryBuilder
}

func createQuery(
	vmBytes vm_bytes.Bytes,
	vmLayer vm_layers.Layer,
	resultBuilder results.Builder,
	queryBuilder results.QueryBuilder,
) Query {
	out := query{
		vmBytes:       vmBytes,
		vmLayer:       vmLayer,
		resultBuilder: resultBuilder,
		queryBuilder:  queryBuilder,
	}

	return &out
}

// Execute executes the query
func (app *query) Execute(input programs.Program, frame frames.Frame) (results.Result, error) {
	layerName := input.Layer()
	assignment, err := frame.Fetch(layerName)
	if err != nil {
		return nil, err
	}

	assignable := assignment.Assignable()
	if !assignable.IsLayer() {
		// error
	}

	layer := assignable.Layer()
	execLayer, err := app.vmLayer.Execute(layer, frame)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
