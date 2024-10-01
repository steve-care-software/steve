package operations

import "github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/operations/selectors"

type actor struct {
	selector  selectors.Selector
	operation Operation
}

func createActorWithSelector(
	selector selectors.Selector,
) Actor {
	return createActorInternally(
		selector,
		nil,
	)
}

func createActorWithOperation(
	operation Operation,
) Actor {
	return createActorInternally(
		nil,
		operation,
	)
}

func createActorInternally(
	selector selectors.Selector,
	operation Operation,
) Actor {
	out := actor{
		selector:  selector,
		operation: operation,
	}

	return &out
}

// IsSelector returns true if there is a selector, false otherwise
func (obj *actor) IsSelector() bool {
	return obj.selector != nil
}

// Selector returns the selector, if any
func (obj *actor) Selector() selectors.Selector {
	return obj.selector
}

// IsOperation returns true if there is an operation, false otherwise
func (obj *actor) IsOperation() bool {
	return obj.operation != nil
}

// Operation returns the operation, if any
func (obj *actor) Operation() Operation {
	return obj.operation
}
