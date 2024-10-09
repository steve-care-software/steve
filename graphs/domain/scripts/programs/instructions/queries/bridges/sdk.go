package bridges

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/bridges/links"
)

// Bridges represents bridges
type Bridges interface {
	List() []Bridge
}

// Bridge represents a bridge
type Bridge interface {
	Weight() uint
	Origin() links.Link
	Target() links.Link
}
