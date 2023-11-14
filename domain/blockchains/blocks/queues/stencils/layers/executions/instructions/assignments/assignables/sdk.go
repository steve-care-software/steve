package assignables

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/executions/instructions/assignments/assignables/compares"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/executions/queries"
)

// Builder represents the assignable builder
type Builder interface {
	Create() Builder
	WithCompare(compare compares.Compare) Builder
	WithQuery(query queries.Query) Builder
	Now() (Assignable, error)
}

// Assignable represents the assignable
type Assignable interface {
	IsCompare() bool
	Compare() compares.Compare
	IsQuery() bool
	Query() queries.Query
}
