package parameters

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/containers"
)

// Parameters represents func parameters
type Parameters interface {
	Hash() hash.Hash
	List() []Parameter
}

// Parameter represents a func parameter
type Parameter interface {
	Hash() hash.Hash
	Name() string
	Container() containers.Container
	IsMandatory() bool
}
