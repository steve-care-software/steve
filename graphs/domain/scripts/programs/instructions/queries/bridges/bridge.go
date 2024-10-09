package bridges

import "github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/bridges/links"

// bridge represents the implementation of the Bridge interface
type bridge struct {
	weight uint
	origin links.Link
	target links.Link
}

func createBridge(weight uint, origin, target links.Link) Bridge {
	return &bridge{
		weight: weight,
		origin: origin,
		target: target,
	}
}

// Weight returns the weight of the bridge
func (obj *bridge) Weight() uint {
	return obj.weight
}

// Origin returns the origin link of the bridge
func (obj *bridge) Origin() links.Link {
	return obj.origin
}

// Target returns the target link of the bridge
func (obj *bridge) Target() links.Link {
	return obj.target
}
