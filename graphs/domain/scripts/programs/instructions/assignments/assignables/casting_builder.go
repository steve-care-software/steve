package assignables

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/assignments/assignables"
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/kinds"
)

type castingBuilder struct {
	assignable assignables.Assignable
	kind       kinds.Kind
}

func createCastingBuilder() CastingBuilder {
	return &castingBuilder{
		assignable: nil,
		kind:       nil,
	}
}

// Create initializes the casting builder
func (app *castingBuilder) Create() CastingBuilder {
	return createCastingBuilder()
}

// WithAssignable adds an assignable to the casting builder
func (app *castingBuilder) WithAssignable(assignable assignables.Assignable) CastingBuilder {
	app.assignable = assignable
	return app
}

// WithKind adds a kind to the casting builder
func (app *castingBuilder) WithKind(kind kinds.Kind) CastingBuilder {
	app.kind = kind
	return app
}

// Now builds a new Casting instance
func (app *castingBuilder) Now() (Casting, error) {
	if app.assignable == nil {
		return nil, errors.New("the assignable is mandatory to build a Casting instance")
	}
	if app.kind == nil {
		return nil, errors.New("the kind is mandatory to build a Casting instance")
	}

	return createCasting(app.assignable, app.kind), nil
}
