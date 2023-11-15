package queries

import (
	"github.com/steve-care-software/steve/vms/children/queries/programs"
	"github.com/steve-care-software/steve/vms/children/queries/results"
)

// Query represents the query
type Query interface {
	Execute(input programs.Program) (results.Result, error)
}
