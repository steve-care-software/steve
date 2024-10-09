package bridges

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/bridges/links"
)

type builder struct {
	weight uint
	origin links.Link
	target links.Link
}

func createBuilder() Builder {
	return &builder{
		weight: 0,
		origin: nil,
		target: nil,
	}
}

// Create initializes the bridge builder
func (obj *builder) Create() Builder {
	return createBuilder()
}

// WithWeight adds a weight to the bridge builder
func (obj *builder) WithWeight(weight uint) Builder {
	obj.weight = weight
	return obj
}

// WithOrigin adds an origin link to the bridge builder
func (obj *builder) WithOrigin(origin links.Link) Builder {
	obj.origin = origin
	return obj
}

// WithTarget adds a target link to the bridge builder
func (obj *builder) WithTarget(target links.Link) Builder {
	obj.target = target
	return obj
}

// Now builds a new Bridge instance
func (obj *builder) Now() (Bridge, error) {
	if obj.weight <= 0 {
		return nil, errors.New("the weight must be greater than zero (0) in order to build a Bridge instance")
	}

	if obj.origin == nil {
		return nil, errors.New("the origin is mandatory in order to build a Bridge instance")
	}

	if obj.target == nil {
		return nil, errors.New("the target is mandatory in order to build a Bridge instance")
	}

	return createBridge(obj.weight, obj.origin, obj.target), nil
}
