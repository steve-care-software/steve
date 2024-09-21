package compensations

import (
	"github.com/steve-care-software/steve/domain/hash"
)

// Builder represents the compensation builder
type Builder interface {
	Create() Builder
	WithWrite(write float64) Builder
	WithReview(review float64) Builder
	Now() (Compensation, error)
}

// Compensation represents a compensation
type Compensation interface {
	Hash() hash.Hash
	HasWrite() bool
	Write() float64
	HasReview() bool
	Review() float64
}
