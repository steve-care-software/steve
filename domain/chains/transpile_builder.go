package chains

import (
	"errors"

	"github.com/steve-care-software/steve/domain/chains/nfts"
	"github.com/steve-care-software/steve/domain/hash"
)

type transpileBuilder struct {
	hashAdapter hash.Adapter
	bridge      nfts.NFT
	target      nfts.NFT
	next        Chain
}

func createTranspileBuilder(
	hashAdapter hash.Adapter,
) TranspileBuilder {
	out := transpileBuilder{
		hashAdapter: hashAdapter,
		bridge:      nil,
		target:      nil,
		next:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *transpileBuilder) Create() TranspileBuilder {
	return createTranspileBuilder(
		app.hashAdapter,
	)
}

// WithBridge adds a bridge to the builder
func (app *transpileBuilder) WithBridge(bridge nfts.NFT) TranspileBuilder {
	app.bridge = bridge
	return app
}

// WithTarget adds a bridge to the builder
func (app *transpileBuilder) WithTarget(target nfts.NFT) TranspileBuilder {
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

	data := [][]byte{
		app.bridge.Hash().Bytes(),
		app.target.Hash().Bytes(),
	}

	if app.next != nil {
		data = append(data, app.next.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.next != nil {
		return createTranspileWithNext(*pHash, app.bridge, app.target, app.next), nil
	}

	return createTranspile(*pHash, app.bridge, app.target), nil
}
