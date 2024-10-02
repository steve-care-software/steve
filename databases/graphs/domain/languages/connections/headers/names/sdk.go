package names

import "github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/headers/names/cardinalities"

// Name represents an header name
type Name interface {
	Name() string
	Cardinality() cardinalities.Cardinality
}
