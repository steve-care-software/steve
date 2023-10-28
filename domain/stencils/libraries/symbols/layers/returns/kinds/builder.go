package kinds

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	isContinue  bool
	isPrompt    bool
	execute     []string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		isContinue:  false,
		isPrompt:    false,
		execute:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithExecute adds execute commands to the builder
func (app *builder) WithExecute(execute []string) Builder {
	app.execute = execute
	return app
}

// IsContinue flags the builder as a continue
func (app *builder) IsContinue() Builder {
	app.isContinue = true
	return app
}

// IsPrompt flags the builder as a prompt
func (app *builder) IsPrompt() Builder {
	app.isPrompt = true
	return app
}

// Now builds a new Kind instance
func (app *builder) Now() (Kind, error) {

	data := [][]byte{}
	if app.isContinue {
		data = append(data, []byte{0})
	}

	if app.isPrompt {
		data = append(data, []byte{1})
	}

	if app.execute != nil {
		data = append(data, []byte{2})
		for _, oneCmd := range app.execute {
			data = append(data, []byte(oneCmd))
		}
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.isContinue {
		return createKindWithContinue(*pHash), nil
	}

	if app.isPrompt {
		return createKindWithPrompt(*pHash), nil
	}

	if app.execute != nil && len(app.execute) <= 0 {
		app.execute = nil
	}

	if app.execute != nil {
		return createKindWithExecute(*pHash, app.execute), nil
	}

	return nil, errors.New("the Kind is invalid")
}
