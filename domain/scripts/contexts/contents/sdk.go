package contents

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/contexts/contents/formats"
	"github.com/steve-care-software/steve/domain/scripts/contexts/contents/weights"
)

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
