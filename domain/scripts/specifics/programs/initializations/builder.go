package initializations

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/assignments"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/containers"
)

type builder struct {
	hashAdapter hash.Adapter
	container   containers.Container
	assignment  assignments.Assignment
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		container:   nil,
		assignment:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithContainer adds a container to the builder
func (app *builder) WithContainer(container containers.Container) Builder {
	app.container = container
	return app
}

// WithAssignment adds an assignment to the builder
func (app *builder) WithAssignment(assignment assignments.Assignment) Builder {
	app.assignment = assignment
	return app
}

// Now builds a new Initialization instance
func (app *builder) Now() (Initialization, error) {
	if app.container == nil {
		return nil, errors.New("the container is mandatory in order to build an Initialization instance")
	}

	if app.assignment == nil {
		return nil, errors.New("the assignment is mandatory in order to build an Initialization instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.container.Hash().Bytes(),
		app.assignment.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createInitialization(
		*pHash,
		app.container,
		app.assignment,
	), nil
}
