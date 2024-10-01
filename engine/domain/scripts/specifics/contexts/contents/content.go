package contents

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/contexts/contents/formats"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/contexts/contents/weights"
)

type content struct {
	hash    hash.Hash
	formats formats.Formats
	weights weights.Weights
}

func createContentWithFormats(
	hash hash.Hash,
	formats formats.Formats,
) Content {
	return createContentInternally(hash, formats, nil)
}

func createContentWithWeights(
	hash hash.Hash,
	weights weights.Weights,
) Content {
	return createContentInternally(hash, nil, weights)
}

func createContentWithFormatsAndWeights(
	hash hash.Hash,
	formats formats.Formats,
	weights weights.Weights,
) Content {
	return createContentInternally(hash, formats, weights)
}

func createContentInternally(
	hash hash.Hash,
	formats formats.Formats,
	weights weights.Weights,
) Content {
	out := content{
		hash:    hash,
		formats: formats,
		weights: weights,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// HasFormats returns true if there is formats, false otherwise
func (obj *content) HasFormats() bool {
	return obj.formats != nil
}

// Formats returns the formats, if any
func (obj *content) Formats() formats.Formats {
	return obj.formats
}

// HasWeights returns true if there is weights, false otherwise
func (obj *content) HasWeights() bool {
	return obj.weights != nil
}

// Weights returns the weights, if any
func (obj *content) Weights() weights.Weights {
	return obj.weights
}
