package failures

import "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/bytevalues"

// Builder represents the failure builder
type Builder interface {
	Create() Builder
	WithCouldNotCompare(couldNotCompare bytevalues.ByteValues) Builder
	Now() (Failure, error)
}

// Failure represents a failure
type Failure interface {
	IsCouldNotCompare() bool
	CouldNotCompare() bytevalues.ByteValues
}
