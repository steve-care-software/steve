package roots

import (
	"time"

	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/blockchains/roots/resolutions"
)

// Builder represents a root builder
type Builder interface {
	Create() Builder
	WithResolution(resolution resolutions.Resolution) Builder
	WithOwner(owner hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Root, error)
}

// Root represents a root block
type Root interface {
	Hash() hash.Hash
	Resolution() resolutions.Resolution
	CreatedOn() time.Time
	HasOwner() bool
	Owner() hash.Hash
}
