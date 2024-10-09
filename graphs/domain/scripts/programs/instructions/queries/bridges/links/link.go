package links

import "github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/assignments"

type link struct {
	origin assignments.Assignment
	target assignments.Assignment
}

func createLink(origin, target assignments.Assignment) Link {
	return &link{
		origin: origin,
		target: target,
	}
}

// Origin returns the origin assignment
func (obj *link) Origin() assignments.Assignment {
	return obj.origin
}

// Target returns the target assignment
func (obj *link) Target() assignments.Assignment {
	return obj.target
}
