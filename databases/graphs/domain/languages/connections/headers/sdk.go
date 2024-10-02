package headers

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/headers/names"
)

// Header represents an header
type Header interface {
	Name() names.Name
	HasReverse() bool
	Reverse() names.Name
}
