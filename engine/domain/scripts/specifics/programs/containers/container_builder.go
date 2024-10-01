package containers

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/containers/kinds"
)

type containerBuilder struct {
	hashAdapter hash.Adapter
	pFlag       *uint8
	kind        kinds.Kind
}

func createContainerBuilder(
	hashAdapter hash.Adapter,
) ContainerBuilder {
	out := containerBuilder{
		hashAdapter: hashAdapter,
		pFlag:       nil,
		kind:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *containerBuilder) Create() ContainerBuilder {
	return createContainerBuilder(
		app.hashAdapter,
	)
}

// WithFlag adds a flag to the builder
func (app *containerBuilder) WithFlag(flag uint8) ContainerBuilder {
	app.pFlag = &flag
	return app
}

// WithKind adds a kind to the builder
func (app *containerBuilder) WithKind(kind kinds.Kind) ContainerBuilder {
	app.kind = kind
	return app
}

// Now builds a new Container
func (app *containerBuilder) Now() (Container, error) {
	if app.pFlag == nil {
		return nil, errors.New("the flag is mandatory in order to build a Container instance")
	}

	if app.kind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Container instance")
	}

	flag := *app.pFlag
	if flag > FlagSortedSet {
		str := fmt.Sprintf("the flag (%d) is invalid when building the Container instance", flag)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		{flag},
		app.kind.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createContainer(
		*pHash,
		flag,
		app.kind,
	), nil
}
