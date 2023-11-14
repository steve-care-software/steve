package bytevalues

import "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/bytevalues/failures"

// Builder represents the assignables builder
type Builder interface {
	Create() Builder
	WithList(list []ByteValue) Builder
	Now() (ByteValues, error)
}

// ByteValues represents the byteValues
type ByteValues interface {
	List() []ByteValue
	IsSuccess() bool
	Compare() (bool, error)
}

// ByteValueBuilder represents the byteValue builder
type ByteValueBuilder interface {
	Create() ByteValueBuilder
	WithSuccess(success []byte) ByteValueBuilder
	WithFailure(failure failures.Failure) ByteValueBuilder
	Now() (ByteValue, error)
}

// ByteValue represents the byteValue
type ByteValue interface {
	IsSuccess() bool
	Success() []byte
	IsFailure() bool
	Failure() failures.Failure
}
