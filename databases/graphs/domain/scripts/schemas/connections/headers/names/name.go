package names

import "github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/headers/names/cardinalities"

type name struct {
	name        string
	cardinality cardinalities.Cardinality
}

func createName(
	nameValue string,
	cardinality cardinalities.Cardinality,
) Name {
	out := name{
		name:        nameValue,
		cardinality: cardinality,
	}

	return &out
}

// Name returns the name
func (obj *name) Name() string {
	return obj.name
}

// Cardinality returns the cardinality
func (obj *name) Cardinality() cardinalities.Cardinality {
	return obj.cardinality
}
