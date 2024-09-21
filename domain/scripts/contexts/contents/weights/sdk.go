package weights

import (
	"github.com/steve-care-software/steve/domain/scripts/contexts/contents/suites"
)

// Weights represents weights
type Weights interface {
	List() []Weight
}

// Weight represents a weight
type Weight interface {
	Name() string
	Value() uint
	HasReverse() bool
	Reverse() string
	HasSuites() bool
	Suites() suites.Suites
}
