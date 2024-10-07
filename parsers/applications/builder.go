package applications

import (
	"github.com/steve-care-software/steve/parsers/domain/asts"
	"github.com/steve-care-software/steve/parsers/domain/asts/instructions"
	"github.com/steve-care-software/steve/parsers/domain/queries"
	"github.com/steve-care-software/steve/parsers/domain/walkers"
	"github.com/steve-care-software/steve/parsers/domain/walkers/elements"
)

type builder struct {
	queryAdapterFactory queries.AdapterFactory
	elementsAdapter     instructions.ElementsAdapter
	astAdapter          asts.Adapter
	elementAdapter      elements.Adapter
	tokensBuilder       instructions.TokensBuilder
	pElement            *elements.Element
}

func createBuilder(
	queryAdapterFactory queries.AdapterFactory,
	elementsAdapter instructions.ElementsAdapter,
	astAdapter asts.Adapter,
	elementAdapter elements.Adapter,
	tokensBuilder instructions.TokensBuilder,
) Builder {
	out := builder{
		queryAdapterFactory: queryAdapterFactory,
		elementsAdapter:     elementsAdapter,
		astAdapter:          astAdapter,
		elementAdapter:      elementAdapter,
		tokensBuilder:       tokensBuilder,
		pElement:            nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.queryAdapterFactory,
		app.elementsAdapter,
		app.astAdapter,
		app.elementAdapter,
		app.tokensBuilder,
	)
}

// WithElement adds an element to the builder
func (app *builder) WithElement(ins elements.Element) Builder {
	app.pElement = &ins
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	var walker walkers.Walker
	if app.pElement != nil {
		retWalker, err := app.elementAdapter.ToWalker(*app.pElement)
		if err != nil {
			return nil, err
		}

		walker = retWalker
	}

	queryAdapter, err := app.queryAdapterFactory.Create()
	if err != nil {
		return nil, err
	}

	return createApplication(
		app.elementsAdapter,
		app.astAdapter,
		queryAdapter,
		app.tokensBuilder,
		walker,
	), nil
}
