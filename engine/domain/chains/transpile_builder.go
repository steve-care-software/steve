package chains

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/hash"
)

type transpileBuilder struct {
	bridge hash.Hash
	target hash.Hash
	next   Chain
}

func createTranspileBuilder() TranspileBuilder {
	out := transpileBuilder{
		bridge: nil,
		target: nil,
		next:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *transpileBuilder) Create() TranspileBuilder {
	return createTranspileBuilder()
}

// WithBridge adds a bridge to the builder
func (app *transpileBuilder) WithBridge(bridge hash.Hash) TranspileBuilder {
	app.bridge = bridge
	return app
}

// WithTarget adds a bridge to the builder
func (app *transpileBuilder) WithTarget(target hash.Hash) TranspileBuilder {
	app.target = target
	return app
}

// WithNext adds a next chain to the builder
func (app *transpileBuilder) WithNext(next Chain) TranspileBuilder {
	app.next = next
	return app
}

// Now builds a new Transpile instance
func (app *transpileBuilder) Now() (Transpile, error) {
	if app.bridge == nil {
		return nil, errors.New("the bridge is mandatory in order to build a Transpile instance")
	}

	if app.target == nil {
		return nil, errors.New("the target is mandatory in order to build a Transpile instance")
	}

	if app.next != nil {
		return createTranspileWithNext(app.bridge, app.target, app.next), nil
	}

	return createTranspile(app.bridge, app.target), nil
}
