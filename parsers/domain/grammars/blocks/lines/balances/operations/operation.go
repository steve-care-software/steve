package operations

type operation struct {
	actor Actor
	tail  Tail
	isNot bool
}

func createOperation(
	actor Actor,
	isNot bool,
) Operation {
	return createOperationInternally(actor, isNot, nil)
}

func createOperationWithTail(
	actor Actor,
	isNot bool,
	tail Tail,
) Operation {
	return createOperationInternally(actor, isNot, tail)
}

func createOperationInternally(
	actor Actor,
	isNot bool,
	tail Tail,
) Operation {
	out := operation{
		actor: actor,
		isNot: isNot,
		tail:  tail,
	}

	return &out
}

// Actor returns the actor
func (obj *operation) Actor() Actor {
	return obj.actor
}

// IsNot returns true if not, false otherwise
func (obj *operation) IsNot() bool {
	return obj.isNot
}

// HasTail returns true if there is a tail, false otherwise
func (obj *operation) HasTail() bool {
	return obj.tail != nil
}

// Tail returns the tail
func (obj *operation) Tail() Tail {
	return obj.tail
}
