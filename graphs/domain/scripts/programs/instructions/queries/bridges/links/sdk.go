package links

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/assignments"
)

// Link represents a bridge link
type Link interface {
	Origin() assignments.Assignment
	Target() assignments.Assignment
}
