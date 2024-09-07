package resources

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/commits/modifications/resources/pointers"
)

type resourceBuilder struct {
	hashAdapter hash.Adapter
	identifier  string
	pointer     pointers.Pointer
}

func createResourceBuilder(
	hashAdapter hash.Adapter,
) ResourceBuilder {
	out := resourceBuilder{
		hashAdapter: hashAdapter,
		identifier:  "",
		pointer:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *resourceBuilder) Create() ResourceBuilder {
	return createResourceBuilder(
		app.hashAdapter,
	)
}

// WithIdentifier adds an identifier to the builder
func (app *resourceBuilder) WithIdentifier(identifier string) ResourceBuilder {
	app.identifier = identifier
	return app
}

// WithPointer add pointer to the builder
func (app *resourceBuilder) WithPointer(pointer pointers.Pointer) ResourceBuilder {
	app.pointer = pointer
	return app
}

// Now builds a new Resource instance
func (app *resourceBuilder) Now() (Resource, error) {
	if app.identifier == "" {
		return nil, errors.New("the identifier is mandatory in order to build a Resource instance")
	}

	if app.pointer == nil {
		return nil, errors.New("the pointer is mandatory in order to build a Resource instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.identifier),
		app.pointer.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createResource(
		*pHash,
		app.identifier,
		app.pointer,
	), nil
}
