package links

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/assignments"
)

type builder struct {
	origin assignments.Assignment
	target assignments.Assignment
}

func createBuilder() Builder {
	out := builder{
		origin: nil,
		target: nil,
	}

	return &out
}

// Create initializes the link builder
func (obj *builder) Create() Builder {
	return createBuilder()
}

// WithOrigin adds an origin assignment to the link builder
func (obj *builder) WithOrigin(origin assignments.Assignment) Builder {
	obj.origin = origin
	return obj
}

// WithTarget adds a target assignment to the link builder
func (obj *builder) WithTarget(target assignments.Assignment) Builder {
	obj.target = target
	return obj
}

// Now builds a new Link instance
func (obj *builder) Now() (Link, error) {
	if obj.origin == nil {
		return nil, errors.New("the origin is mandatory in order to build a Link instance")
	}

	if obj.target == nil {
		return nil, errors.New("the target is mandatory in order to build a Link instance")
	}

	return createLink(obj.origin, obj.target), nil
}
