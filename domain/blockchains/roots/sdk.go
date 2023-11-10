package roots

import (
	"time"

	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/blockchains/roots/resolutions"
)

// Root represents a root block
type Root interface {
	Hash() hash.Hash
	Resolution() resolutions.Resolution
	Owner() hash.Hash
	CreatedOn() time.Time
}
