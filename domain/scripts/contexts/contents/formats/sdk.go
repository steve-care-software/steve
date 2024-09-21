package formats

import (
	"github.com/steve-care-software/steve/domain/scripts/contexts/contents/suites"
)

// Formats represents formats
type Formats interface {
	List() []Format
}

// Format represents a format
type Format interface {
	Point() string
	Grammar() string
	HasSuites() bool
	Suites() suites.Suites
}
