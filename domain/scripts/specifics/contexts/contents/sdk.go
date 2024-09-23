package contents

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/contexts/contents/formats"
	"github.com/steve-care-software/steve/domain/scripts/specifics/contexts/contents/weights"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a content builder
type Builder interface {
	Create() Builder
	WithFormats(formats formats.Formats) Builder
	WithWeights(weights weights.Weights) Builder
	Now() (Content, error)
}

// Content represents a context content
type Content interface {
	Hash() hash.Hash
	HasFormats() bool
	Formats() formats.Formats
	HasWeights() bool
	Weights() weights.Weights
}
