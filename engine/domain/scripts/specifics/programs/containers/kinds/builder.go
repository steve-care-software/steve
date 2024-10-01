package kinds

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/containers/kinds/numerics"
)

type builder struct {
	hashAdapter hash.Adapter
	numeric     numerics.Numeric
	pEngine     *uint8
	pRemaining  *uint8
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		numeric:     nil,
		pEngine:     nil,
		pRemaining:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithNumeric adds a numeric to the builder
func (app *builder) WithNumeric(numeric numerics.Numeric) Builder {
	app.numeric = numeric
	return app
}

// WithEngine adds an engine to the builder
func (app *builder) WithEngine(engine uint8) Builder {
	app.pEngine = &engine
	return app
}

// WithRemaining adds a remaining to the builder
func (app *builder) WithRemaining(remaining uint8) Builder {
	app.pRemaining = &remaining
	return app
}

// Now builds a new Kind instance
func (app *builder) Now() (Kind, error) {
	data := [][]byte{}
	if app.numeric != nil {
		data = append(data, app.numeric.Hash().Bytes())
	}

	if app.pEngine != nil {
		engine := *app.pEngine
		if engine > EngineRoute {
			str := fmt.Sprintf("the engine (%d) is invalid when building the Kind instance", engine)
			return nil, errors.New(str)
		}

		data = append(data, []byte{engine})
	}

	if app.pRemaining != nil {
		remaining := *app.pRemaining
		if remaining > RemainingBool {
			str := fmt.Sprintf("the remaining (%d) is invalid when building the Kind instance", remaining)
			return nil, errors.New(str)
		}

		data = append(data, []byte{remaining})
	}

	if len(data) != 1 {
		return nil, errors.New("the Kind is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.numeric != nil {
		return createKindWithNumeric(
			*pHash,
			app.numeric,
		), nil
	}

	if app.pEngine != nil {
		return createKindWithEngine(
			*pHash,
			app.pEngine,
		), nil
	}

	return createKindWithRemaining(
		*pHash,
		app.pEngine,
	), nil
}
