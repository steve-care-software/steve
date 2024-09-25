package functions

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/calls/functions/parameters"
)

// Function represents a func call
type Function interface {
	Hash() hash.Hash
	Name() string
	Parameters() parameters.Parameters
	IsEngine() bool
}
