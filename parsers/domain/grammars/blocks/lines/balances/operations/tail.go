package operations

type tail struct {
	operator uint8
	actor    Actor
}

func createTail(
	operator uint8,
	actor Actor,
) Tail {
	out := tail{
		operator: operator,
		actor:    actor,
	}

	return &out
}

// Operator returns the operator
func (obj *tail) Operator() uint8 {
	return obj.operator
}

// Actor returns the actor
func (obj *tail) Actor() Actor {
	return obj.actor
}
