package contents

import (
	"github.com/steve-care-software/steve/domain/scripts/contexts/contents/formats"
	"github.com/steve-care-software/steve/domain/scripts/contexts/contents/weights"
)

// Content represents a context content
type Content interface {
	HasFormats() bool
	Formats() formats.Formats
	HasWeights() bool
	Weights() weights.Weights
}
