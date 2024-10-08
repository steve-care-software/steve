package formats

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/contexts/contents/formats/suites"
	"github.com/steve-care-software/steve/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewFormatBuilder creates a format builder
func NewFormatBuilder() FormatBuilder {
	hashAdapter := hash.NewAdapter()
	return createFormatBuilder(
		hashAdapter,
	)
}

// Builder represents a formats builder
type Builder interface {
	Create() Builder
	WithList(list []Format) Builder
	Now() (Formats, error)
}

// Formats represents formats
type Formats interface {
	Hash() hash.Hash
	List() []Format
}

// FormatBuilder represents a format builder
type FormatBuilder interface {
	Create() FormatBuilder
	WithPoint(point string) FormatBuilder
	WithGrammar(grammar string) FormatBuilder
	WithSuites(suites suites.Suites) FormatBuilder
	Now() (Format, error)
}

// Format represents a format
type Format interface {
	Hash() hash.Hash
	Point() string
	Grammar() string
	HasSuites() bool
	Suites() suites.Suites
}
