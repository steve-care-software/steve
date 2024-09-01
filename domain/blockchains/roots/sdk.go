package roots

import (
	"github.com/steve-care-software/steve/domain/blockchains/roots/units"
	"github.com/steve-care-software/steve/domain/hash"
)

// Root represents a root block
type Root interface {
	Hash() hash.Hash
	Database() hash.Hash
	Units() units.Unit
}
