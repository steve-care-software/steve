package operations

type operation struct {
	actor Actor
	tail  Tail
	isNot bool
}

func createOperation(
	actor Actor,
	tail Tail,
	isNot bool,
) Operation {
	out := operation{
		actor: actor,
		tail:  tail,
		isNot: isNot,
	}

	return &out
}

// Actor returns the actor
func (obj *operation) Actor() Actor {
	return obj.actor
}

// Tail returns the tail
func (obj *operation) Tail() Tail {
	return obj.tail
}

// IsNot returns true if not, false otherwise
func (obj *operation) IsNot() bool {
	return obj.isNot
}
