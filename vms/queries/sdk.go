package queries

import (
	"github.com/steve-care-software/steve/vms/queries/programs"
	"github.com/steve-care-software/steve/vms/queries/results"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/frames"
)

// Query represents the query
type Query interface {
	Execute(input programs.Program, frame frames.Frame) (results.Result, error)
}
