package points

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/steve-care-software/steve/commons/hash"
)

type pointBuilder struct {
	hashAdapter hash.Adapter
	name        string
	pKind       *uint8
	pStructure  *uint8
}

func createPointBuilder(
	hashAdapter hash.Adapter,
) PointBuilder {
	out := pointBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		pKind:       nil,
		pStructure:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *pointBuilder) Create() PointBuilder {
	return createPointBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *pointBuilder) WithName(name string) PointBuilder {
	app.name = name
	return app
}

// WithKind adds a kind to the builder
func (app *pointBuilder) WithKind(kind uint8) PointBuilder {
	app.pKind = &kind
	return app
}

// WithStructure adds a structure to the builder
func (app *pointBuilder) WithStructure(structure uint8) PointBuilder {
	app.pStructure = &structure
	return app
}

// Now builds a new Point instance
func (app *pointBuilder) Now() (Point, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Point instance")
	}

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Point instance")
	}

	kind := *app.pKind
	if kind > KindFloat {
		str := fmt.Sprintf("the kind (%d) is invalid while building the Point instance", kind)
		return nil, errors.New(str)
	}

	data := [][]byte{
		[]byte(app.name),
		[]byte(strconv.Itoa(int(*app.pKind))),
	}

	if app.pStructure != nil {
		structure := *app.pStructure
		if structure > StructureVector {
			str := fmt.Sprintf("the structure (%d) is invalid while building the Point instance", structure)
			return nil, errors.New(str)
		}

		data = append(data, []byte(strconv.Itoa(int(*app.pStructure))))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.pStructure != nil {
		return createPointWithStructure(
			*pHash,
			app.name,
			*app.pKind,
			app.pStructure,
		), nil
	}

	return createPoint(
		*pHash,
		app.name,
		*app.pKind,
	), nil
}
