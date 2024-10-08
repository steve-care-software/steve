package roots

import (
	"errors"
	"strings"

	"github.com/steve-care-software/steve/engine/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/pipelines/executions/suites"
	"github.com/steve-care-software/steve/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	head        heads.Head
	pipelines   []string
	suites      suites.Suites
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		head:        nil,
		pipelines:   nil,
		suites:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithHead adds an head to the builder
func (app *builder) WithHead(head heads.Head) Builder {
	app.head = head
	return app
}

// WithPipelines add pipelines to the builder
func (app *builder) WithPipelines(pipelines []string) Builder {
	app.pipelines = pipelines
	return app
}

// WithSuites add suites to the builder
func (app *builder) WithSuites(suites suites.Suites) Builder {
	app.suites = suites
	return app
}

// Now builds a new Root instance
func (app *builder) Now() (Root, error) {
	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build a Root instance")
	}

	if app.pipelines != nil && len(app.pipelines) <= 0 {
		app.pipelines = nil
	}

	if app.pipelines == nil {
		return nil, errors.New("the pipelines are mandatory in order to build a Root instance")
	}

	data := [][]byte{
		app.head.Hash().Bytes(),
		[]byte(strings.Join(app.pipelines, ",")),
	}

	if app.suites != nil {
		data = append(data, app.suites.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.suites != nil {
		return createRootWithSuites(
			*pHash,
			app.head,
			app.pipelines,
			app.suites,
		), nil
	}

	return createRoot(
		*pHash,
		app.head,
		app.pipelines,
	), nil
}
