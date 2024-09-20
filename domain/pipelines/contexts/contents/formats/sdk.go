package formats

import (
	"github.com/steve-care-software/steve/domain/hash"
)

// Formats represents formats
type Formats interface {
	List() []Format
}

// Format represents a format
type Format interface {
	Point() string
	Grammar() hash.Hash
}
